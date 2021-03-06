package sbf

import (
	"log"
	"regexp"
	"strings"

	"github.com/snksoft/crc"
)

const promptRegExp = `COM\\d>|USB\\d>|OTG\\d>|IP\\d{2}>|BT\\d{2}>` ///< regular expression defining what a prompt looks like
const promptLength = 5                                             ///< length of a prompt
const maxASCIIDisplaySize = 16384
const maxFormattedInformationBlockSize = 4096
const maxASCIICommandReplySize = 4096
const maxEventSize = 256

var hasPrompt bool = false

//Taken from parse function in ssnrx.cpp
func Parse(buffer *[]byte) []interface{} {
	done := false
	for !done {
		// We are looking for either
		// - a prompt (which may be caused by sending [Enter])
		// - a message, starting with '$', and followed by a character
		//   indicating the kind of message
		bufferSize := len(*buffer)
		ndx := 0
		for ndx = 0; ndx < bufferSize; ndx++ {
			if (*buffer)[ndx] == '>' || (*buffer)[ndx] == '$' {
				break
			}
		}

		//Make sure we haven't gone past the end of the buffer
		if ndx < bufferSize {
			if (*buffer)[ndx] == '>' {
				// '>' terminates a prompt
				done = handleCommandPrompt(buffer, ndx)
			} else {
				// '$' was found
				done = handleReceivedData(buffer, ndx)
			}
		} else {
			// We've reached the end of the buffer with nothing found. Discard all data,
			//except for the last PromptLength-1, because we may have the start of a new prompt
			done = true
			if bufferSize > promptLength-1 {
				*buffer = (*buffer)[promptLength-1:]
				if bufferSize-len(*buffer) > 0 {
					log.Printf("[DEBUG] parse - Discarding %d bytes from buffer\n", bufferSize-len(*buffer))
				}
			}
		}
	}
}

func handleCommandPrompt(buffer *[]byte, ndx int) (bool, []interface{}) {
	var prompt []byte
	payloads := []interface{}{}

	//See if we have enough characters to qualify for a prompt
	if ndx+1 >= promptLength {
		//See if the preceding characters match what we would expect a prompt to look like
		matched, err := regexp.Match(promptRegExp, (*buffer)[(ndx-promptLength+1):ndx])
		if err != nil {
			log.Printf("[ERROR] parse - Error evaluating regular expression: %s\n", err.Error())
		} else {
			if matched {
				prompt = (*buffer)[(ndx - promptLength + 1):ndx]
			}
		}
	}

	if len(prompt) > 0 {
		if ndx-len(prompt)+1 > 0 {
			log.Printf("[DEBUG] parse - Command prompt received: %s\n", string(prompt))
		}
		// TODO - Implement this code if we need the adapter to handle command prompts
		//setPrompt(prompt);
	}
	log.Printf("[DEBUG] parse - Removing prompt from buffer: %s\n", string(prompt))
	*buffer = append((*buffer)[:ndx-promptLength], (*buffer)[ndx+1:]...)
	return false, payloads
}

func handleReceivedData(buffer *[]byte, ndx int) (bool, []interface{}) {
	payloads := []interface{}{}

	if ndx > 0 {
		*buffer = (*buffer)[ndx:]
	}
	if len(*buffer) >= 2 {
		if (*buffer)[1] == '@' || (*buffer)[1] == 'R' || (*buffer)[1] == 'T' || (*buffer)[1] == '-' {
			notEnoughData := false
			processedBytes := 0

			if (*buffer)[1] == '@' {
				processedBytes, notEnoughData = parseSBF(buffer, 0)
			} else if (*buffer)[1] == 'R' {
				processedBytes, notEnoughData = parseASCIICommandReply(buffer, 0)
			} else if (*buffer)[1] == 'T' {
				if len(*buffer) < 3 {
					notEnoughData = true
				} else {
					if (*buffer)[2] == 'D' {
						processedBytes, notEnoughData = parseASCIIDisplay(buffer, 0)
					} else if (*buffer)[2] == 'E' {
						processedBytes, notEnoughData = parseEvent(buffer, 0)
					}
				}
			} else if (*buffer)[1] == '-' {
				processedBytes, notEnoughData = parseFormattedInformationBlock(buffer, 0)
			}

			if processedBytes > 0 {
				*buffer = (*buffer)[processedBytes:]
			} else {
				if notEnoughData {
					// the buffer does not yet contain enough data to parse the message
					// so we return control, and parsing will be re-attempted upon receiving more data
					return true, payloads
				} else {
					// an error has occured when trying to parse the message
					*buffer = (*buffer)[2:]
					log.Printf("[DEBUG] handleReceivedData - Discarding %d bytes from buffer\n", 2)
				}
			}
		} else {
			*buffer = (*buffer)[1:]
		}
		return false, payloads
	} else {
		// the '$' was not followed yet by a byte indicating the type of message
		// so we return control, and parsing will be reattempted upon receiving more data
		return true, payloads
	}
}

func parseSBF(buffer *[]byte, ndx int) (int, bool) {
	bufferSize := len(*buffer)
	notEnoughData := false

	if string((*buffer)[ndx:ndx+2]) != "$@" {
		return -1, notEnoughData
	}
	if bufferSize-ndx < 8 {
		notEnoughData = true
		return -1, notEnoughData
	}

	// get the length field from the SBF header
	length := uint16((*buffer)[6])<<8 + uint16((*buffer)[7])
	if length < 8 {
		log.Printf("[ERROR] parseSBF - Invalid SBF block length: %d\n", length)
		return -1, notEnoughData
	}
	if length > uint16(bufferSize-ndx) {
		notEnoughData = true
		return -1, notEnoughData
	}

	//Parse the CRC
	expectedCRC := uint16((*buffer)[ndx+2])<<8 + uint16((*buffer)[ndx+3])
	//Recalculate the CRC
	actualCRC := crc.CalculateCRC(crc.CCITT, (*buffer)[ndx+4:length-4])

	if actualCRC != uint64(expectedCRC) {
		log.Printf("[ERROR] parseSBF - SBF CRC error. Expected: %d, calculated:%d\n", expectedCRC, actualCRC)
		return -1, notEnoughData
	}
	// emit newSBFBlock(mBuffer.mid(startIndex, length));
	// emit newSBFBlockWithId(mBuffer.mid(startIndex, length), (actualID & 0x1fff), (actualID >> 13));

	//TODO - Send the SBF block for parsing
	handleSbfBlock((*buffer)[ndx:length])
	return int(length), notEnoughData
}

func parseASCIICommandReply(buffer *[]byte, ndx int) (int, bool) {
	notEnoughData := false
	if string((*buffer)[ndx:ndx+2]) != "$R" {
		return -1, notEnoughData
	}
	// mPromptTimer.start();  // restart timer (from zero again)
	// mHasPrompt = false;
	// int endIndex = searchEndOfAsciiMessage(startIndex, sMaxASCIICommandReplySize, outNotEnoughData);
	// if (endIndex != -1) {
	//   QString prompt = mBuffer.mid(endIndex - sPromptLength, sPromptLength);
	//   bool error = (mBuffer.at(startIndex + 2) == '?');
	//   emit newCommandReply(mBuffer.mid(startIndex, endIndex - startIndex - prompt.size() - 2), error);

	//   if (prompt == "STOP>") {
	//     emit stopReceived();
	//     mPromptTimer.stop();
	//   }
	//   else if (prompt != "---->") {
	//     setPrompt(prompt);
	//   }

	// consume "\r\n" if present
	//   if ((mBuffer.size() > endIndex) && (mBuffer.at(endIndex) == '\r')) {
	//     endIndex++;
	//     if ((mBuffer.size() > endIndex) && (mBuffer.at(endIndex) == '\n')) {
	//       endIndex++;
	//     }
	//   }
	//   return endIndex - startIndex;
	// } else {
	return -1, notEnoughData
	// }
}

func parseASCIIDisplay(buffer *[]byte, ndx int) (int, bool) {
	notEnoughData := false
	if len(*buffer)-ndx < 3 {
		notEnoughData = true
		return -1, notEnoughData
	}
	if string(*buffer)[ndx:ndx+3] != "$TD" {
		return -1, notEnoughData
	}

	endIndex := strings.Index(string((*buffer)[ndx:]), "\r\n####>\r\n")
	if endIndex != -1 && (endIndex < ndx+maxASCIIDisplaySize) {
		// the $TD is followed by \r\n, which is also stripped

		// TODO - Do we publish this to the platform?
		//
		//asciiDisplayContents := (*buffer)[ndx+5 : endIndex-(ndx+5)]
		// emit newASCIIDisplay(asciiDisplayContents);
		return endIndex + 9 - ndx, notEnoughData // total processed bytes, including start and end delimiters
	} else {
		if endIndex == -1 && (len(*buffer) < ndx+maxASCIIDisplaySize) {
			notEnoughData = true
		} else {
			// maximum length of ASCII display exceeded without finding the end
			notEnoughData = false
		}
		return -1, notEnoughData
	}
}

func parseEvent(buffer *[]byte, ndx int) (int, bool) {
	notEnoughData := false
	if len(*buffer)-ndx < 3 {
		notEnoughData = true
		return -1, notEnoughData
	}
	if string(*buffer)[ndx:ndx+3] != "$TE" {
		return -1, notEnoughData
	}

	endIndex := strings.Index(string((*buffer)[ndx:]), "\r\n")

	if endIndex != -1 && endIndex < ndx+maxEventSize {
		// TODO - Do we publish this to the platform?
		//
		//event := string((*buffer)[ndx+4 : endIndex-(ndx+4)])
		//emit newEvent(event);
		return endIndex + 2 - ndx, notEnoughData // total processed bytes, including start and end delimiters
	} else {
		if endIndex == -1 && len(*buffer) < ndx+maxEventSize {
			notEnoughData = true
		} else {
			// maximum length of event line exceeded without finding the end
			notEnoughData = false
		}
		return -1, notEnoughData
	}
}

func parseFormattedInformationBlock(buffer *[]byte, ndx int) (int, bool) {
	notEnoughData := false
	// mHasPrompt = false;
	// mPromptTimer.start(); // restart timer (from zero again)

	// try to parse the first line as "<-- BLOCK I / N\r\n"
	// int indexOfEOL = mBuffer.indexOf("\r\n", startIndex);
	// if (indexOfEOL == -1)
	// {
	//   if ((mBuffer.size() - startIndex) < 30)
	//   {
	// 30 is taken (arbitrarily) as a reasonable small maximum lenght
	// for the first line of the formatted information block reply.
	//     *outNotEnoughData = true;
	//   }
	//   return -1;
	// }
	// else
	// {
	//   static QRegExp blockHeaderRegExp("\\$-- BLOCK (\\d+) / (\\d+)"); // (greedy by default)
	//   QString firstLine = mBuffer.mid(startIndex, indexOfEOL - startIndex);
	//   if (blockHeaderRegExp.indexIn(firstLine) != 0)
	//   {
	// the first line should immediately start matching the given regular expression
	//     return -1;
	//   }
	//   else
	//   {
	//     int blockIndex = blockHeaderRegExp.cap(1).toInt();
	//     int nrOfBlocks = blockHeaderRegExp.cap(2).toInt();

	// now search for the end of the block
	//     int endIndex = searchEndOfAsciiMessage(startIndex, sMaxFormattedInformationBlockSize, outNotEnoughData);
	//     if (endIndex != -1)
	//     {
	//       QString prompt = mBuffer.mid(endIndex - sPromptLength, sPromptLength);
	//       emit newFormattedInformationBlock(mBuffer.mid(startIndex, endIndex - startIndex - prompt.size() - 2), blockIndex, nrOfBlocks);

	//       if (prompt == "STOP>")
	//       {
	//         emit stopReceived();
	//       }
	//       else if (prompt != "---->")
	//       {
	//         setPrompt(prompt);
	//       }
	//       return endIndex - startIndex;
	//     }
	//     else
	//     {
	return -1, notEnoughData
	//     }
	//   }
	// }
}

func handleSbfBlock(buffer []byte) {
	//Parse the SBF ID
	sbfID := uint16(buffer[4])<<8 + uint16(buffer[5])

	switch sbfID {
	/* Measurement Blocks */
	case sbfnr_GenMeasEpoch_1: //    = 5944
	//case sbfid_GenMeasEpoch_1_0: //= 5944 | 0x0
	case sbfnr_MeasEpoch_2: //= 4027
	//case sbfid_MeasEpoch_2_0: //= 4027 | 0x0
	case sbfid_MeasEpoch_2_1: //= 4027 | 0x2000
	case sbfnr_MeasExtra_1: //= 4000
	//case sbfid_MeasExtra_1_0: //= 4000 | 0x0
	case sbfid_MeasExtra_1_1: //= 4000 | 0x2000
	case sbfid_MeasExtra_1_2: //= 4000 | 0x4000
	case sbfid_MeasExtra_1_3: //= 4000 | 0x6000
	case sbfnr_MeasFullRange_1: //= 4098
	//case sbfid_MeasFullRange_1_0: //= 4098 | 0x0
	case sbfid_MeasFullRange_1_1: //= 4098 | 0x2000
	case sbfnr_Meas3Ranges_1: //= 4109
	//case sbfid_Meas3Ranges_1_0: //= 4109 | 0x0
	case sbfnr_Meas3CN0HiRes_1: //= 4110
	//case sbfid_Meas3CN0HiRes_1_0: //= 4110 | 0x0
	case sbfnr_Meas3Doppler_1: //= 4111
	//case sbfid_Meas3Doppler_1_0: //= 4111 | 0x0
	case sbfnr_Meas3PP_1: //= 4112
	//case sbfid_Meas3PP_1_0: //= 4112 | 0x0
	case sbfnr_Meas3MP_1: //= 4113
	//case sbfid_Meas3MP_1_0: //= 4113 | 0x0
	case sbfnr_IQCorr_1: //= 4046
	//case sbfid_IQCorr_1_0: //= 4046 | 0x0
	case sbfid_IQCorr_1_1: //= 4046 | 0x2000
	case sbfnr_ISMR_1: //= 4086
	//case sbfid_ISMR_1_0: //= 4086 | 0x0
	case sbfnr_SQMSamples_1: //= 4087
	//case sbfid_SQMSamples_1_0: //= 4087 | 0x0
	case sbfnr_EndOfMeas_1: //= 5922
	//case sbfid_EndOfMeas_1_0: //= 5922 | 0x0

	/* Navigation Page Blocks */
	case sbfnr_GPSRaw_1: //= 5895
	//case sbfid_GPSRaw_1_0: //= 5895 | 0x0
	case sbfnr_CNAVRaw_1: //= 5947
	//case sbfid_CNAVRaw_1_0: //= 5947 | 0x0
	case sbfnr_GEORaw_1: //= 5898
	//case sbfid_GEORaw_1_0: //= 5898 | 0x0
	case sbfnr_GPSRawCA_1: //= 4017
	//case sbfid_GPSRawCA_1_0: //= 4017 | 0x0
	case sbfnr_GPSRawL2C_1: //= 4018
	//case sbfid_GPSRawL2C_1_0: //= 4018 | 0x0
	case sbfnr_GPSRawL5_1: //= 4019
	//case sbfid_GPSRawL5_1_0: //= 4019 | 0x0
	case sbfnr_GPSRawL1C_1: //= 4221
	//case sbfid_GPSRawL1C_1_0: //= 4221 | 0x0
	case sbfnr_GLORawCA_1: //= 4026
	//case sbfid_GLORawCA_1_0: //= 4026 | 0x0
	case sbfnr_GALRawFNAV_1: //= 4022
	//case sbfid_GALRawFNAV_1_0: //= 4022 | 0x0
	case sbfnr_GALRawINAV_1: //= 4023
	//case sbfid_GALRawINAV_1_0: //= 4023 | 0x0
	case sbfnr_GALRawCNAV_1: //= 4024
	//case sbfid_GALRawCNAV_1_0: //= 4024 | 0x0
	case sbfnr_GALRawGNAV_1: //= 4025
	//case sbfid_GALRawGNAV_1_0: //= 4025 | 0x0
	case sbfnr_GALRawGNAVe_1: //= 4029
	//case sbfid_GALRawGNAVe_1_0: //= 4029 | 0x0
	case sbfnr_GEORawL1_1: //= 4020
	//case sbfid_GEORawL1_1_0: //= 4020 | 0x0
	case sbfnr_GEORawL5_1: //= 4021
	//case sbfid_GEORawL5_1_0: //= 4021 | 0x0
	case sbfnr_BDSRaw_1: //= 4047
	//case sbfid_BDSRaw_1_0: //= 4047 | 0x0
	case sbfnr_BDSRawB1C_1: //= 4218
	//case sbfid_BDSRawB1C_1_0: //= 4218 | 0x0
	case sbfnr_BDSRawB2a_1: //= 4219
	//case sbfid_BDSRawB2a_1_0: //= 4219 | 0x0
	case sbfnr_BDSRawB2b_1: //= 4242
	//case sbfid_BDSRawB2b_1_0: //= 4242 | 0x0
	case sbfnr_QZSRawL1CA_1: //= 4066
	//case sbfid_QZSRawL1CA_1_0: //= 4066 | 0x0
	case sbfnr_QZSRawL2C_1: //= 4067
	//case sbfid_QZSRawL2C_1_0: //= 4067 | 0x0
	case sbfnr_QZSRawL5_1: //= 4068
	//case sbfid_QZSRawL5_1_0: //= 4068 | 0x0
	case sbfnr_QZSRawL6_1: //= 4069
	//case sbfid_QZSRawL6_1_0: //= 4069 | 0x0
	case sbfnr_QZSRawL1C_1: //= 4227
	//case sbfid_QZSRawL1C_1_0: //= 4227 | 0x0
	case sbfnr_QZSRawL1S_1: //= 4228
	//case sbfid_QZSRawL1S_1_0: //= 4228 | 0x0
	case sbfnr_NAVICRaw_1: //= 4093
	//case sbfid_NAVICRaw_1_0: //= 4093 | 0x0
	case sbfnr_GNSSNavBits_1: //= 4088
	//case sbfid_GNSSNavBits_1_0: //= 4088 | 0x0
	case sbfnr_GNSSSymbols_1: //= 4099
	//case sbfid_GNSSSymbols_1_0: //= 4099 | 0x0

	/* GPS Decoded Message Blocks */
	case sbfnr_GPSNav_1: //= 5891
	//case sbfid_GPSNav_1_0: //= 5891 | 0x0
	case sbfnr_GPSAlm_1: //= 5892
	//case sbfid_GPSAlm_1_0: //= 5892 | 0x0
	case sbfnr_GPSIon_1: //= 5893
	//case sbfid_GPSIon_1_0: //= 5893 | 0x0
	case sbfnr_GPSUtc_1: //= 5894
	//case sbfid_GPSUtc_1_0: //= 5894 | 0x0
	case sbfnr_GPSCNav_1: //= 4042
	//case sbfid_GPSCNav_1_0: //= 4042 | 0x0

	/* GLONASS Decoded Message Blocks */
	case sbfnr_GLONav_1: //= 4004
	//case sbfid_GLONav_1_0: //= 4004 | 0x0
	case sbfnr_GLOAlm_1: //= 4005
	//case sbfid_GLOAlm_1_0: //= 4005 | 0x0
	case sbfnr_GLOTime_1: //= 4036
	//case sbfid_GLOTime_1_0: //= 4036 | 0x0

	/* Galileo Decoded Message Blocks */
	case sbfnr_GALNav_1: //= 4002
	//case sbfid_GALNav_1_0: //= 4002 | 0x0
	case sbfnr_GALAlm_1: //= 4003
	//case sbfid_GALAlm_1_0: //= 4003 | 0x0
	case sbfnr_GALIon_1: //= 4030
	//case sbfid_GALIon_1_0: //= 4030 | 0x0
	case sbfnr_GALUtc_1: //= 4031
	//case sbfid_GALUtc_1_0: //= 4031 | 0x0
	case sbfnr_GALGstGps_1: //= 4032
	//case sbfid_GALGstGps_1_0: //= 4032 | 0x0
	case sbfnr_GALSARRLM_1: //= 4034
	//case sbfid_GALSARRLM_1_0: //= 4034 | 0x0

	/* BeiDou Decoded Message Blocks */
	case sbfnr_BDSNav_1: //= 4081
	//case sbfid_BDSNav_1_0: //= 4081 | 0x0
	case sbfnr_BDSAlm_1: //= 4119
	//case sbfid_BDSAlm_1_0: //= 4119 | 0x0
	case sbfnr_BDSIon_1: //= 4120
	//case sbfid_BDSIon_1_0: //= 4120 | 0x0
	case sbfnr_BDSUtc_1: //= 4121
	//case sbfid_BDSUtc_1_0: //= 4121 | 0x0

	/* QZSS Decoded Message Blocks */
	case sbfnr_QZSNav_1: //= 4095
	//case sbfid_QZSNav_1_0: //= 4095 | 0x0
	case sbfnr_QZSAlm_1: //= 4116
	//case sbfid_QZSAlm_1_0: //= 4116 | 0x0

	/* SBAS L1 Decoded Message Blocks */
	case sbfnr_GEOMT00_1: //= 5925
	//case sbfid_GEOMT00_1_0: //= 5925 | 0x0
	case sbfnr_GEOPRNMask_1: //= 5926
	//case sbfid_GEOPRNMask_1_0: //= 5926 | 0x0
	case sbfnr_GEOFastCorr_1: //= 5927
	//case sbfid_GEOFastCorr_1_0: //= 5927 | 0x0
	case sbfnr_GEOIntegrity_1: //= 5928
	//case sbfid_GEOIntegrity_1_0: //= 5928 | 0x0
	case sbfnr_GEOFastCorrDegr_1: //= 5929
	//case sbfid_GEOFastCorrDegr_1_0: //= 5929 | 0x0
	case sbfnr_GEONav_1: //= 5896
	//case sbfid_GEONav_1_0: //= 5896 | 0x0
	case sbfnr_GEODegrFactors_1: //= 5930
	//case sbfid_GEODegrFactors_1_0: //= 5930 | 0x0
	case sbfnr_GEONetworkTime_1: //= 5918
	//case sbfid_GEONetworkTime_1_0: //= 5918 | 0x0
	case sbfnr_GEOAlm_1: //= 5897
	//case sbfid_GEOAlm_1_0: //= 5897 | 0x0
	case sbfnr_GEOIGPMask_1: //= 5931
	//case sbfid_GEOIGPMask_1_0: //= 5931 | 0x0
	case sbfnr_GEOLongTermCorr_1: //= 5932
	//case sbfid_GEOLongTermCorr_1_0: //= 5932 | 0x0
	case sbfnr_GEOIonoDelay_1: //= 5933
	//case sbfid_GEOIonoDelay_1_0: //= 5933 | 0x0
	case sbfnr_GEOServiceLevel_1: //= 5917
	//case sbfid_GEOServiceLevel_1_0: //= 5917 | 0x0
	case sbfnr_GEOClockEphCovMatrix_1: //= 5934
	//case sbfid_GEOClockEphCovMatrix_1_0: //= 5934 | 0x0

	/* SBAS L5 Decoded Message Blocks */
	case sbfnr_SBASL5Nav_1: //= 5958
	//case sbfid_SBASL5Nav_1_0: //= 5958 | 0x0
	case sbfnr_SBASL5Alm_1: //= 5959
	//case sbfid_SBASL5Alm_1_0: //= 5959 | 0x0

	/* Position, Velocity and Time Blocks */
	case sbfnr_PVTCartesian_1: //= 5903
	//case sbfid_PVTCartesian_1_0: //= 5903 | 0x0
	case sbfnr_PVTGeodetic_1: //= 5904
	//case sbfid_PVTGeodetic_1_0: //= 5904 | 0x0
	case sbfnr_DOP_1: //= 5909
	//case sbfid_DOP_1_0: //= 5909 | 0x0
	case sbfnr_PVTResiduals_1: //= 5910
	//case sbfid_PVTResiduals_1_0: //= 5910 | 0x0
	case sbfnr_RAIMStatistics_1: //= 5915
	//case sbfid_RAIMStatistics_1_0: //= 5915 | 0x0
	case sbfnr_PVTCartesian_2: //= 4006
	//case sbfid_PVTCartesian_2_0: //= 4006 | 0x0
	case sbfid_PVTCartesian_2_1: //= 4006 | 0x2000
	case sbfid_PVTCartesian_2_2: //= 4006 | 0x4000
	case sbfnr_PVTGeodetic_2: //= 4007
	//case sbfid_PVTGeodetic_2_0: //= 4007 | 0x0
	case sbfid_PVTGeodetic_2_1: //= 4007 | 0x2000
	case sbfid_PVTGeodetic_2_2: //= 4007 | 0x4000
	case sbfnr_PVTGeodeticAuth_1: //= 4232
	//case sbfid_PVTGeodeticAuth_1_0: //= 4232 | 0x0
	case sbfid_PVTGeodeticAuth_1_1: //= 4232 | 0x2000
	case sbfid_PVTGeodeticAuth_1_2: //= 4232 | 0x4000
	case sbfnr_PosCovCartesian_1: //= 5905
	//case sbfid_PosCovCartesian_1_0: //= 5905 | 0x0
	case sbfnr_PosCovGeodetic_1: //= 5906
	//case sbfid_PosCovGeodetic_1_0: //= 5906 | 0x0
	case sbfnr_VelCovCartesian_1: //= 5907
	//case sbfid_VelCovCartesian_1_0: //= 5907 | 0x0
	case sbfnr_VelCovGeodetic_1: //= 5908
	//case sbfid_VelCovGeodetic_1_0: //= 5908 | 0x0
	case sbfnr_DOP_2: //= 4001
	//case sbfid_DOP_2_0: //= 4001 | 0x0
	case sbfnr_PosCart_1: //= 4044
	//case sbfid_PosCart_1_0: //= 4044 | 0x0
	case sbfnr_PosLocal_1: //= 4052
	//case sbfid_PosLocal_1_0: //= 4052 | 0x0
	case sbfnr_PosProjected_1: //= 4094
	//case sbfid_PosProjected_1_0: //= 4094 | 0x0
	case sbfnr_PVTSatCartesian_1: //= 4008
	//case sbfid_PVTSatCartesian_1_0: //= 4008 | 0x0
	case sbfid_PVTSatCartesian_1_1: //= 4008 | 0x2000
	case sbfnr_PVTResiduals_2: //= 4009
	//case sbfid_PVTResiduals_2_0: //= 4009 | 0x0
	case sbfid_PVTResiduals_2_1: //= 4009 | 0x2000
	case sbfnr_RAIMStatistics_2: //= 4011
	//case sbfid_RAIMStatistics_2_0: //= 4011 | 0x0
	case sbfnr_GEOCorrections_1: //= 5935
	//case sbfid_GEOCorrections_1_0: //= 5935 | 0x0
	case sbfnr_BaseVectorCart_1: //= 4043
	//case sbfid_BaseVectorCart_1_0: //= 4043 | 0x0
	case sbfnr_BaseVectorGeod_1: //= 4028
	//case sbfid_BaseVectorGeod_1_0: //= 4028 | 0x0
	case sbfnr_Ambiguities_1: //= 4240
	//case sbfid_Ambiguities_1_0: //= 4240 | 0x0
	case sbfnr_EndOfPVT_1: //= 5921
	//case sbfid_EndOfPVT_1_0: //= 5921 | 0x0
	case sbfnr_BaseLine_1: //= 5950
	//case sbfid_BaseLine_1_0: //= 5950 | 0x0

	/* INS/GNSS Integrated Blocks */
	case sbfnr_IntPVCart_1: //= 4060
	//case sbfid_IntPVCart_1_0: //= 4060 | 0x0
	case sbfnr_IntPVGeod_1: //= 4061
	//case sbfid_IntPVGeod_1_0: //= 4061 | 0x0
	case sbfnr_IntPosCovCart_1: //= 4062
	//case sbfid_IntPosCovCart_1_0: //= 4062 | 0x0
	case sbfnr_IntVelCovCart_1: //= 4063
	//case sbfid_IntVelCovCart_1_0: //= 4063 | 0x0
	case sbfnr_IntPosCovGeod_1: //= 4064
	//case sbfid_IntPosCovGeod_1_0: //= 4064 | 0x0
	case sbfnr_IntVelCovGeod_1: //= 4065
	//case sbfid_IntVelCovGeod_1_0: //= 4065 | 0x0
	case sbfnr_IntAttEuler_1: //= 4070
	//case sbfid_IntAttEuler_1_0: //= 4070 | 0x0
	case sbfid_IntAttEuler_1_1: //= 4070 | 0x2000
	case sbfnr_IntAttCovEuler_1: //= 4072
	//case sbfid_IntAttCovEuler_1_0: //= 4072 | 0x0
	case sbfnr_IntPVAAGeod_1: //= 4045
	//case sbfid_IntPVAAGeod_1_0: //= 4045 | 0x0
	case sbfnr_INSNavCart_1: //= 4225
	//case sbfid_INSNavCart_1_0: //= 4225 | 0x0
	case sbfnr_INSNavGeod_1: //= 4226
	//case sbfid_INSNavGeod_1_0: //= 4226 | 0x0
	case sbfnr_IMUBias_1: //= 4241
	//case sbfid_IMUBias_1_0: //= 4241 | 0x0

	/* GNSS Attitude Blocks */
	case sbfnr_AttEuler_1: //= 5938
	//case sbfid_AttEuler_1_0: //= 5938 | 0x0
	case sbfnr_AttCovEuler_1: //= 5939
	//case sbfid_AttCovEuler_1_0: //= 5939 | 0x0
	case sbfnr_AuxAntPositions_1: //= 5942
	//case sbfid_AuxAntPositions_1_0: //= 5942 | 0x0
	case sbfnr_EndOfAtt_1: //= 5943
	//case sbfid_EndOfAtt_1_0: //= 5943 | 0x0
	case sbfnr_AttQuat_1: //= 5940
	//case sbfid_AttQuat_1_0: //= 5940 | 0x0
	case sbfnr_AttCovQuat_1: //= 5941
	//case sbfid_AttCovQuat_1_0: //= 5941 | 0x0

	/* Receiver Time Blocks */
	case sbfnr_ReceiverTime_1: //= 5914
	//case sbfid_ReceiverTime_1_0: //= 5914 | 0x0
	case sbfnr_xPPSOffset_1: //= 5911
	//case sbfid_xPPSOffset_1_0: //= 5911 | 0x0
	case sbfnr_SysTimeOffset_1: //= 4039
	//case sbfid_SysTimeOffset_1_0: //= 4039 | 0x0
	case sbfid_SysTimeOffset_1_1: //= 4039 | 0x2000

	/* External Event Blocks */
	case sbfnr_ExtEvent_1: //= 5924
	//case sbfid_ExtEvent_1_0: //= 5924 | 0x0
	case sbfid_ExtEvent_1_1: //= 5924 | 0x2000
	case sbfnr_ExtEventPVTCartesian_1: //= 4037
	//case sbfid_ExtEventPVTCartesian_1_0: //= 4037 | 0x0
	case sbfid_ExtEventPVTCartesian_1_1: //= 4037 | 0x2000
	case sbfid_ExtEventPVTCartesian_1_2: //= 4037 | 0x4000
	case sbfnr_ExtEventPVTGeodetic_1: //= 4038
	//case sbfid_ExtEventPVTGeodetic_1_0: //= 4038 | 0x0
	case sbfid_ExtEventPVTGeodetic_1_1: //= 4038 | 0x2000
	case sbfid_ExtEventPVTGeodetic_1_2: //= 4038 | 0x4000
	case sbfnr_ExtEventBaseVectCart_1: //= 4216
	//case sbfid_ExtEventBaseVectCart_1_0: //= 4216 | 0x0
	case sbfnr_ExtEventBaseVectGeod_1: //= 4217
	//case sbfid_ExtEventBaseVectGeod_1_0: //= 4217 | 0x0
	case sbfnr_ExtEventINSNavCart_1: //= 4229
	//case sbfid_ExtEventINSNavCart_1_0: //= 4229 | 0x0
	case sbfnr_ExtEventINSNavGeod_1: //= 4230
	//case sbfid_ExtEventINSNavGeod_1_0: //= 4230 | 0x0
	case sbfnr_ExtEventAttEuler_1: //= 4237
	//case sbfid_ExtEventAttEuler_1_0: //= 4237 | 0x0

	/* Differential Correction Blocks */
	case sbfnr_DiffCorrIn_1: //= 5919
	//case sbfid_DiffCorrIn_1_0: //= 5919 | 0x0
	case sbfnr_BaseStation_1: //= 5949
	//case sbfid_BaseStation_1_0: //= 5949 | 0x0
	case sbfnr_RTCMDatum_1: //= 4049
	//case sbfid_RTCMDatum_1_0: //= 4049 | 0x0
	case sbfnr_BaseLink_1: //= 5948
	//case sbfid_BaseLink_1_0: //= 5948 | 0x0

	/* L-Band Demodulator Blocks */
	case sbfnr_LBandReceiverStatus_1: //= 4200
	//case sbfid_LBandReceiverStatus_1_0: //= 4200 | 0x0
	case sbfnr_LBandTrackerStatus_1: //= 4201
	//case sbfid_LBandTrackerStatus_1_0: //= 4201 | 0x0
	case sbfid_LBandTrackerStatus_1_1: //= 4201 | 0x2000
	case sbfid_LBandTrackerStatus_1_2: //= 4201 | 0x4000
	case sbfid_LBandTrackerStatus_1_3: //= 4201 | 0x6000
	case sbfnr_LBAS1DecoderStatus_1: //= 4202
	//case sbfid_LBAS1DecoderStatus_1_0: //= 4202 | 0x0
	case sbfid_LBAS1DecoderStatus_1_1: //= 4202 | 0x2000
	case sbfid_LBAS1DecoderStatus_1_2: //= 4202 | 0x4000
	case sbfnr_LBAS1Messages_1: //= 4203
	//case sbfid_LBAS1Messages_1_0: //= 4203 | 0x0
	case sbfnr_LBandBeams_1: //= 4204
	//case sbfid_LBandBeams_1_0: //= 4204 | 0x0
	case sbfnr_LBandRaw_1: //= 4212
	//case sbfid_LBandRaw_1_0: //= 4212 | 0x0
	case sbfnr_FugroStatus_1: //= 4214
	//case sbfid_FugroStatus_1_0: //= 4214 | 0x0

	/* External Sensor Blocks */
	case sbfnr_ExtSensorMeas_1: //= 4050
	//case sbfid_ExtSensorMeas_1_0: //= 4050 | 0x0
	case sbfnr_ExtSensorStatus_1: //= 4056
	//case sbfid_ExtSensorStatus_1_0: //= 4056 | 0x0
	case sbfnr_ExtSensorSetup_1: //= 4057
	//case sbfid_ExtSensorSetup_1_0: //= 4057 | 0x0
	case sbfid_ExtSensorSetup_1_1: //= 4057 | 0x2000
	case sbfid_ExtSensorSetup_1_2: //= 4057 | 0x4000
	case sbfnr_ExtSensorStatus_2: //= 4223
	//case sbfid_ExtSensorStatus_2_0: //= 4223 | 0x0
	case sbfnr_ExtSensorInfo_1: //= 4222
	//case sbfid_ExtSensorInfo_1_0: //= 4222 | 0x0
	case sbfnr_IMUSetup_1: //= 4224
	//case sbfid_IMUSetup_1_0: //= 4224 | 0x0

	/* Status Blocks */
	case sbfnr_ReceiverStatus_1: //= 5913
	//case sbfid_ReceiverStatus_1_0: //= 5913 | 0x0
	case sbfnr_TrackingStatus_1: //= 5912
	//case sbfid_TrackingStatus_1_0: //= 5912 | 0x0
	case sbfnr_ChannelStatus_1: //= 4013
	//case sbfid_ChannelStatus_1_0: //= 4013 | 0x0
	case sbfnr_ReceiverStatus_2: //= 4014
	//case sbfid_ReceiverStatus_2_0: //= 4014 | 0x0
	case sbfid_ReceiverStatus_2_1: //= 4014 | 0x2000
	case sbfnr_SatVisibility_1: //= 4012
	//case sbfid_SatVisibility_1_0: //= 4012 | 0x0
	case sbfnr_InputLink_1: //= 4090
	//case sbfid_InputLink_1_0: //= 4090 | 0x0
	case sbfnr_OutputLink_1: //= 4091
	//case sbfid_OutputLink_1_0: //= 4091 | 0x0
	case sbfid_OutputLink_1_1: //= 4091 | 0x2000
	//case sbfnr_NTRIPClientStatus_1: //= 4053
	case sbfid_NTRIPClientStatus_1_0: //= 4053 | 0x0
	case sbfnr_NTRIPServerStatus_1: //= 4122
	//case sbfid_NTRIPServerStatus_1_0: //= 4122 | 0x0
	case sbfnr_IPStatus_1: //= 4058
	//case sbfid_IPStatus_1_0: //= 4058 | 0x0
	case sbfid_IPStatus_1_1: //= 4058 | 0x2000
	case sbfnr_WiFiAPStatus_1: //= 4054
	//case sbfid_WiFiAPStatus_1_0: //= 4054 | 0x0
	case sbfnr_WiFiClientStatus_1: //= 4096
	//case sbfid_WiFiClientStatus_1_0: //= 4096 | 0x0
	case sbfnr_CellularStatus_1: //= 4055
	//case sbfid_CellularStatus_1_0: //= 4055 | 0x0
	case sbfid_CellularStatus_1_1: //= 4055 | 0x2000
	case sbfnr_BluetoothStatus_1: //= 4051
	//case sbfid_BluetoothStatus_1_0: //= 4051 | 0x0
	case sbfnr_DynDNSStatus_1: //= 4105
	//case sbfid_DynDNSStatus_1_0: //= 4105 | 0x0
	case sbfid_DynDNSStatus_1_1: //= 4105 | 0x2000
	case sbfnr_BatteryStatus_1: //= 4083
	//case sbfid_BatteryStatus_1_0: //= 4083 | 0x0
	case sbfid_BatteryStatus_1_1: //= 4083 | 0x2000
	case sbfid_BatteryStatus_1_2: //= 4083 | 0x4000
	case sbfnr_PowerStatus_1: //= 4101
	//case sbfid_PowerStatus_1_0: //= 4101 | 0x0
	case sbfnr_QualityInd_1: //= 4082
	//case sbfid_QualityInd_1_0: //= 4082 | 0x0
	case sbfnr_DiskStatus_1: //= 4059
	//case sbfid_DiskStatus_1_0: //= 4059 | 0x0
	case sbfid_DiskStatus_1_1: //= 4059 | 0x2000
	case sbfnr_LogStatus_1: //= 4102
	//case sbfid_LogStatus_1_0: //= 4102 | 0x0
	case sbfnr_UHFStatus_1: //= 4085
	//case sbfid_UHFStatus_1_0: //= 4085 | 0x0
	case sbfnr_RFStatus_1: //= 4092
		//case sbfid_RFStatus_1_0: //= 4092 | 0x0
		handleRFStatus(buffer)
	case sbfnr_RIMSHealth_1: //= 4089
	//case sbfid_RIMSHealth_1_0: //= 4089 | 0x0
	case sbfnr_OSNMAStatus_1: //= 4231
	//case sbfid_OSNMAStatus_1_0: //= 4231 | 0x0
	case sbfnr_GALNavMonitor_1: //= 4108
	//case sbfid_GALNavMonitor_1_0: //= 4108 | 0x0
	case sbfnr_INAVmonitor_1: //= 4233
	//case sbfid_INAVmonitor_1_0: //= 4233 | 0x0
	case sbfnr_P2PPStatus_1: //= 4238
	//case sbfid_P2PPStatus_1_0: //= 4238 | 0x0
	case sbfnr_AuthenticationStatus_1: //= 4239
	//case sbfid_AuthenticationStatus_1_0: //= 4239 | 0x0
	case sbfnr_CosmosStatus_1: //= 4243
	//case sbfid_CosmosStatus_1_0: //= 4243 | 0x0

	/* Miscellaneous Blocks */
	case sbfnr_ReceiverSetup_1: //= 5902
	//case sbfid_ReceiverSetup_1_0: //= 5902 | 0x0
	case sbfid_ReceiverSetup_1_1: //= 5902 | 0x2000
	case sbfid_ReceiverSetup_1_2: //= 5902 | 0x4000
	case sbfid_ReceiverSetup_1_3: //= 5902 | 0x6000
	case sbfid_ReceiverSetup_1_4: //= 5902 | 0x8000
	case sbfnr_RxComponents_1: //= 4084
	//case sbfid_RxComponents_1_0: //= 4084 | 0x0
	case sbfnr_RxMessage_1: //= 4103
	//case sbfid_RxMessage_1_0: //= 4103 | 0x0
	case sbfnr_Commands_1: //= 4015
	//case sbfid_Commands_1_0: //= 4015 | 0x0
	case sbfnr_Comment_1: //= 5936
	//case sbfid_Comment_1_0: //= 5936 | 0x0
	case sbfnr_BBSamples_1: //= 4040
	//case sbfid_BBSamples_1_0: //= 4040 | 0x0
	case sbfnr_ASCIIIn_1: //= 4075
	//case sbfid_ASCIIIn_1_0: //= 4075 | 0x0
	case sbfnr_EncapsulatedOutput_1: //= 4097
	//case sbfid_EncapsulatedOutput_1_0: //= 4097 | 0x0
	case sbfnr_RawDataIn_1: //= 4236
	//case sbfid_RawDataIn_1_0: //= 4236 | 0x0

	/* TUR Specific Blocks */
	case sbfnr_TURPVTSatCorrections_1: //= 4035
	//case sbfid_TURPVTSatCorrections_1_0: //= 4035 | 0x0
	case sbfid_TURPVTSatCorrections_1_1: //= 4035 | 0x2000
	case sbfnr_TURHPCAInfo_1: //= 4010
	//case sbfid_TURHPCAInfo_1_0: //= 4010 | 0x0
	case sbfnr_CorrPeakSample_1: //= 4016
	//case sbfid_CorrPeakSample_1_0: //= 4016 | 0x0
	case sbfnr_CorrValues_1: //= 4100
	//case sbfid_CorrValues_1_0: //= 4100 | 0x0
	case sbfnr_TURStatus_1: //= 4041
	//case sbfid_TURStatus_1_0: //= 4041 | 0x0
	case sbfid_TURStatus_1_1: //= 4041 | 0x2000
	case sbfid_TURStatus_1_2: //= 4041 | 0x4000
	case sbfnr_GALIntegrity_1: //= 4033
	//case sbfid_GALIntegrity_1_0: //= 4033 | 0x0
	case sbfnr_TURFormat_1: //= 4080
	//case sbfid_TURFormat_1_0: //= 4080 | 0x0
	case sbfnr_CalibrationValues_1: //= 4215
	//case sbfid_CalibrationValues_1_0: //= 4215 | 0x0
	case sbfnr_MultipathMonitor_1: //= 4220
	//case sbfid_MultipathMonitor_1_0: //= 4220 | 0x0
	case sbfnr_FOCTURNStatus_1: //= 4234
	//case sbfid_FOCTURNStatus_1_0: //= 4234 | 0x0
	case sbfnr_TGVFXStatus_1: //= 4235
	//case sbfid_TGVFXStatus_1_0: //= 4235 | 0x0

	/* PinPoint-GIS RX */
	case sbfnr_GISAction_1: //= 4106
	//case sbfid_GISAction_1_0: //= 4106 | 0x0
	case sbfnr_GISStatus_1: //= 4107
		//case sbfid_GISStatus_1_0: //= 4107 | 0x0
	default:
		log.Printf("[ERROR] handleSbfBlock - Unsupported SBF block ID %d\n", sbfID)
	}
}

func handleRFStatus(buffer []byte) {

}
