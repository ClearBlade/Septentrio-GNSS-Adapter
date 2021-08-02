package sbf

/**
 * Struct definitions of SBF (Septentrio Binary Format) blocks.
 *
 * Adapted from sbfdef.h in the sdk provided in the RxTools suite
 */

//  #if defined(__GNUC__) || defined(__ARMCC__)
//  /* Before the advent of the CPMF platform, double data types were always
// 	* 32-bit aligned. On the CPMF, double data types are 64-bit aligned. The
// 	* "packed, aligned" attribute combination is necessary to enforce 32-bit
// 	* alignment for double data types and to port the sbf encoding/decoding
// 	* functionality to the CPMF.
// 	*/
//  #  define SBFDOUBLE double __attribute__((packed, aligned(4)))
//  #else
//  #  define SBFDOUBLE double
//  #endif
type SBFDOUBLE float64

/*==SBF-BLOCKS definition====================================================*/

/*--Block Header : ----------------------------------------------------------*/
/** Block Header */
type BlockHeader_t struct {
	Sync   uint16
	CRC    uint16
	ID     uint16
	Length uint16
}

/*--Time Header : -----------------------------------------------------------*/
/** Time Header */
type TimeHeader_t struct {
	Header BlockHeader_t
	TOW    uint32
	WNc    uint16
}

//! type for the SBF Header structure
type HeaderBlock_t BlockHeader_t
type HeaderAndTimeBlock_t TimeHeader_t

/*--Header-VoidBlock definition ---------------------------------------------*/
/** \brief type to get a generic SBF block header
*
*  The VoidBlock_t is used in a lot of PPSDK functions to get a generic
*  SBF block. Since the VoidBlock_t has no data associated with it,
*  first a uint8[] buffer should be created that is large enough
*  to contain the full SBF message. Then the buffer should be casted to
*  VoidBlock_t before passing it the PPSDK function. If the size of the
*  SBF block is unknown the buffer can be created with a size of
*  \ref MAX_SBFSIZE so that it can fit any SBF block.
 */
type VoidBlock_t HeaderBlock_t

/*--Receiver Time Header : --------------------------------------------------*/
/** Receiver Time Header */
type ReceiverTimeHeader_t struct {
	Header TimeHeader_t
}

/*--SIS Time Header : -------------------------------------------------------*/
/** SIS Time Header */
type SISTimeHeader_t struct {
	Header TimeHeader_t
}

/*--External Time Header : --------------------------------------------------*/
/** External Time Header */
type ExternalTimeHeader_t struct {
	Header TimeHeader_t
}

/*--Navigation Header : -----------------------------------------------------*/
/** Navigation Header */
type NavigationHeader_t struct {
	Header    SISTimeHeader_t
	SVID      uint8
	CRCPassed uint8
}

/*--MeasEpoch Header : ------------------------------------------------------*/
/** MeasEpoch Header */
type MeasEpochHeader_t struct {
	Header    ReceiverTimeHeader_t
	N1        uint8
	SB1Length uint8
	SB2Length uint8
}

/*==Measurement Blocks=======================================================*/

/** GenMeasEpoch_1_0_t */
/* Measurement set of one epoch */
type GenMeasEpoch_1_0_t struct {
	Header BlockHeader_t

	/* Time Header */
	TOW uint32
	WNc uint16

	/* MeasEpoch Header */
	N       uint8
	SB1Size uint8
	SB2Size uint8

	CommonFlags uint8
	Reserved    [2]uint8
	Data        [SBF_GENMEASEPOCH_1_0_DATA_LENGTH]uint8
}
type GenMeasEpoch_1_t GenMeasEpoch_1_0_t

/** MeasEpochChannelType1_1_0_t */
type MeasEpochChannelType1_1_0_t struct {
	RXChannel  uint8
	Type       uint8
	SVID       uint8
	Misc       uint8
	CodeLSB    uint32
	Doppler    int32
	CarrierLSB uint16
	CarrierMSB int8
	CN0        uint8
	LockTime   uint16
	ObsInfo    uint8
	N_Type2    uint8
}

/** MeasEpochChannelType2_1_0_t */
type MeasEpochChannelType2_1_0_t struct {
	Type             uint8
	LockTime         uint8
	CN0              uint8
	OffsetsMSB       uint8
	CarrierMSB       int8
	ObsInfo          uint8
	CodeOffsetLSB    uint16
	CarrierLSB       uint16
	DopplerOffsetLSB uint16
}

type MeasEpochChannelType1_1_t MeasEpochChannelType1_1_0_t
type MeasEpochChannelType2_1_t MeasEpochChannelType2_1_0_t

/** MeasEpochChannelType2_2_0_t */
type MeasEpochChannelType2_2_0_t struct {
	Type             uint8
	LockTime         uint8
	CN0              uint8
	OffsetsMSB       uint8
	CarrierMSB       int8
	ObsInfo          uint8
	CodeOffsetLSB    uint16
	CarrierLSB       uint16
	DopplerOffsetLSB uint16
}

/** MeasEpochChannelType1_2_0_t */
type MeasEpochChannelType1_2_0_t struct {
	RXChannel  uint8
	Type       uint8
	SVID       uint8
	Misc       uint8
	CodeLSB    uint32
	Doppler    int32
	CarrierLSB uint16
	CarrierMSB int8
	CN0        uint8
	LockTime   uint16
	ObsInfo    uint8
	N_Type2    uint8
}

/** MeasEpoch_2_0_t */
type MeasEpoch_2_0_t struct {
	Header BlockHeader_t

	/* Time Header */
	TOW uint32
	WNc uint16

	/* MeasEpoch Header */
	N       uint8
	SB1Size uint8
	SB2Size uint8

	CommonFlags uint8
	Reserved2   uint8 /* Reserved for future use */
	Reserved    uint8
	Data        [SBF_MEASEPOCH_2_0_DATA_LENGTH]uint8
}

/*--MeasEpoch_2_1_t : -------------------------------------------------------*/
/* Measurement set of one epoch */
/** MeasEpoch_2_1_t */
type MeasEpoch_2_1_t struct {
	Header BlockHeader_t

	/* Time Header */
	TOW uint32
	WNc uint16

	/* MeasEpoch Header */
	N           uint8
	SB1Size     uint8
	SB2Size     uint8
	CommonFlags uint8
	CumClkJumps uint8
	Reserved    uint8
	Data        [SBF_MEASEPOCH_2_1_DATA_LENGTH]uint8
}
type MeasEpoch_2_t MeasEpoch_2_1_t

/** MeasEpochChannelType2_2_1_t */
type MeasEpochChannelType2_2_1_t struct {
	Type             uint8
	LockTime         uint8
	CN0              uint8
	OffsetsMSB       uint8
	CarrierMSB       int8
	ObsInfo          uint8
	CodeOffsetLSB    uint16
	CarrierLSB       uint16
	DopplerOffsetLSB uint16
}

/** MeasEpochChannelType1_2_1_t */
type MeasEpochChannelType1_2_1_t struct {
	RXChannel  uint8
	Type       uint8
	SVID       uint8
	Misc       uint8
	CodeLSB    uint32
	Doppler    int32
	CarrierLSB uint16
	CarrierMSB int8
	CN0        uint8
	LockTime   uint16
	ObsInfo    uint8
	N_Type2    uint8
}

type MeasEpochChannelType2_2_t MeasEpochChannelType2_2_1_t
type MeasEpochChannelType1_2_t MeasEpochChannelType1_2_1_t

/*--MeasExtra_1_0_t : -------------------------------------------------------*/
/* Additional info such as observable variance */

/** MeasExtraChannelSub_1_0_t */
type MeasExtraChannelSub_1_0_t struct {
	RXChannel     uint8
	Type          uint8
	MPCorr        int16
	SmoothingCorr int16
	CodeVar       uint16
	CarrierVar    uint16
	LockTime      uint16
}

/** MeasExtra_1_0_t */
type MeasExtra_1_0_t struct {
	Header BlockHeader_t

	/* Time Header */
	TOW uint32
	WNc uint16

	N                uint8
	SBSize           uint8
	DopplerVarFactor float32
	MeasExtraChannel [SBF_MEASEXTRA_1_0_MEASEXTRACHANNEL_LENGTH]MeasExtraChannelSub_1_0_t
}

/*--MeasExtra_1_1_t : -------------------------------------------------------*/
/* Additional info such as observable variance */

/** MeasExtraChannelSub_1_1_t */
type MeasExtraChannelSub_1_1_t struct {
	RXChannel     uint8
	Type          uint8
	MPCorr        int16
	SmoothingCorr int16
	CodeVar       uint16
	CarrierVar    uint16
	LockTime      uint16
	CumLossCont   uint8
	CarMPCorr     int8
	_padding      [SBF_MEASEXTRACHANNELSUB_1_1__PADDING_LENGTH]uint8
}

/** MeasExtra_1_1_t */
type MeasExtra_1_1_t struct {
	Header BlockHeader_t

	/* Time Header */
	TOW uint32
	WNc uint16

	N                uint8
	SBSize           uint8
	DopplerVarFactor float32
	MeasExtraChannel [SBF_MEASEXTRA_1_1_MEASEXTRACHANNEL_LENGTH]MeasExtraChannelSub_1_1_t
}

/*--MeasExtra_1_2_t : -------------------------------------------------------*/
/* Additional info such as observable variance */

/** MeasExtraChannelSub_1_2_t */
type MeasExtraChannelSub_1_2_t struct {
	RXChannel     uint8
	Type          uint8
	MPCorr        int16
	SmoothingCorr int16
	CodeVar       uint16
	CarrierVar    uint16
	LockTime      uint16
	CumLossCont   uint8
	CarMPCorr     int8
	Info          uint8
	_padding      [SBF_MEASEXTRACHANNELSUB_1_2__PADDING_LENGTH]uint8
}

/** MeasExtra_1_2_t */
type MeasExtra_1_2_t struct {
	Header BlockHeader_t

	/* Time Header */
	TOW uint32
	WNc uint16

	N                uint8
	SBSize           uint8
	DopplerVarFactor float32
	MeasExtraChannel [SBF_MEASEXTRA_1_2_MEASEXTRACHANNEL_LENGTH]MeasExtraChannelSub_1_2_t
}

/*--MeasExtra_1_3_t : -------------------------------------------------------*/
/* Additional info such as observable variance */
/** MeasExtraChannelSub_1_3_t */
type MeasExtraChannelSub_1_3_t struct {
	RXChannel     uint8
	Type          uint8
	MPCorr        int16
	SmoothingCorr int16
	CodeVar       uint16
	CarrierVar    uint16
	LockTime      uint16
	CumLossCont   uint8
	CarMPCorr     int8
	Info          uint8
	Misc          uint8
}

/** MeasExtra_1_3_t */
type MeasExtra_1_3_t struct {
	Header BlockHeader_t

	/* Time Header */
	TOW uint32
	WNc uint16

	N                uint8
	SBSize           uint8
	DopplerVarFactor float32
	MeasExtraChannel [SBF_MEASEXTRA_1_3_MEASEXTRACHANNEL_LENGTH]MeasExtraChannelSub_1_3_t
}
type MeasExtra_1_t MeasExtra_1_3_t
type MeasExtraChannelSub_1_t MeasExtraChannelSub_1_3_t

/*--MeasFullRange_1_0_t : ---------------------------------------------------*/
/* Extended-range code and phase measurements */
/** MeasFullRangeSub_1_0_t */
type MeasFullRangeSub_1_0_t struct {
	RxChannel      uint8
	Type           uint8
	SVID           uint8
	FreqNrAnt      uint8
	CodeObs        SBFDOUBLE
	CarrierMinCode float32
}

/** MeasFullRange_1_0_t */
type MeasFullRange_1_0_t struct {
	Header BlockHeader_t

	/* Time Header */
	TOW uint32
	WNc uint16

	N                uint8
	SBLength         uint8
	Reserved         uint32
	MeasFullRangeSub [SBF_MEASFULLRANGE_1_0_MEASFULLRANGESUB_LENGTH]MeasFullRangeSub_1_0_t
}

/*--MeasFullRange_1_1_t : ---------------------------------------------------*/
/* Extended-range code and phase measurements */
/** MeasFullRangeSub_1_1_t */
type MeasFullRangeSub_1_1_t struct {
	RxChannel      uint8
	Type           uint8
	SVID           uint8
	FreqNrAnt      uint8
	CodeObs        SBFDOUBLE
	CarrierMinCode float32
	CN0            int16
	ObsInfo        uint8
	_padding       [SBF_MEASFULLRANGESUB_1_1__PADDING_LENGTH]uint8
}

/** MeasFullRange_1_1_t */
type MeasFullRange_1_1_t struct {
	Header BlockHeader_t

	/* Time Header */
	TOW uint32
	WNc uint16

	N                uint8
	SBLength         uint8
	Reserved         uint32
	MeasFullRangeSub [SBF_MEASFULLRANGE_1_1_MEASFULLRANGESUB_LENGTH]MeasFullRangeSub_1_1_t
}

type MeasFullRangeSub_1_t MeasFullRangeSub_1_1_t
type MeasFullRange_1_t MeasFullRange_1_1_t

/*--Meas3Ranges_1_0_t : -----------------------------------------------------*/
/* Code, phase and CN0 measurements */
/** M3MasterLong_1_0_t */
type M3MasterLong_1_0_t struct {
	BF1   uint32
	PRLSB uint32
	BF2   uint16
	BF3   uint8
	PRR   int16
}

/** M3MasterShort_1_0_t */
type M3MasterShort_1_0_t struct {
	BF1   uint32
	PRLSB uint32
	PRR   int16
}

/** M3MasterDeltaS_1_0_t */
type M3MasterDeltaS_1_0_t struct {
	BF1 uint32
}

/** M3MasterDeltaL_1_0_t */
type M3MasterDeltaL_1_0_t struct {
	BF1 uint8
	BF2 uint32
}

/** M3SlaveLong_1_0_t */
type M3SlaveLong_1_0_t struct {
	BF1      uint32
	PRLSBrel uint16
	BF3      uint8
}

/** M3SlaveShort_1_0_t */
type M3SlaveShort_1_0_t struct {
	BF1 uint32
	BF2 uint8
}

/** M3SlaveDelta_1_0_t */
type M3SlaveDelta_1_0_t struct {
	BF1      uint16
	dCarrier uint8
}

/** M3SatData_1_0_t */
type M3SatData_1_0_t struct {
	BF1            uint8
	SatMask        [SBF_M3SATDATA_1_0_SATMASK_LENGTH]uint8
	GLOFnList      [SBF_M3SATDATA_1_0_GLOFNLIST_LENGTH]uint8
	BDSLongRange   uint16
	SigExcluded    uint8
	M3MasterLong   M3MasterLong_1_0_t
	M3MasterShort  M3MasterShort_1_0_t
	M3MasterDeltaS M3MasterDeltaS_1_0_t
	M3MasterDeltaL M3MasterDeltaL_1_0_t
	M3SlaveLong    M3SlaveLong_1_0_t
	M3SlaveShort   M3SlaveShort_1_0_t
	M3SlaveDelta   M3SlaveDelta_1_0_t
}

/** Meas3Ranges_1_0_t */
type Meas3Ranges_1_0_t struct {
	Header BlockHeader_t

	/* Time Header */
	TOW uint32
	WNc uint16

	CommonFlags    uint8
	CumClkJumps    uint8
	Constellations uint16
	Misc           uint8
	Reserved       uint8
	Data           [SBF_MEAS3RANGES_1_0_DATA_LENGTH]uint8
}
type M3SatData_1_t M3SatData_1_0_t
type Meas3Ranges_1_t Meas3Ranges_1_0_t

type M3MasterLong_1_t M3MasterLong_1_0_t
type M3MasterShort_1_t M3MasterShort_1_0_t
type M3MasterDeltaS_1_t M3MasterDeltaS_1_0_t
type M3MasterDeltaL_1_t M3MasterDeltaL_1_0_t
type M3SlaveLong_1_t M3SlaveLong_1_0_t
type M3SlaveShort_1_t M3SlaveShort_1_0_t
type M3SlaveDelta_1_t M3SlaveDelta_1_0_t

/*--Meas3CN0HiRes_1_0_t : ---------------------------------------------------*/
/* Extension of Meas3Ranges containing fractional C/N0 values */
/** Meas3CN0HiRes_1_0_t */
type Meas3CN0HiRes_1_0_t struct {
	Header BlockHeader_t

	/* Time Header */
	TOW uint32
	WNc uint16

	Flags    uint16
	CN0HiRes [SBF_MEAS3CN0HIRES_1_0_CN0HIRES_LENGTH]uint8
}

type Meas3CN0HiRes_1_t Meas3CN0HiRes_1_0_t

/*--Meas3Doppler_1_0_t : ----------------------------------------------------*/
/* Extension of Meas3Ranges containing Doppler values */
/** Meas3Doppler_1_0_t */
type Meas3Doppler_1_0_t struct {
	Header BlockHeader_t

	/* Time Header */
	TOW uint32
	WNc uint16

	Flags uint16
	Data  [SBF_MEAS3DOPPLER_1_0_DATA_LENGTH]uint8
}
type Meas3Doppler_1_t Meas3Doppler_1_0_t

/*--Meas3PP_1_0_t : ---------------------------------------------------------*/
/* Extension of Meas3Ranges containing proprietary flags for data post-processing. */
/** Meas3PP_1_0_t */
type Meas3PP_1_0_t struct {
	Header BlockHeader_t

	/* Time Header */
	TOW uint32
	WNc uint16

	Flags uint16
	Data  [SBF_MEAS3PP_1_0_DATA_LENGTH]uint8
}

type Meas3PP_1_t Meas3PP_1_0_t

/*--Meas3MP_1_0_t : ---------------------------------------------------------*/
/* Extension of Meas3Ranges containing multipath corrections applied by the receiver. */
/** Meas3MP_1_0_t */
type Meas3MP_1_0_t struct {
	Header BlockHeader_t

	/* Time Header */
	TOW uint32
	WNc uint16

	Flags uint16
	Data  [SBF_MEAS3MP_1_0_DATA_LENGTH]uint8
}

type Meas3MP_1_t Meas3MP_1_0_t

/*--IQCorr_1_0_t : ----------------------------------------------------------*/
/* Real and imaginary post-correlation values */

/** IQCorrChannelSub_1_0_t */
type IQCorrChannelSub_1_0_t struct {
	RXChannel  uint8
	Type       uint8
	SVID       uint8
	CorrIQ_MSB uint8
	CorrI_LSB  uint8
	CorrQ_LSB  uint8
	Reserved   uint16 /* Reserved for future use */
}

/** IQCorr_1_0_t */
type IQCorr_1_0_t struct {
	Header BlockHeader_t

	/* Time Header */
	TOW uint32
	WNc uint16

	N            uint8
	SBSize       uint8
	CorrDuration uint8
	Reserved2    uint8 /* Reserved for future use */
	Reserved     [2]uint8
	CorrChannel  [SBF_IQCORR_1_0_CORRCHANNEL_LENGTH]IQCorrChannelSub_1_0_t
}

/*--IQCorr_1_1_t : ----------------------------------------------------------*/
/* Real and imaginary post-correlation values */

/** IQCorrChannelSub_1_1_t */
type IQCorrChannelSub_1_1_t struct {
	RXChannel       uint8
	Type            uint8
	SVID            uint8
	CorrIQ_MSB      uint8
	CorrI_LSB       uint8
	CorrQ_LSB       uint8
	CarrierPhaseLSB uint16
}

/** IQCorr_1_1_t */
type IQCorr_1_1_t struct {
	Header BlockHeader_t

	/* Time Header */
	TOW uint32
	WNc uint16

	N            uint8
	SBSize       uint8
	CorrDuration uint8
	CumClkJumps  uint8
	Reserved     [2]uint8
	CorrChannel  [SBF_IQCORR_1_1_CORRCHANNEL_LENGTH]IQCorrChannelSub_1_1_t
}

type IQCorrChannelSub_1_t IQCorrChannelSub_1_1_t
type IQCorr_1_t IQCorr_1_1_t

/*--ISMR_1_0_t : ------------------------------------------------------------*/
/* Ionospheric scintillation monitor (ISMR) data */
/** ISMRChannel_1_0_t */
type ISMRChannel_1_0_t struct {
	RXChannel uint8
	Type      uint8
	SVID      uint8
	Reserved  uint8
	S4        uint16
	SigmaPhi  uint16
}

/** ISMR_1_0_t */
type ISMR_1_0_t struct {
	Header BlockHeader_t

	/* Time Header */
	TOW uint32
	WNc uint16

	N           uint8
	SBLength    uint8
	Reserved    [4]uint8
	ISMRChannel [SBF_ISMR_1_0_ISMRCHANNEL_LENGTH]ISMRChannel_1_0_t
}

type ISMRChannel_1_t ISMRChannel_1_0_t
type ISMR_1_t ISMR_1_0_t

/*--SQMSamples_1_0_t : ------------------------------------------------------*/
/* Correlation samples for signal quality monitoring */
/** SQMOneSample_1_0_t */
type SQMOneSample_1_0_t struct {
	Offset int16
	I      int16
	Q      int16
}

/** SQMChannel_1_0_t */
type SQMChannel_1_0_t struct {
	RxChannel uint8
	Type      uint8
	SVID      uint8
	N2        uint8
}

/** SQMSamples_1_0_t */
type SQMSamples_1_0_t struct {
	Header BlockHeader_t

	/* Time Header */
	TOW uint32
	WNc uint16

	N         uint8
	SB1Length uint8
	SB2Length uint8
	Reserved  [3]uint8
	Data      [SBF_SQMSAMPLES_1_0_DATA_LENGTH]uint8
}

type SQMOneSample_1_t SQMOneSample_1_0_t
type SQMChannel_1_t SQMChannel_1_0_t
type SQMSamples_1_t SQMSamples_1_0_t

/*--EndOfMeas_1_0_t : -------------------------------------------------------*/
/* Measurement epoch marker */
/** EndOfMeas_1_0_t */
type EndOfMeas_1_0_t struct {
	Header BlockHeader_t

	/* Time Header */
	TOW uint32
	WNc uint16
}

type EndOfMeas_1_t EndOfMeas_1_0_t

/*==Navigation Page Blocks===================================================*/

/*--GPSRaw_1_0_t : ----------------------------------------------------------*/
/* GPS CA navigation frame */
/** GPSRaw_1_0_t */
type GPSRaw_1_0_t struct {
	Header BlockHeader_t

	/* Time Header */
	TOW uint32
	WNc uint16

	PRN      uint8
	Reserved uint8
	RawBits  [SBF_GPSRAW_1_0_RAWBITS_LENGTH]uint32
}

type GPSRaw_1_t GPSRaw_1_0_t

/*--CNAVRaw_1_0_t : ---------------------------------------------------------*/
/* GPS L2C navigation frame */
/** CNAVRaw_1_0_t */
type CNAVRaw_1_0_t struct {
	Header BlockHeader_t

	/* Time Header */
	TOW uint32
	WNc uint16

	PRN      uint8
	Reserved uint8
	CNAVBits [SBF_CNAVRAW_1_0_CNAVBITS_LENGTH]uint32
}
type CNAVRaw_1_t CNAVRaw_1_0_t

/*--GEORaw_1_0_t : ----------------------------------------------------------*/
/* SBAS L1 navigation frame */
/** GEORaw_1_0_t */
type GEORaw_1_0_t struct {
	Header BlockHeader_t

	/* Time Header */
	TOW uint32
	WNc uint16

	PRN          uint8
	SignalSource uint8
	RawBits      [SBF_GEORAW_1_0_RAWBITS_LENGTH]uint32
}

type GEORaw_1_t GEORaw_1_0_t

/*--GPSRawCA_1_0_t : --------------------------------------------------------*/
/* GPS CA navigation subframe */
/** GPSRawCA_1_0_t */
type GPSRawCA_1_0_t struct {
	Header BlockHeader_t

	/* Time Header */
	TOW uint32
	WNc uint16

	/* Navigation Header */
	SVID      uint8
	CRCStatus uint8

	ViterbiCnt uint8
	Source     uint8
	FreqNr     uint8
	RxChannel  uint8
	NAVBits    [SBF_GPSRAWCA_1_0_NAVBITS_LENGTH + 6]uint32
}
type GPSRawCA_1_t GPSRawCA_1_0_t

/*--GPSRawL2C_1_0_t : -------------------------------------------------------*/
/* GPS L2C navigation frame */
/** GPSRawL2C_1_0_t */
type GPSRawL2C_1_0_t struct {
	Header BlockHeader_t

	/* Time Header */
	TOW uint32
	WNc uint16

	/* Navigation Header */
	SVID      uint8
	CRCStatus uint8

	ViterbiCnt uint8
	Source     uint8
	FreqNr     uint8
	RxChannel  uint8
	NAVBits    [SBF_GPSRAWL2C_1_0_NAVBITS_LENGTH + 6]uint32
}

type GPSRawL2C_1_t GPSRawL2C_1_0_t

/*--GPSRawL5_1_0_t : --------------------------------------------------------*/
/* GPS L5 navigation frame */
/** GPSRawL5_1_0_t */
type GPSRawL5_1_0_t struct {
	Header BlockHeader_t

	/* Time Header */
	TOW uint32
	WNc uint16

	/* Navigation Header */
	SVID      uint8
	CRCStatus uint8

	ViterbiCnt uint8
	Source     uint8
	FreqNr     uint8
	RxChannel  uint8
	NAVBits    [SBF_GPSRAWL5_1_0_NAVBITS_LENGTH + 6]uint32
}
type GPSRawL5_1_t GPSRawL5_1_0_t

/*--GPSRawL1C_1_0_t : -------------------------------------------------------*/
/* GPS L1C navigation frame */
/** GPSRawL1C_1_0_t */
type GPSRawL1C_1_0_t struct {
	Header BlockHeader_t

	/* Time Header */
	TOW uint32
	WNc uint16

	SVID      uint8
	CRCSF2    uint8
	CRCSF3    uint8
	Source    uint8
	Reserved  uint8
	RxChannel uint8
	NAVBits   [SBF_GPSRAWL1C_1_0_NAVBITS_LENGTH]uint32
}
type GPSRawL1C_1_t GPSRawL1C_1_0_t

/*--GLORawCA_1_0_t : --------------------------------------------------------*/
/* GLONASS CA navigation string */
/** GLORawCA_1_0_t */
type GLORawCA_1_0_t struct {
	Header BlockHeader_t

	/* Time Header */
	TOW uint32
	WNc uint16

	/* Navigation Header */
	SVID      uint8
	CRCStatus uint8

	ViterbiCnt uint8
	Source     uint8
	FreqNr     uint8
	RxChannel  uint8
	NAVBits    [SBF_GLORAWCA_1_0_NAVBITS_LENGTH + 13]uint32
}
type GLORawCA_1_t GLORawCA_1_0_t

/*--GALRawFNAV_1_0_t : ------------------------------------------------------*/
/* Galileo F/NAV navigation page */

/** GALRawFNAV_1_0_t */
type GALRawFNAV_1_0_t struct {
	Header BlockHeader_t

	/* Time Header */
	TOW uint32
	WNc uint16

	/* Navigation Header */
	SVID      uint8
	CRCStatus uint8

	ViterbiCnt uint8
	Source     uint8
	FreqNr     uint8
	RxChannel  uint8
	NAVBits    [SBF_GALRAWFNAV_1_0_NAVBITS_LENGTH + 8]uint32
}
type GALRawFNAV_1_t GALRawFNAV_1_0_t

/*--GALRawINAV_1_0_t : ------------------------------------------------------*/
/* Galileo I/NAV navigation page */
/** GALRawINAV_1_0_t */
type GALRawINAV_1_0_t struct {
	Header BlockHeader_t

	/* Time Header */
	TOW uint32
	WNc uint16

	/* Navigation Header */
	SVID      uint8
	CRCStatus uint8

	ViterbiCnt uint8
	Source     uint8
	FreqNr     uint8
	RxChannel  uint8
	NAVBits    [SBF_GALRAWINAV_1_0_NAVBITS_LENGTH + 8]uint32
}
type GALRawINAV_1_t GALRawINAV_1_0_t

/*--GALRawCNAV_1_0_t : ------------------------------------------------------*/
/* Galileo C/NAV navigation page */
/** GALRawCNAV_1_0_t */
type GALRawCNAV_1_0_t struct {
	Header BlockHeader_t

	/* Time Header */
	TOW uint32
	WNc uint16

	/* Navigation Header */
	SVID      uint8
	CRCStatus uint8

	ViterbiCnt uint8
	Source     uint8
	FreqNr     uint8
	RxChannel  uint8
	NAVBits    [SBF_GALRAWCNAV_1_0_NAVBITS_LENGTH + 0]uint32
}
type GALRawCNAV_1_t GALRawCNAV_1_0_t

/*--GALRawGNAV_1_0_t : ------------------------------------------------------*/
/** GALRawGNAV_1_0_t */
type GALRawGNAV_1_0_t struct {
	Header BlockHeader_t

	/* Time Header */
	TOW uint32
	WNc uint16

	/* Navigation Header */
	SVID      uint8
	CRCStatus uint8

	ViterbiCnt uint8
	Source     uint8
	FreqNr     uint8
	RxChannel  uint8
	NAVBits    [SBF_GALRAWGNAV_1_0_NAVBITS_LENGTH + 10]uint32
}
type GALRawGNAV_1_t GALRawGNAV_1_0_t

/*--GALRawGNAVe_1_0_t : -----------------------------------------------------*/
/** GALRawGNAVe_1_0_t */
type GALRawGNAVe_1_0_t struct {
	Header BlockHeader_t

	/* Time Header */
	TOW uint32
	WNc uint16

	/* Navigation Header */
	SVID      uint8
	CRCStatus uint8

	ViterbiCnt uint8
	Source     uint8
	FreqNr     uint8
	RxChannel  uint8
	NAVBits    [SBF_GALRAWGNAVE_1_0_NAVBITS_LENGTH + 5]uint32
}
type GALRawGNAVe_1_t GALRawGNAVe_1_0_t
type GALRawGNAVencrypted_1_0_t GALRawGNAVe_1_0_t

/*--GEORawL1_1_0_t : --------------------------------------------------------*/
/* SBAS L1 navigation message */
/** GEORawL1_1_0_t */
type GEORawL1_1_0_t struct {
	Header BlockHeader_t

	/* Time Header */
	TOW uint32
	WNc uint16

	/* Navigation Header */
	SVID      uint8
	CRCStatus uint8

	ViterbiCnt uint8
	Source     uint8
	FreqNr     uint8
	RxChannel  uint8
	NAVBits    [SBF_GEORAWL1_1_0_NAVBITS_LENGTH + 8]uint32
}
type GEORawL1_1_t GEORawL1_1_0_t

/*--GEORawL5_1_0_t : --------------------------------------------------------*/
/* SBAS L5 navigation message */
/** GEORawL5_1_0_t */
type GEORawL5_1_0_t struct {
	Header BlockHeader_t

	/* Time Header */
	TOW uint32
	WNc uint16

	/* Navigation Header */
	SVID      uint8
	CRCStatus uint8

	ViterbiCnt uint8
	Source     uint8
	FreqNr     uint8
	RxChannel  uint8
	NAVBits    [SBF_GEORAWL5_1_0_NAVBITS_LENGTH + 8]uint32
}
type GEORawL5_1_t GEORawL5_1_0_t

/*--BDSRaw_1_0_t : ----------------------------------------------------------*/
/* BeiDou navigation page */

/** BDSRaw_1_0_t */
type BDSRaw_1_0_t struct {
	Header BlockHeader_t

	/* Time Header */
	TOW uint32
	WNc uint16

	/* Navigation Header */
	SVID      uint8
	CRCStatus uint8

	ViterbiCnt uint8
	Source     uint8
	Reserved   uint8
	RxChannel  uint8
	NAVBits    [SBF_BDSRAW_1_0_NAVBITS_LENGTH + 6]uint32
}
type BDSRaw_1_t BDSRaw_1_0_t

/*--BDSRawB1C_1_0_t : -------------------------------------------------------*/
/* BeiDou B1C navigation frame */
/** BDSRawB1C_1_0_t */
type BDSRawB1C_1_0_t struct {
	Header BlockHeader_t

	/* Time Header */
	TOW uint32
	WNc uint16

	SVID      uint8
	CRCSF2    uint8
	CRCSF3    uint8
	Source    uint8
	Reserved  uint8
	RxChannel uint8
	NAVBits   [SBF_BDSRAWB1C_1_0_NAVBITS_LENGTH]uint32
}
type BDSRawB1C_1_t BDSRawB1C_1_0_t

/*--BDSRawB2a_1_0_t : -------------------------------------------------------*/
/* BeiDou B2a navigation frame */

/** BDSRawB2a_1_0_t */
type BDSRawB2a_1_0_t struct {
	Header BlockHeader_t

	/* Time Header */
	TOW uint32
	WNc uint16

	/* Navigation Header */
	SVID      uint8
	CRCPassed uint8

	ViterbiCnt uint8
	Source     uint8
	Reserved   uint8
	RxChannel  uint8
	NAVBits    [SBF_BDSRAWB2A_1_0_NAVBITS_LENGTH]uint32
}
type BDSRawB2a_1_t BDSRawB2a_1_0_t

/*--BDSRawB2b_1_0_t : -------------------------------------------------------*/
/* BeiDou B2b navigation frame */
/** BDSRawB2b_1_0_t */
type BDSRawB2b_1_0_t struct {
	Header BlockHeader_t

	/* Time Header */
	TOW uint32
	WNc uint16

	SVID      uint8
	CRCPassed uint8
	Reserved1 uint8
	Source    uint8
	Reserved2 uint8
	RxChannel uint8
	NAVBits   [SBF_BDSRAWB2B_1_0_NAVBITS_LENGTH]uint32
}
type BDSRawB2b_1_t BDSRawB2b_1_0_t

/*--QZSRawL1CA_1_0_t : ------------------------------------------------------*/
/* QZSS L1 CA navigation frame */

/** QZSRawL1CA_1_0_t */
type QZSRawL1CA_1_0_t struct {
	Header BlockHeader_t

	/* Time Header */
	TOW uint32
	WNc uint16

	/* Navigation Header */
	SVID      uint8
	CRCPassed uint8

	Reserved  uint8
	Source    uint8
	Reserved2 uint8
	RxChannel uint8
	NAVBits   [SBF_QZSRAWL1CA_1_0_NAVBITS_LENGTH]uint32
}
type QZSRawL1CA_1_t QZSRawL1CA_1_0_t

/*--QZSRawL2C_1_0_t : -------------------------------------------------------*/
/* QZSS L2C navigation frame */

/** QZSRawL2C_1_0_t */
type QZSRawL2C_1_0_t struct {
	Header BlockHeader_t

	/* Time Header */
	TOW uint32
	WNc uint16

	/* Navigation Header */
	SVID      uint8
	CRCPassed uint8

	ViterbiCnt uint8
	Source     uint8
	Reserved   uint8
	RxChannel  uint8
	NAVBits    [SBF_QZSRAWL2C_1_0_NAVBITS_LENGTH]uint32
}
type QZSRawL2C_1_t QZSRawL2C_1_0_t

/*--QZSRawL5_1_0_t : --------------------------------------------------------*/
/* QZSS L5 navigation frame */
/** QZSRawL5_1_0_t */
type QZSRawL5_1_0_t struct {
	Header BlockHeader_t

	/* Time Header */
	TOW uint32
	WNc uint16

	/* Navigation Header */
	SVID      uint8
	CRCPassed uint8

	ViterbiCnt uint8
	Source     uint8
	Reserved   uint8
	RxChannel  uint8
	NAVBits    [SBF_QZSRAWL5_1_0_NAVBITS_LENGTH]uint32
}
type QZSRawL5_1_t QZSRawL5_1_0_t

/*--QZSRawL6_1_0_t : --------------------------------------------------------*/
/* QZSS L6 navigation message */

/** QZSRawL6_1_0_t */
type QZSRawL6_1_0_t struct {
	Header BlockHeader_t

	/* Time Header */
	TOW uint32
	WNc uint16

	SVID      uint8
	Parity    uint8
	RSCnt     uint8
	Source    uint8
	Reserved  uint8
	RxChannel uint8
	NAVBits   [SBF_QZSRAWL6_1_0_NAVBITS_LENGTH]uint32
}
type QZSRawL6_1_t QZSRawL6_1_0_t

/*--QZSRawL1C_1_0_t : -------------------------------------------------------*/
/* QZSS L1C navigation frame */

/** QZSRawL1C_1_0_t */
type QZSRawL1C_1_0_t struct {
	Header BlockHeader_t

	/* Time Header */
	TOW uint32
	WNc uint16

	SVID      uint8
	CRCSF2    uint8
	CRCSF3    uint8
	Source    uint8
	Reserved  uint8
	RxChannel uint8
	NAVBits   [SBF_QZSRAWL1C_1_0_NAVBITS_LENGTH]uint32
}
type QZSRawL1C_1_t QZSRawL1C_1_0_t

/*--QZSRawL1S_1_0_t : -------------------------------------------------------*/
/* QZSS L1S navigation message */
/** QZSRawL1S_1_0_t */
type QZSRawL1S_1_0_t struct {
	Header BlockHeader_t

	/* Time Header */
	TOW uint32
	WNc uint16

	/* Navigation Header */
	SVID      uint8
	CRCPassed uint8

	ViterbiCnt uint8
	Source     uint8
	FreqNr     uint8
	RxChannel  uint8
	NAVBits    [SBF_QZSRAWL1S_1_0_NAVBITS_LENGTH]uint32
}
type QZSRawL1S_1_t QZSRawL1S_1_0_t

/*--NAVICRaw_1_0_t : --------------------------------------------------------*/
/* NavIC/IRNSS subframe */
/** NAVICRaw_1_0_t */
type NAVICRaw_1_0_t struct {
	Header BlockHeader_t

	/* Time Header */
	TOW uint32
	WNc uint16

	/* Navigation Header */
	SVID      uint8
	CRCPassed uint8

	ViterbiCnt uint8
	Source     uint8
	Reserved   uint8
	RxChannel  uint8
	NAVBits    [SBF_NAVICRAW_1_0_NAVBITS_LENGTH]uint32
}
type NAVICRaw_1_t NAVICRaw_1_0_t

/*--GNSSNavBits_1_0_t : -----------------------------------------------------*/
/* Raw navigation bits during last second */
/** RawNavBits_1_0_t */
type RawNavBits_1_0_t struct {
	/* Time Header */
	TOW uint32
	WNc uint16

	Type      uint8
	SVID      uint8
	FreqNr    uint8
	RxChannel uint8
	NBits     uint16
	NavBits   [SBF_RAWNAVBITS_1_0_NAVBITS_LENGTH]uint32
}

/** GNSSNavBits_1_0_t */
type GNSSNavBits_1_0_t struct {
	Header BlockHeader_t

	RawNavBits RawNavBits_1_0_t
}
type RawNavBits_1_t RawNavBits_1_0_t
type GNSSNavBits_1_t GNSSNavBits_1_0_t

/*--GNSSSymbols_1_0_t : -----------------------------------------------------*/
/* Raw navigation symbols */
/** RawSymbols_1_0_t */
type RawSymbols_1_0_t struct {
	/* Time Header */
	TOW uint32
	WNc uint16

	Type      uint8
	SVID      uint8
	FreqNr    uint8
	RxChannel uint8
	NSymbols  uint16
	Symbols   [SBF_RAWSYMBOLS_1_0_SYMBOLS_LENGTH]uint32
}

/** GNSSSymbols_1_0_t */
type GNSSSymbols_1_0_t struct {
	Header     BlockHeader_t
	RawSymbols RawSymbols_1_0_t
}
type RawSymbols_1_t RawSymbols_1_0_t
type GNSSSymbols_1_t GNSSSymbols_1_0_t

/*==GPS Decoded Message Blocks===============================================*/

/*--GPSNav_1_0_t : ----------------------------------------------------------*/
/* GPS ephemeris and clock */

/** gpEph_1_0_t */
type gpEph_1_0_t struct {
	/* Time Header */
	TOW uint32
	WNc uint16

	PRN        uint8
	Reserved   uint8
	WN         uint16    /* Week number (modulo 1024). [061-070:1] */
	CAorPonL2  uint8     /* C/A code or P code on L2. [071-072:1] */
	URA        uint8     /* User range accuracy index. [073-076:1] */
	health     uint8     /* SV health. [077-082:1] */
	L2DataFlag uint8     /* L2 P data flag. [091-091:1] */
	IODC       uint16    /* Issue of data clock. [083-084:1] [211-218:1] */
	IODE2      uint8     /* Issue of data eph. frame 2. [061-068:2] */
	IODE3      uint8     /* Issue of data eph. frame 3. [271-278:3] */
	FitIntFlg  uint8     /* fit interval flag [287-287:2] */
	dummy      uint8     /* introduced w.r.t. 32-bits memory alignment */
	T_gd       float32   /* Correction term T_gd (s). [197-204:1] */
	t_oc       uint32    /* Clock correction t_oc (s). [219-234:1] */
	a_f2       float32   /* Clock correction a_f2 (s/s^2). [241-248:1] */
	a_f1       float32   /* Clock correction a_f1 (s/s). [249-264:1] */
	a_f0       float32   /* Clock correction a_f0 (s). [271-292:1] */
	C_rs       float32   /* radius sin ampl (m) [069-084:2] */
	DEL_N      float32   /* mean motion diff (semi-circles/s) [091-106:2] */
	M_0        SBFDOUBLE /* Mean Anom (semi-circles) [107-114:2] [121-144:2] */
	C_uc       float32   /* lat cosine ampl (r) [151-166:2] */
	e          SBFDOUBLE /* Eccentricity [167-174:2] [181-204:2] */
	C_us       float32   /* Lat sine ampl   (r) [211-226:2] */
	SQRT_A     SBFDOUBLE /* SQRT(A) (m^1/2) [227-234:2 241-264:2] */
	t_oe       uint32    /* Reference time of ephemeris (s) [271-286:2] */
	C_ic       float32   /* inclin cos ampl (r) [061-076:3] */
	OMEGA_0    SBFDOUBLE /* Right Ascen at TOA (semi-circles) [077-084:3] [091-114:3] */
	C_is       float32   /* inclin sin ampl (r) [121-136:3] */
	i_0        SBFDOUBLE /* Orbital Inclination (semi-circles) [137-144:3] [151-174:3] */
	C_rc       float32   /* radius cos ampl (m) [181-196:3] */
	omega      SBFDOUBLE /* Argument of Perigee(semi-circle) [197-204:3] [211-234:3] */
	OMEGADOT   float32   /* Rate of Right Ascen(semi-circles/s) [241-264:3] */
	IDOT       float32   /* Rate of inclin (semi-circles/s) [279-292:3] */
	WNt_oc     uint16    /* modified WN to go with t_oc (still modulo 1024) */
	WNt_oe     uint16    /* modified WN to go with t_oe (still modulo 1024) */
}

/** GPSNav_1_0_t */
type GPSNav_1_0_t struct {
	Header BlockHeader_t

	Eph gpEph_1_0_t
}

type gpEph_1_t gpEph_1_0_t

type GPSNav_1_t GPSNav_1_0_t

/*--GPSAlm_1_0_t : ----------------------------------------------------------*/
/* Almanac data for a GPS satellite */
/** gpAlm_1_0_t */
type gpAlm_1_0_t struct {
	/* Time Header */
	TOW uint32
	WNc uint16

	PRN      uint8
	Reserved uint8
	e        float32 /* Eccentricity */
	t_oa     uint32  /* Reference time of almanac (s) */
	delta_i  float32 /* Orbital Inclination (semi-circles) */
	OMEGADOT float32 /* Rate of Right Ascen (r/s) */
	SQRT_A   float32 /* SQRT(A) (m^1/2) */
	OMEGA_0  float32 /* Right Ascen at TOA (semi-circles) */
	omega    float32 /* Argument of Perigee (semi-circles) */
	M_0      float32 /* Mean Anom (semi-circles) */
	a_f1     float32 /* Clock correction a_f1 (s/s). */
	a_f0     float32 /* Clock correction a_f0 (s). */
	WN_a     uint8   /* Reference week number of almanac. */
	config   uint8   /* See GPS SPS Sig. Spec. 2nd Edition 1995 2.4.5.4 p31.
	-bit 00-03: 4bit anti spoofing & configuration. */
	health8 uint8 /* -bit 00-07: 8bit health data in alm. subframes. */
	health6 uint8 /* -bit 08-13: 6bit health data SF4P25 and SF5P25. */
}

/** GPSAlm_1_0_t */
type GPSAlm_1_0_t struct {
	Header BlockHeader_t
	Alm    gpAlm_1_0_t
}
type gpAlm_1_t gpAlm_1_0_t
type GPSAlm_1_t GPSAlm_1_0_t

/*--GPSIon_1_0_t : ----------------------------------------------------------*/
/* Ionosphere data from the GPS subframe 5 */

/** gpIon_1_0_t */
type gpIon_1_0_t struct {
	/* Time Header */
	TOW uint32
	WNc uint16

	PRN      uint8
	Reserved uint8
	alpha_0  float32 /* (sec) [069-076:4p18] */
	alpha_1  float32 /* (sec/semicircle) [077-084:4p18] */
	alpha_2  float32 /* (sec/semicircle^2) [091-098:4p18] */
	alpha_3  float32 /* (sec/semicircle^3) [099-106:4p18] */
	beta_0   float32 /* (sec) [107-114:4p18] */
	beta_1   float32 /* (sec/semicircle) [121-128:4p18] */
	beta_2   float32 /* (sec/semicircle^2) [129-136:4p18] */
	beta_3   float32 /* (sec/semicircle^3) [137-144:4p18] */
}

/** GPSIon_1_0_t */
type GPSIon_1_0_t struct {
	Header BlockHeader_t

	Ion gpIon_1_0_t
}
type gpIon_1_t gpIon_1_0_t
type GPSIon_1_t GPSIon_1_0_t

/*--GPSUtc_1_0_t : ----------------------------------------------------------*/
/* GPS-UTC data from GPS subframe 5 */

/** gpUtc_1_0_t */
type gpUtc_1_0_t struct {
	/* Time Header */
	TOW uint32
	WNc uint16

	PRN       uint8
	Reserved  uint8
	A_1       float32   /* (sec/sec) [151-176:4p18] */
	A_0       SBFDOUBLE /* (sec) [181-194:4] [211-218:4] */
	t_ot      uint32    /* (sec) [219-226:4p18] */
	WN_t      uint8     /* (wk) [227-234:4p18] */
	DEL_t_LS  int8      /* (sec) [241-218:4p18] */
	WN_LSF    uint8     /* (wk) [219-226:4p18] */
	DN        uint8     /* (days) [227-234:4p18] */
	DEL_t_LSF int8      /* (sec) [271-278:4p18] */
}

/** GPSUtc_1_0_t */
type GPSUtc_1_0_t struct {
	Header BlockHeader_t
	Utc    gpUtc_1_0_t
}

type gpUtc_1_t gpUtc_1_0_t
type GPSUtc_1_t GPSUtc_1_0_t

/*--GPSCNav_1_0_t : ---------------------------------------------------------*/
/* CNAV Ephemeris data for one satellite. */
/** gpEphCNAV_1_0_t */
type gpEphCNAV_1_0_t struct {
	/* Time Header */
	TOW uint32
	WNc uint16

	PRN         uint8 /* PRN [M10/11] */
	Alert       uint8
	WN          uint16    /* WN [MT10] */
	Health      uint8     /* Health [MT10] */
	URA_ED      int8      /* ED Accuracy Index [MT10] */
	t_op        uint32    /* Data Predict TOW (s) [MT10/30] */
	t_oe        uint32    /* Ephemeris Reference TOW (s) [MT10/11] */
	A           SBFDOUBLE /* Semi Major Axis (m) [MT10] */
	A_DOT       SBFDOUBLE /* Semi Major Axis Change Rate (m/s) [MT10] */
	DELTA_N     float32   /* Mean Motion Delta (semi-circles/s) [MT10] */
	DELTA_N_DOT float32   /* Mean Motion Rate Delta (semi-circles/s^2) [MT10] */
	M_0         SBFDOUBLE /* Mean anomaly (semi-circles) [MT10] */
	e           SBFDOUBLE /* Eccentricity [MT10] */
	omega       SBFDOUBLE /* Argument of Perigee (semi-circles) [MT10] */
	OMEGA_0     SBFDOUBLE /* Reference Right Ascension Angle (semi-circles) [MT11] */
	OMEGADOT    float32   /* Right Ascension Rate (semi-circles/s) [MT11] */
	i_0         SBFDOUBLE /* Inclination Angle (semi-circles) [MT11] */
	IDOT        float32   /* Inclination Angle Rate (semi-circles/s) [MT11] */
	C_is        float32   /* Amplitude is (r) [MT11] */
	C_ic        float32   /* Amplitude ic (r) [MT11] */
	C_rs        float32   /* Amplitude rs (m) [MT11] */
	C_rc        float32   /* Amplitude rc (m) [MT11] */
	C_us        float32   /* Amplitude us (r) [MT11] */
	C_uc        float32   /* Amplitude uc (r) [MT11] */
	t_oc        uint32    /* Clock Reference Time (s) [MT30] */
	URA_NED0    int8      /* NED Accuracy Index [MT30] */
	URA_NED1    uint8     /* NED Accuracy Change Index [MT30] */
	URA_NED2    uint8     /* NED Accuracy Change Rate Index [MT30] */
	WN_op       uint8     /* Data Predict WN [MT30] */
	a_f2        float32   /* Clock Drift Rate Correction Coefficient (s/s^2) [MT30] */
	a_f1        float32   /* Clock Drift Correction Coefficient (s/s) [MT30] */
	a_f0        SBFDOUBLE /* Clock Bias Correction Coefficient (s) [MT30] */
	T_gd        float32   /* Tgd (s) [MT30] */
	ISC_L1CA    float32   /* ISC L1CA (s) [MT30] */
	ISC_L2C     float32   /* ISC L2C (s) [MT30] */
	ISC_L5I5    float32   /* ISC L5I5 (s) [MT30] */
	ISC_L5Q5    float32   /* ISC L5Q5 (s) [MT30] */
}

/** GPSCNav_1_0_t */
type GPSCNav_1_0_t struct {
	Header BlockHeader_t
	Eph    gpEphCNAV_1_0_t
}

type gpEphCNAV_1_t gpEphCNAV_1_0_t
type GPSCNav_1_t GPSCNav_1_0_t

/*==GLONASS Decoded Message Blocks===========================================*/

/*--GLONav_1_0_t : ----------------------------------------------------------*/
/* GLONASS ephemeris and clock */

/** glEph_1_0_t */
type glEph_1_0_t struct {
	/* Time Header */
	TOW uint32
	WNc uint16

	SVID   uint8     /* Slot number + 37 */
	FreqNr uint8     /* Frequency number + 8 */
	x      SBFDOUBLE /* X component of satellite position in PZ-90 (km) */
	y      SBFDOUBLE /* Y component of satellite position in PZ-90 (km) */
	z      SBFDOUBLE /* Z component of satellite position in PZ-90 (km) */
	dx     float32   /* X component of satellite velocity in PZ-90 (km/s) */
	dy     float32   /* Y component of satellite velocity in PZ-90 (km/s) */
	dz     float32   /* Z component of satellite velocity in PZ-90 (km/s) */
	ddx    float32   /* X component of satellite acceleration in PZ-90 (km/s2) */
	ddy    float32   /* Y component of satellite acceleration in PZ-90 (km/s2) */
	ddz    float32   /* Z component of satellite acceleration in PZ-90 (km/s2) */
	gamma  float32   /* relative deviation of predicted carrier frequency */
	tau    float32   /* corr to nth satellite time rel to GLONASS time t_c */
	dtau   float32   /* time difference between L2 and L1 sub-band (s) */
	t_oe   uint32    /* reference time of GLONASS ephemeris in GPS time frame */
	WNt_oe uint16    /* reference WN of GLONASS ephemeris in GPS time frame */
	P1     uint8     /* length of applicability interval (min) */
	P2     uint8     /* odd/even flag of t_b */
	E      uint8     /* age of immediate data (days) */
	B      uint8     /* health flag, unhealthy if MSB is set */
	tb     uint16    /* time of day [min] defining middle of validity interval */
	M      uint8     /* GLONASS-M only: GLONASS-M satellite identifier (01, otherwise 00) */
	P      uint8     /* GLONASS-M only: mode of computation of freq/time corr data */
	l      uint8     /* GLONASS-M only: health flag, 0=healthy, 1=unhealthy */
	P4     uint8     /* GLONASS-M only: 'updated' flag of ephemeris data */
	N_T    uint16    /* GLONASS-M only: current day within 4-year interval */
	F_T    uint16    /* GLONASS-M only: predicted user range accuracy at time t_b */
}

/** GLONav_1_0_t */
type GLONav_1_0_t struct {
	Header BlockHeader_t
	Eph    glEph_1_0_t
}

type glEph_1_t glEph_1_0_t
type GLONav_1_t GLONav_1_0_t

/*--GLOAlm_1_0_t : ----------------------------------------------------------*/
/* Almanac data for a GLONASS satellite */
/** glAlm_1_0_t */
type glAlm_1_0_t struct {
	/* Time Header */
	TOW uint32
	WNc uint16

	SVID     uint8   /* Slot number + 37 */
	FreqNr   uint8   /* Frequency number + 8 */
	epsilon  float32 /* eccentricity of satellite at t_ln */
	t_oa     uint32  /* Reference time of almanac in GPS time frame (TOW) */
	Delta_i  float32 /* corr to mean value of incl at t_ln (semi-circle) */
	lambda   float32 /* longitude of first ascending node (semi-circle) */
	t_ln     float32 /* time of first ascending node passage (s) */
	omega    float32 /* argument of perigee at t_ln (semi-circle) */
	Delta_T  float32 /* corr to mean value of T_dr at t_ln (s/orb period) */
	dDelta_T float32 /* rate of change of T_dr (s/orb period^2) */
	tau      float32 /* coarse value of satellite time corr at t_ln (s) */
	WN_a     uint8   /* Reference week number of almanac in GPS time frame %256 */
	C        uint8   /* general health flag (1 indicates healthy) */
	N        uint16  /* calender day number within 4 year period */
	M        uint8   /* GLONASS-M only: GLONASS-M satellite identifier (01, otherwise 00) */
	N_4      uint8   /* GLONASS-M only: 4 year interval number, starting from 1996 */
	_padding [SBF_GLALM_1_0__PADDING_LENGTH]uint8
}

/** GLOAlm_1_0_t */
type GLOAlm_1_0_t struct {
	Header BlockHeader_t
	Alm    glAlm_1_0_t
}

type glAlm_1_t glAlm_1_0_t
type GLOAlm_1_t GLOAlm_1_0_t

/*--GLOTime_1_0_t : ---------------------------------------------------------*/
/* GLO-UTC, GLO-GPS and GLO-UT1 data */

/** glTime_1_0_t */
type glTime_1_0_t struct {
	/* Time Header */
	TOW uint32
	WNc uint16

	SVID    uint8 /* Slot number + 37 */
	FreqNr  uint8 /* Frequency number + 8 */
	N_4     uint8
	KP      uint8 /* notification of leap second */
	N       uint16
	tau_GPS float32
	tau_c   SBFDOUBLE
	B1      float32 /* diff between UT1 and UTC at start of curr day(s) */
	B2      float32 /* daily change to diff delta UT1 (s/msd) */
}

/** GLOTime_1_0_t */
type GLOTime_1_0_t struct {
	Header  BlockHeader_t
	GLOTime glTime_1_0_t
}

type glTime_1_t glTime_1_0_t
type GLOTime_1_t GLOTime_1_0_t

/*==Galileo Decoded Message Blocks===========================================*/

/*--GALNav_1_0_t : ----------------------------------------------------------*/
/* Galileo ephemeris, clock, health and BGD */

/** gaEph_1_0_t */
type gaEph_1_0_t struct {
	/* Time Header */
	TOW uint32
	WNc uint16

	SVID         uint8     /* SBF range 71-102 SIS range 1-64 F1/I4/G1 */
	Source       uint8     /* obtained from INAV (2) or FNAV (16) */
	SQRT_A       SBFDOUBLE /* SQRT(A) (m^1/2) [F2/I1/G1] */
	M_0          SBFDOUBLE /* Mean Anom (semi-circles) [F2/I1/G1] */
	e            SBFDOUBLE /* Eccentricity [F2/I1/G1] */
	i_0          SBFDOUBLE /* Orbital Inclination (semi-circles) [F3/I2/G2] */
	omega        SBFDOUBLE /* Argument of Perigee (semi-circles) [F3/I2/G2] */
	OMEGA_0      SBFDOUBLE /* Right Ascen at TOA (semi-circles) [F2/I2/G2] */
	OMEGADOT     float32   /* Rate of Right Ascen(semi-circles/s) [F2/I3/G1] */
	IDOT         float32   /* Rate of inclin (semi-circles/s) [F2/I2/G1] */
	DEL_N        float32   /* mean motion diff (semi-circles/s) [F3/I3/G2] */
	C_uc         float32   /* lat cosine ampl (rad) [F3/I3/G2] */
	C_us         float32   /* Lat sine ampl   (rad) [F3/I3/G3] */
	C_rc         float32   /* radius cos ampl (m) [F3/I3/G3] */
	C_rs         float32   /* radius sin ampl (m) [F3/I3/G3] */
	C_ic         float32   /* inclin cos ampl (r) [F4/I4/G3] */
	C_is         float32   /* inclin sin ampl (r) [F4/I4/G3] */
	t_oe         uint32    /* Reference time of ephemeris (s) [F3/I1/G2] */
	t_oc         uint32    /* Clock correction t_oc (s) [F1/I4/G1] */
	a_f2         float32   /* Clock correction a_f2 (s/s^2) [F1/I4/G1] */
	a_f1         float32   /* Clock correction a_f1 (s/s) [F1/I4/G1] */
	a_f0         SBFDOUBLE /* Clock correction a_f0 (s) [F1/I4/G1] */
	WNt_oe       uint16    /* modified WN to go with t_oe */
	WNt_oc       uint16    /* modified WN to go with t_oc */
	IODnav       uint16    /* 0 - 1023 [Fx/Ix/Gx] */
	Health_OSSOL uint16    /* contains DVS,HS for OS an SOL */
	Health_PRS   uint8     /* contains DVS,HS for PRS */
	SISA_L1E5a   uint8     /* Signal In Space Accuracy [F1] */
	SISA_L1E5b   uint8     /* Signal In Space Accuracy [I3] */
	SISA_L1AE6A  uint8     /* Signal In Space Accuracy [G3] */
	BGD_L1E5a    float32   /* Correction term T_gd (s). [F1/I5] */
	BGD_L1E5b    float32   /* Correction term T_gd (s). [I5] */
	BGD_L1AE6A   float32   /* Correction term T_gd (s). [G4] */
	CNAVenc      uint8     /* C/NAV encryption status [I1] */
}

/** GALNav_1_0_t */
type GALNav_1_0_t struct {
	Header BlockHeader_t
	Eph    gaEph_1_0_t
}

type gaEph_1_t gaEph_1_0_t
type GALNav_1_t GALNav_1_0_t

/*--GALAlm_1_0_t : ----------------------------------------------------------*/
/* Almanac data for a Galileo satellite */

/** gaAlm_1_0_t */
type gaAlm_1_0_t struct {
	/* Time Header */
	TOW uint32
	WNc uint16

	SVID     uint8   /* SBF range 71-102 SIS range 1-64 F1/I4/G1 */
	Source   uint8   /* bitfield according to sigType */
	e        float32 /* Eccentricity */
	t_oa     uint32  /* Reference time of almanac (s) */
	delta_i  float32 /* Orbital Inclination (semi-circles) */
	OMEGADOT float32 /* Rate of Right Ascen (r/s) */
	SQRT_A   float32 /* SQRT(A) (m^1/2) */
	OMEGA_0  float32 /* Right Ascen at TOA (semi-circles) */
	omega    float32 /* Argument of Perigee (semi-circles) */
	M_0      float32 /* Mean Anom (semi-circles) */
	a_f1     float32 /* Clock correction a_f1 (s/s). */
	a_f0     float32 /* Clock correction a_f0 (s). */
	WN_a     uint8   /* Reference week number of almanac. */
	SVID_A   uint8
	health   uint16
	IODa     uint8
}

/** GALAlm_1_0_t */
type GALAlm_1_0_t struct {
	Header BlockHeader_t
	Alm    gaAlm_1_0_t
}

type gaAlm_1_t gaAlm_1_0_t
type GALAlm_1_t GALAlm_1_0_t

/*--GALIon_1_0_t : ----------------------------------------------------------*/
/* NeQuick Ionosphere model parameters */

/** gaIon_1_0_t */
type gaIon_1_0_t struct {
	/* Time Header */
	TOW uint32
	WNc uint16

	SVID       uint8   /* SBF range 71-102 SIS range 1-64 F1/I4/G1 */
	Source     uint8   /* bitfield according to sigType */
	a_i0       float32 /* (sec) [F1/I5/C4] */
	a_i1       float32 /* (sec/semicircle) [F1/I5/C4] */
	a_i2       float32 /* (sec/semicircle^2) [F1/I5/C4] */
	StormFlags uint8   /* 5 bits: 1 for each region [F1/I5/C4] */
}

/** GALIon_1_0_t */
type GALIon_1_0_t struct {
	Header BlockHeader_t
	Ion    gaIon_1_0_t
}

type gaIon_1_t gaIon_1_0_t
type GALIon_1_t GALIon_1_0_t

/*--GALUtc_1_0_t : ----------------------------------------------------------*/
/* GST-UTC data */

/** gaUtc_1_0_t */
type gaUtc_1_0_t struct {
	/* Time Header */
	TOW uint32
	WNc uint16

	SVID      uint8     /* SBF range 71-102 SIS range 1-64 [F1/I4/G1] */
	Source    uint8     /* bitfield according to sigType */
	A_1       float32   /* (sec/sec) [F4/I6/C5] */
	A_0       SBFDOUBLE /* (sec) [F4/I6/C5] */
	t_ot      uint32    /* (sec) [F4/I6/C5] */
	WN_ot     uint8     /* (wk) [F4/I6/C5] */
	DEL_t_LS  int8      /* (sec) [F4/I6/C5] */
	WN_LSF    uint8     /* (wk) [F4/I6/C5] */
	DN        uint8     /* (days) [F4/I6/C5] */
	DEL_t_LSF int8      /* (sec) [F4/I6/C5] */
}

/** GALUtc_1_0_t */
type GALUtc_1_0_t struct {
	Header BlockHeader_t
	Utc    gaUtc_1_0_t
}

type gaUtc_1_t gaUtc_1_0_t
type GALUtc_1_t GALUtc_1_0_t

/*--GALGstGps_1_0_t : -------------------------------------------------------*/
/* GST-GPS data */

/** gaGstGps_1_0_t */
type gaGstGps_1_0_t struct {
	/* Time Header */
	TOW uint32
	WNc uint16

	SVID   uint8   /* SBF range 71-102 SIS range 1-64 [F1/I4/G1] */
	Source uint8   /* bitfield according to sigType */
	A_1G   float32 /* (sec/sec) [F4/I10/C5] */
	A_0G   float32 /* (sec) [F4/I10/C5] */
	t_oG   uint32  /* (sec) [F4/I10/C5] */
	WN_oG  uint8   /* (wk) [F4/I10/C5] */
}

/** GALGstGps_1_0_t */
type GALGstGps_1_0_t struct {
	Header BlockHeader_t
	GstGps gaGstGps_1_0_t
}

type gaGstGps_1_t gaGstGps_1_0_t
type GALGstGps_1_t GALGstGps_1_0_t

/*--GALSARRLM_1_0_t : -------------------------------------------------------*/
/* Search-and-rescue return link message */

/** gaSARRLM_1_0_t */
type gaSARRLM_1_0_t struct {
	/* Time Header */
	TOW uint32
	WNc uint16

	SVID      uint8
	Source    uint8 /* bitfield according to sigType */
	RLMLength uint8
	Reserved  [3]uint8
	RLMBits   [SBF_GASARRLM_1_0_RLMBITS_LENGTH]uint32
}

/** GALSARRLM_1_0_t */
type GALSARRLM_1_0_t struct {
	Header    BlockHeader_t
	galSARRLM gaSARRLM_1_0_t
}

type gaSARRLM_1_t gaSARRLM_1_0_t
type GALSARRLM_1_t GALSARRLM_1_0_t

/*==BeiDou Decoded Message Blocks============================================*/

/*--BDSNav_1_0_t : ----------------------------------------------------------*/
/* BeiDou ephemeris and clock */

/** cmpEph_1_0_t */
type cmpEph_1_0_t struct {
	/* Time Header */
	TOW uint32
	WNc uint16

	PRN       uint8
	Reserved  uint8
	WN        uint16
	URA       uint8
	SatH1     uint8
	IODC      uint8
	IODE      uint8
	Reserved2 uint16
	T_GD1     float32
	T_GD2     float32
	t_oc      uint32
	a_f2      float32
	a_f1      float32
	a_f0      float32
	C_rs      float32
	DEL_N     float32
	M_0       SBFDOUBLE
	C_uc      float32
	e         SBFDOUBLE
	C_us      float32
	SQRT_A    SBFDOUBLE
	t_oe      uint32
	C_ic      float32
	OMEGA_0   SBFDOUBLE
	C_is      float32
	i_0       SBFDOUBLE
	C_rc      float32
	omega     SBFDOUBLE
	OMEGADOT  float32
	IDOT      float32
	WNt_oc    uint16
	WNt_oe    uint16
}

/** BDSNav_1_0_t */
type BDSNav_1_0_t struct {
	Header BlockHeader_t
	Eph    cmpEph_1_0_t
}

type cmpEph_1_t cmpEph_1_0_t
type BDSNav_1_t BDSNav_1_0_t

/*--BDSAlm_1_0_t : ----------------------------------------------------------*/
/* Almanac data for a BeiDou satellite */

/** cmpAlm_1_0_t */
type cmpAlm_1_0_t struct {
	/* Time Header */
	TOW uint32
	WNc uint16

	PRN      uint8
	WN_a     uint8   /* Reference week number of almanac. */
	t_oa     uint32  /* Reference time of almanac (s) */
	SQRT_A   float32 /* SQRT(A) (m^1/2) */
	e        float32 /* Eccentricity */
	omega    float32 /* Argument of Perigee (semi-circles) */
	M_0      float32 /* Mean Anom (semi-circles) */
	OMEGA_0  float32 /* Right Ascen at TOA (semi-circles) */
	OMEGADOT float32 /* Rate of Right Ascen (r/s) */
	delta_i  float32 /* Orbital Inclination (semi-circles) */
	a_f0     float32 /* Clock correction a_0 (s). */
	a_f1     float32 /* Clock correction a_1 (s/s). */
	Health   uint16  /* Health. */
	Reserved [2]uint8
}

/** BDSAlm_1_0_t */
type BDSAlm_1_0_t struct {
	Header BlockHeader_t
	Alm    cmpAlm_1_0_t
}

type cmpAlm_1_t cmpAlm_1_0_t
type BDSAlm_1_t BDSAlm_1_0_t

/*--BDSIon_1_0_t : ----------------------------------------------------------*/
/* BeiDou Ionospheric delay model parameters */

/** cmpIon_1_0_t */
type cmpIon_1_0_t struct {
	/* Time Header */
	TOW uint32
	WNc uint16

	PRN      uint8
	Reserved uint8
	alpha_0  float32 /* (sec) [069-076:4p18] */
	alpha_1  float32 /* (sec/semicircle) [077-084:4p18] */
	alpha_2  float32 /* (sec/semicircle^2) [091-098:4p18] */
	alpha_3  float32 /* (sec/semicircle^3) [099-106:4p18] */
	beta_0   float32 /* (sec) [107-114:4p18] */
	beta_1   float32 /* (sec/semicircle) [121-128:4p18] */
	beta_2   float32 /* (sec/semicircle^2) [129-136:4p18] */
	beta_3   float32 /* (sec/semicircle^3) [137-144:4p18] */
}

/** BDSIon_1_0_t */
type BDSIon_1_0_t struct {
	Header BlockHeader_t
	Ion    cmpIon_1_0_t
}

type cmpIon_1_t cmpIon_1_0_t
type BDSIon_1_t BDSIon_1_0_t

/*--BDSUtc_1_0_t : ----------------------------------------------------------*/
/* BDT-UTC data */

/** cmpUtc_1_0_t */
type cmpUtc_1_0_t struct {
	/* Time Header */
	TOW uint32
	WNc uint16

	PRN       uint8
	Reserved  uint8
	A_1       float32   /* (sec/sec) [151-176:4p18] */
	A_0       SBFDOUBLE /* (sec) [181-194:4] [211-218:4] */
	DEL_t_LS  int8      /* (sec) [241-218:4p18] */
	WN_LSF    uint8     /* (wk) [219-226:4p18] */
	DN        uint8
	DEL_t_LSF int8 /* (sec) [271-278:4p18] */
}

/** BDSUtc_1_0_t */
type BDSUtc_1_0_t struct {
	Header BlockHeader_t
	Utc    cmpUtc_1_0_t
}

type cmpUtc_1_t cmpUtc_1_0_t
type BDSUtc_1_t BDSUtc_1_0_t

/*==QZSS Decoded Message Blocks==============================================*/

/*--QZSNav_1_0_t : ----------------------------------------------------------*/
/* QZSS ephemeris and clock */

/** qzEph_1_0_t */
type qzEph_1_0_t struct {
	/* Time Header */
	TOW uint32
	WNc uint16

	PRN        uint8
	Reserved   uint8
	WN         uint16    /* Week number (modulo 1024). [061-070:1] */
	CAorPonL2  uint8     /* C/A code or P code on L2. [071-072:1] */
	URA        uint8     /* User range accuracy index. [073-076:1] */
	health     uint8     /* SV health. [077-082:1] */
	L2DataFlag uint8     /* L2 P data flag. [091-091:1] */
	IODC       uint16    /* Issue of data clock. [083-084:1] [211-218:1] */
	IODE2      uint8     /* Issue of data eph. frame 2. [061-068:2] */
	IODE3      uint8     /* Issue of data eph. frame 3. [271-278:3] */
	FitIntFlg  uint8     /* fit interval flag [287-287:2] */
	Reserved2  uint8     /* introduced w.r.t. 32-bits memory alignment */
	T_gd       float32   /* Correction term T_gd (s). [197-204:1] */
	t_oc       uint32    /* Clock correction t_oc (s). [219-234:1] */
	a_f2       float32   /* Clock correction a_f2 (s/s^2). [241-248:1] */
	a_f1       float32   /* Clock correction a_f1 (s/s). [249-264:1] */
	a_f0       float32   /* Clock correction a_f0 (s). [271-292:1] */
	C_rs       float32   /* radius sin ampl (m) [069-084:2] */
	DEL_N      float32   /* mean motion diff (semi-circles/s) [091-106:2] */
	M_0        SBFDOUBLE /* Mean Anom (semi-circles) [107-114:2] [121-144:2] */
	C_uc       float32   /* lat cosine ampl (r) [151-166:2] */
	e          SBFDOUBLE /* Eccentricity [167-174:2] [181-204:2] */
	C_us       float32   /* Lat sine ampl   (r) [211-226:2] */
	SQRT_A     SBFDOUBLE /* SQRT(A) (m^1/2) [227-234:2 241-264:2] */
	t_oe       uint32    /* Reference time of ephemeris (s) [271-286:2] */
	C_ic       float32   /* inclin cos ampl (r) [061-076:3] */
	OMEGA_0    SBFDOUBLE /* Right Ascen at TOA (semi-circles) [077-084:3] [091-114:3] */
	C_is       float32   /* inclin sin ampl (r) [121-136:3] */
	i_0        SBFDOUBLE /* Orbital Inclination (semi-circles) [137-144:3] [151-174:3] */
	C_rc       float32   /* radius cos ampl (m) [181-196:3] */
	omega      SBFDOUBLE /* Argument of Perigee(semi-circle) [197-204:3] [211-234:3] */
	OMEGADOT   float32   /* Rate of Right Ascen(semi-circles/s) [241-264:3] */
	IDOT       float32   /* Rate of inclin (semi-circles/s) [279-292:3] */
	WNt_oc     uint16    /* modified WN to go with t_oc (still modulo 1024) */
	WNt_oe     uint16    /* modified WN to go with t_oe (still modulo 1024) */
}

/** QZSNav_1_0_t */
type QZSNav_1_0_t struct {
	Header BlockHeader_t
	Eph    qzEph_1_0_t
}

type qzEph_1_t qzEph_1_0_t
type QZSNav_1_t QZSNav_1_0_t

/*--QZSAlm_1_0_t : ----------------------------------------------------------*/
/* Almanac data for a QZSS satellite */

/** qzAlm_1_0_t */
type qzAlm_1_0_t struct {
	/* Time Header */
	TOW uint32
	WNc uint16

	PRN       uint8
	Reserved  uint8
	e         float32 /* Difference from reference eccentricity */
	t_oa      uint32  /* Reference time of almanac (s) */
	delta_i   float32 /* Orbital Inclination (semi-circles) */
	OMEGADOT  float32 /* Rate of Right Ascen (r/s) */
	SQRT_A    float32 /* SQRT(A) (m^1/2) */
	OMEGA_0   float32 /* Right Ascen at TOA (semi-circles) */
	omega     float32 /* Argument of Perigee (semi-circles) */
	M_0       float32 /* Mean Anom (semi-circles) */
	a_f1      float32 /* Clock correction a_f1 (s/s). */
	a_f0      float32 /* Clock correction a_f0 (s). */
	WN_a      uint8   /* Reference week number of almanac. */
	Reserved2 uint8
	health8   uint8 /* -bit 00-07: 8bit health data in alm. subframes. */
	health6   uint8 /* -bit 08-13: 6bit health data SF4P25 and SF5P25. */
}

/** QZSAlm_1_0_t */
type QZSAlm_1_0_t struct {
	Header BlockHeader_t
	Alm    qzAlm_1_0_t
}

type qzAlm_1_t qzAlm_1_0_t
type QZSAlm_1_t QZSAlm_1_0_t

/*==SBAS L1 Decoded Message Blocks===========================================*/

/*--GEOMT00_1_0_t : ---------------------------------------------------------*/
/* MT00 : SBAS Don't use for safety applications */

/** raMT00_1_0_t */
type raMT00_1_0_t struct {
	/* Time Header */
	TOW uint32
	WNc uint16

	PRN uint8
}

/** GEOMT00_1_0_t */
type GEOMT00_1_0_t struct {
	Header  BlockHeader_t
	GeoMT00 raMT00_1_0_t
}

type raMT00_1_t raMT00_1_0_t
type GEOMT00_1_t GEOMT00_1_0_t

/*--GEOPRNMask_1_0_t : ------------------------------------------------------*/
/* MT01 : PRN Mask assignments */

/** raPRNMask_1_0_t */
type raPRNMask_1_0_t struct {
	/* Time Header */
	TOW uint32
	WNc uint16

	PRN     uint8
	IODP    uint8
	NbrPRNs uint8
	PRNMask [SBF_RAPRNMASK_1_0_PRNMASK_LENGTH]uint8
}

/** GEOPRNMask_1_0_t */
type GEOPRNMask_1_0_t struct {
	Header     BlockHeader_t
	GeoPRNMask raPRNMask_1_0_t
}

type raPRNMask_1_t raPRNMask_1_0_t
type GEOPRNMask_1_t GEOPRNMask_1_0_t

/*--GEOFastCorr_1_0_t : -----------------------------------------------------*/
/* MT02-05/24: Fast Corrections */

/** FastCorr_1_0_t */
type FastCorr_1_0_t struct {
	PRNMaskNo uint8
	UDREI     uint8
	Reserved  [2]uint8
	PRC       float32
}

/** raFastCorr_1_0_t */
type raFastCorr_1_0_t struct {
	/* Time Header */
	TOW uint32
	WNc uint16

	PRN      uint8
	MT       uint8
	IODP     uint8
	IODF     uint8
	N        uint8
	SBLength uint8
	FastCorr [SBF_RAFASTCORR_1_0_FASTCORR_LENGTH]FastCorr_1_0_t
}

/** GEOFastCorr_1_0_t */
type GEOFastCorr_1_0_t struct {
	Header      BlockHeader_t
	GeoFastCorr raFastCorr_1_0_t
}

type FastCorr_1_t FastCorr_1_0_t
type raFastCorr_1_t raFastCorr_1_0_t
type GEOFastCorr_1_t GEOFastCorr_1_0_t

/*--GEOIntegrity_1_0_t : ----------------------------------------------------*/
/* MT06 : Integrity information */

/** raIntegrity_1_0_t */
type raIntegrity_1_0_t struct {
	/* Time Header */
	TOW uint32
	WNc uint16

	PRN      uint8
	Reserved uint8
	IODF     [SBF_RAINTEGRITY_1_0_IODF_LENGTH]uint8
	UDREI    [SBF_RAINTEGRITY_1_0_UDREI_LENGTH]uint8
}

/** GEOIntegrity_1_0_t */
type GEOIntegrity_1_0_t struct {
	Header       BlockHeader_t
	GeoIntegrity raIntegrity_1_0_t
}

type raIntegrity_1_t raIntegrity_1_0_t
type GEOIntegrity_1_t GEOIntegrity_1_0_t

/*--GEOFastCorrDegr_1_0_t : -------------------------------------------------*/
/* MT07 : Fast correction degradation factors */

/** raFastCorrDegr_1_0_t */
type raFastCorrDegr_1_0_t struct {
	/* Time Header */
	TOW uint32
	WNc uint16

	PRN   uint8
	IODP  uint8
	t_lat uint8
	ai    [SBF_RAFASTCORRDEGR_1_0_AI_LENGTH]uint8
}

/** GEOFastCorrDegr_1_0_t */
type GEOFastCorrDegr_1_0_t struct {
	Header          BlockHeader_t
	GeoFastCorrDegr raFastCorrDegr_1_0_t
}

type raFastCorrDegr_1_t raFastCorrDegr_1_0_t
type GEOFastCorrDegr_1_t GEOFastCorrDegr_1_0_t

/*--GEONav_1_0_t : ----------------------------------------------------------*/
/* MT09 : SBAS navigation message */

/** raEph_1_0_t */
type raEph_1_0_t struct {
	/* Time Header */
	TOW uint32
	WNc uint16

	PRN      uint8
	Reserved uint8
	IODN     uint16 /* Issue of data, 8-bit, cycles from 0 to 255 */
	URA      uint16 /* user range accuracy [0,15] */
	t0       uint32 /* time of day of ephemeris */
	Xg       SBFDOUBLE
	Yg       SBFDOUBLE
	Zg       SBFDOUBLE
	Xgd      SBFDOUBLE
	Ygd      SBFDOUBLE
	Zgd      SBFDOUBLE
	Xgdd     SBFDOUBLE
	Ygdd     SBFDOUBLE
	Zgdd     SBFDOUBLE
	aGf0     float32 /* Clock bias  [s] */
	aGf1     float32 /* Clock drift [s/s] */
}

/** GEONav_1_0_t */
type GEONav_1_0_t struct {
	Header BlockHeader_t
	Eph    raEph_1_0_t
}

type raEph_1_t raEph_1_0_t
type GEONav_1_t GEONav_1_0_t

/*--GEODegrFactors_1_0_t : --------------------------------------------------*/
/* MT10 : Degradation factors */

/** raDF_1_0_t */
type raDF_1_0_t struct {
	Brrc        SBFDOUBLE
	Cltc_lsb    SBFDOUBLE
	Cltc_v1     SBFDOUBLE
	Iltc_v1     uint32
	Cltc_v0     SBFDOUBLE
	Iltc_v0     uint32
	Cgeo_lsb    SBFDOUBLE
	Cgeo_v      SBFDOUBLE
	Igeo        uint32
	Cer         float32
	Ciono_step  SBFDOUBLE
	Iiono       uint32
	Ciono_ramp  SBFDOUBLE
	RSSudre     uint8
	RSSiono     uint8
	Reserved2   [2]uint8
	Ccovariance SBFDOUBLE
}

/** raDegrFactors_1_0_t */
type raDegrFactors_1_0_t struct {
	/* Time Header */
	TOW uint32
	WNc uint16

	PRN      uint8
	Reserved uint8
	DFactors raDF_1_0_t
}

/** GEODegrFactors_1_0_t */
type GEODegrFactors_1_0_t struct {
	Header         BlockHeader_t
	GeoDegrFactors raDegrFactors_1_0_t
}

type raDF_1_t raDF_1_0_t
type raDegrFactors_1_t raDegrFactors_1_0_t
type GEODegrFactors_1_t GEODegrFactors_1_0_t

/*--GEONetworkTime_1_0_t : --------------------------------------------------*/
/* MT12 : SBAS Network Time/UTC offset parameters */

/** NetworkTimeMsg_1_0_t */
type NetworkTimeMsg_1_0_t struct {
	A_1       float32
	A_0       SBFDOUBLE
	t_ot      uint32
	WN_t      uint8
	DEL_t_LS  int8
	WN_LSF    uint8
	DN        uint8
	DEL_t_LSF int8
	UTC_std   uint8
	GPS_WN    uint16
	GPS_TOW   uint32
	GlonassID uint8
}

/** raNetworkTime_1_0_t */
type raNetworkTime_1_0_t struct {
	/* Time Header */
	TOW uint32
	WNc uint16

	PRN        uint8
	Reserved   uint8
	NWTMessage NetworkTimeMsg_1_0_t
}

/** GEONetworkTime_1_0_t */
type GEONetworkTime_1_0_t struct {
	Header         BlockHeader_t
	GeoNetworkTime raNetworkTime_1_0_t
}

type NetworkTimeMsg_1_t NetworkTimeMsg_1_0_t
type raNetworkTime_1_t raNetworkTime_1_0_t
type GEONetworkTime_1_t GEONetworkTime_1_0_t

/*--GEOAlm_1_0_t : ----------------------------------------------------------*/
/* MT17 : SBAS satellite almanac */

/** raAlm_1_0_t */
type raAlm_1_0_t struct {
	/* Time Header */
	TOW uint32
	WNc uint16

	PRN       uint8
	Reserved0 uint8
	DataID    uint8
	Reserved1 uint8
	Health    uint16
	t_oa      uint32
	Xg        SBFDOUBLE
	Yg        SBFDOUBLE
	Zg        SBFDOUBLE
	Xgd       SBFDOUBLE
	Ygd       SBFDOUBLE
	Zgd       SBFDOUBLE
}

/** GEOAlm_1_0_t */
type GEOAlm_1_0_t struct {
	Header BlockHeader_t
	Alm    raAlm_1_0_t
}

type raAlm_1_t raAlm_1_0_t
type GEOAlm_1_t GEOAlm_1_0_t

/*--GEOIGPMask_1_0_t : ------------------------------------------------------*/
/* MT18 : Ionospheric grid point mask */

/** raIGPMask_1_0_t */
type raIGPMask_1_0_t struct {
	/* Time Header */
	TOW uint32
	WNc uint16

	PRN      uint8
	NbrBands uint8
	BandNbr  uint8
	IODI     uint8
	NbrIGPs  uint8
	IGPMask  [SBF_RAIGPMASK_1_0_IGPMASK_LENGTH]uint8
}

/** GEOIGPMask_1_0_t */
type GEOIGPMask_1_0_t struct {
	Header     BlockHeader_t
	GeoIGPMask raIGPMask_1_0_t
}

type raIGPMask_1_t raIGPMask_1_0_t
type GEOIGPMask_1_t GEOIGPMask_1_0_t

/*--GEOLongTermCorr_1_0_t : -------------------------------------------------*/
/* MT24/25 : Long term satellite error corrections */

/** LTCorr_1_0_t */
type LTCorr_1_0_t struct {
	VelocityCode uint8
	PRNMaskNo    uint8
	IODP         uint8
	IODE         uint8
	dx           float32
	dy           float32
	dz           float32
	dxRate       float32 /* set to 0 if velocity code = 0 */
	dyRate       float32 /* set to 0 if velocity code = 0 */
	dzRate       float32 /* set to 0 if velocity code = 0 */
	da_f0        float32
	da_f1        float32 /* set to 0 if velocity code = 0 */
	t_oe         uint32  /* set to 0 if velocity code = 0 */
}

/** raLongTermCorr_1_0_t */
type raLongTermCorr_1_0_t struct {
	/* Time Header */
	TOW uint32
	WNc uint16

	PRN      uint8
	N        uint8
	SBLength uint8
	Reserved [3]uint8
	LTCorr   [SBF_RALONGTERMCORR_1_0_LTCORR_LENGTH]LTCorr_1_0_t
}

/** GEOLongTermCorr_1_0_t */
type GEOLongTermCorr_1_0_t struct {
	Header          BlockHeader_t
	GeoLongTermCorr raLongTermCorr_1_0_t
}

type LTCorr_1_t LTCorr_1_0_t
type raLongTermCorr_1_t raLongTermCorr_1_0_t
type GEOLongTermCorr_1_t GEOLongTermCorr_1_0_t

/*--GEOIonoDelay_1_0_t : ----------------------------------------------------*/
/* MT26 : Ionospheric delay corrections */

/** IDC_1_0_t */
type IDC_1_0_t struct {
	IGPMaskNo     uint8
	GIVEI         uint8
	Reserved      [2]uint8
	VerticalDelay float32
}

/** raIonoDelay_1_0_t */
type raIonoDelay_1_0_t struct {
	/* Time Header */
	TOW uint32
	WNc uint16

	PRN      uint8
	BandNbr  uint8
	IODI     uint8
	N        uint8
	SBLength uint8
	Reserved uint8
	IDC      [SBF_RAIONODELAY_1_0_IDC_LENGTH]IDC_1_0_t
}

/** GEOIonoDelay_1_0_t */
type GEOIonoDelay_1_0_t struct {
	Header       BlockHeader_t
	GeoIonoDelay raIonoDelay_1_0_t
}

type IDC_1_t IDC_1_0_t
type raIonoDelay_1_t raIonoDelay_1_0_t
type GEOIonoDelay_1_t GEOIonoDelay_1_0_t

/*--GEOServiceLevel_1_0_t : -------------------------------------------------*/
/* MT27 : SBAS Service Message */

/** ServiceRegion_1_0_t */
type ServiceRegion_1_0_t struct {
	Latitude1   int8 /* [-90,90] */
	Latitude2   int8
	Longitude1  int16 /* [-180,180] */
	Longitude2  int16
	RegionShape uint8
	_padding    [SBF_SERVICEREGION_1_0__PADDING_LENGTH]uint8
}

/** raServiceMsg_1_0_t */
type raServiceMsg_1_0_t struct {
	IODS         uint8 /* Issue of Data [0,7] */
	nrMessages   uint8 /* Number of Messages [1,8] */
	MessageNR    uint8 /* Message nr [1,8] */
	PriorityCode uint8 /* [0,3] */
	dUDREI_In    uint8 /* Delta UDRE Indicator [0,15] */
	dUDREI_Out   uint8 /* [0,15] */
	N            uint8 /* Number of Regions [0,7] */
	SBLength     uint8 /* for SBF sub-block */
	Regions      [SBF_RASERVICEMSG_1_0_REGIONS_LENGTH]ServiceRegion_1_0_t
}

/** raServiceLevel_1_0_t */
type raServiceLevel_1_0_t struct {
	/* Time Header */
	TOW uint32
	WNc uint16

	PRN            uint8
	Reserved       uint8
	ServiceMessage raServiceMsg_1_0_t
}

/** GEOServiceLevel_1_0_t */
type GEOServiceLevel_1_0_t struct {
	Header          BlockHeader_t
	GeoServiceLevel raServiceLevel_1_0_t
}

type ServiceRegion_1_t ServiceRegion_1_0_t
type raServiceMsg_1_t raServiceMsg_1_0_t
type raServiceLevel_1_t raServiceLevel_1_0_t
type GEOServiceLevel_1_t GEOServiceLevel_1_0_t

/*--GEOClockEphCovMatrix_1_0_t : --------------------------------------------*/
/* MT28 : Clock-Ephemeris Covariance Matrix */

/** CovMatrix_1_0_t */
type CovMatrix_1_0_t struct {
	PRNMaskNo uint8
	Reserved  [2]uint8
	ScaleExp  uint8
	E11       uint16
	E22       uint16
	E33       uint16
	E44       uint16
	E12       int16
	E13       int16
	E14       int16
	E23       int16
	E24       int16
	E34       int16
}

/** raClockEphCovMatrix_1_0_t */
type raClockEphCovMatrix_1_0_t struct {
	/* Time Header */
	TOW uint32
	WNc uint16

	PRN       uint8
	IODP      uint8
	N         uint8
	SBLength  uint8
	Reserved  [2]uint8
	CovMatrix [SBF_RACLOCKEPHCOVMATRIX_1_0_COVMATRIX_LENGTH]CovMatrix_1_0_t
}

/** GEOClockEphCovMatrix_1_0_t */
type GEOClockEphCovMatrix_1_0_t struct {
	Header               BlockHeader_t
	GeoClockEphCovMatrix raClockEphCovMatrix_1_0_t
}

type CovMatrix_1_t CovMatrix_1_0_t
type raClockEphCovMatrix_1_t raClockEphCovMatrix_1_0_t
type GEOClockEphCovMatrix_1_t GEOClockEphCovMatrix_1_0_t

/*==SBAS L5 Decoded Message Blocks===========================================*/

/*--SBASL5Nav_1_0_t : -------------------------------------------------------*/
/* DFMC SBAS ephemeris and clock data */

/** sbsEph_1_0_t */
type sbsEph_1_0_t struct {
	/* Time Header */
	TOW uint32
	WNc uint16

	PRN        uint8
	Reserved   [2]uint8
	Slot_Delta uint8
	IODG       uint8
	providerID uint8
	C_uc       float32
	C_us       float32
	IDOT       float32
	omega      SBFDOUBLE
	OMEGA_0    SBFDOUBLE
	M_0        SBFDOUBLE
	a_f1       float32
	a_f0       float32
	i_0        SBFDOUBLE
	e          SBFDOUBLE
	A          SBFDOUBLE
	t_oe       uint32
}

/** SBASL5Nav_1_0_t */
type SBASL5Nav_1_0_t struct {
	Header BlockHeader_t
	Eph    sbsEph_1_0_t
}

type sbsEph_1_t sbsEph_1_0_t
type SBASL5Nav_1_t SBASL5Nav_1_0_t

/*--SBASL5Alm_1_0_t : -------------------------------------------------------*/
/* DFMC SBAS almanac data */

/** sbsAlm_1_0_t */
type sbsAlm_1_0_t struct {
	/* Time Header */
	TOW uint32
	WNc uint16

	PRN                uint8
	Reserved           uint8
	providerID         uint8
	BroadcastIndicator uint8
	omega              float32
	OMEGA_0            float32
	OMEGADOT           float32
	M_0                float32
	i_0                float32
	e                  float32
	A                  float32
	t_oa               uint32
	PRN_A              uint8
	_padding           [SBF_SBSALM_1_0__PADDING_LENGTH]uint8
}

/** SBASL5Alm_1_0_t */
type SBASL5Alm_1_0_t struct {
	Header BlockHeader_t
	Alm    sbsAlm_1_0_t
}

type sbsAlm_1_t sbsAlm_1_0_t
type SBASL5Alm_1_t SBASL5Alm_1_0_t

/*==Position, Velocity and Time Blocks=======================================*/

/*--PVTCartesian_1_0_t : ----------------------------------------------------*/
/* PVT in Cartesian coordinates */

/** PVTCartesian_1_0_t */
type PVTCartesian_1_0_t struct {
	Header BlockHeader_t

	/* Time Header */
	TOW uint32
	WNc uint16

	NrSV          uint8
	Error         uint8
	Mode          uint8
	System        uint8
	Info          uint8
	SBASprn       uint8
	X             SBFDOUBLE
	Y             SBFDOUBLE
	Z             SBFDOUBLE
	Vx            float32
	Vy            float32
	Vz            float32
	RxClkBias     SBFDOUBLE
	RxClkDrift    float32
	MeanCorrAge   uint16
	BaseStationID uint16
	Cog           float32
}

type PVTCartesian_1_t PVTCartesian_1_0_t

/*--PVTGeodetic_1_0_t : -----------------------------------------------------*/
/* PVT in geodetic coordinates */

/** PVTGeodetic_1_0_t */
type PVTGeodetic_1_0_t struct {
	Header BlockHeader_t

	/* Time Header */
	TOW uint32
	WNc uint16

	NrSV          uint8
	Error         uint8
	Mode          uint8
	System        uint8
	Info          uint8
	SBASprn       uint8
	Lat           SBFDOUBLE
	Lon           SBFDOUBLE
	Alt           SBFDOUBLE
	Vn            float32
	Ve            float32
	Vu            float32
	RxClkBias     SBFDOUBLE
	RxClkDrift    float32
	GeoidHeight   float32
	MeanCorrAge   uint16
	BaseStationID uint16
	Cog           float32
}

type PVTGeodetic_1_t PVTGeodetic_1_0_t

/*--DOP_1_0_t : -------------------------------------------------------------*/
/* Dilution of precision */

/** DOP_1_0_t */
type DOP_1_0_t struct {
	Header BlockHeader_t

	/* Time Header */
	TOW uint32
	WNc uint16

	NrSV  uint8
	Error uint8
	PDOP  uint16
	TDOP  uint16
	HDOP  uint16
	VDOP  uint16
	HPL   float32
	VPL   float32
}

type DOP_1_t DOP_1_0_t

/*--PVTResiduals_1_0_t : ----------------------------------------------------*/
/* Measurement residuals */

/** PVTResidual_1_0_t */
type PVTResidual_1_0_t struct {
	CACodeRes    int16
	P1CodeRes    int16
	P2CodeRes    int16
	DopplerL1Res int16
	DopplerL2Res int16
	PRN          uint8
	_padding     [SBF_PVTRESIDUAL_1_0__PADDING_LENGTH]uint8
}

/** PVTResiduals_1_0_t */
type PVTResiduals_1_0_t struct {
	Header BlockHeader_t

	/* Time Header */
	TOW uint32
	WNc uint16

	N           uint8
	SBSize      uint8
	SatResidual [SBF_PVTRESIDUALS_1_0_SATRESIDUAL_LENGTH]PVTResidual_1_0_t
}

type PVTResidual_1_t PVTResidual_1_0_t
type PVTResiduals_1_t PVTResiduals_1_0_t

/*--RAIMStatistics_1_0_t : --------------------------------------------------*/
/* Integrity statistics */

/** RAIMSatData_1_0_t */
type RAIMSatData_1_0_t struct {
	PRN         uint8
	AntennaID   uint8
	TestResults uint8
	Reserved    uint8
	UnityRangeW uint16
	UnityRrateW uint16
	RangeMDB    uint16
	RrateMDB    uint16
}

/** RAIMStatistics_1_0_t */
type RAIMStatistics_1_0_t struct {
	Header BlockHeader_t

	/* Time Header */
	TOW uint32
	WNc uint16

	Integrity          uint8
	Reserved1          uint8
	PositionHERL       float32
	PositionVERL       float32
	VelocityHERL       float32
	VelocityVERL       float32
	UnityOverallModelP uint16
	UnityOverallModelV uint16
	Reserved2          [2]uint8
	N                  uint8
	SBSize             uint8
	RAIMChannel        [SBF_RAIMSTATISTICS_1_0_RAIMCHANNEL_LENGTH]RAIMSatData_1_0_t
}

type RAIMSatData_1_t RAIMSatData_1_0_t
type RAIMStatistics_1_t RAIMStatistics_1_0_t

/*--PVTCartesian_2_0_t : ----------------------------------------------------*/
/* Position, velocity, and time in Cartesian coordinates */

/** PVTCartesian_2_0_t */
type PVTCartesian_2_0_t struct {
	Header BlockHeader_t

	/* Time Header */
	TOW uint32
	WNc uint16

	Mode        uint8
	Error       uint8
	X           SBFDOUBLE
	Y           SBFDOUBLE
	Z           SBFDOUBLE
	Undulation  float32
	Vx          float32
	Vy          float32
	Vz          float32
	COG         float32
	RxClkBias   SBFDOUBLE
	RxClkDrift  float32
	TimeSystem  uint8
	Datum       uint8
	NrSV        uint8
	WACorrInfo  uint8
	ReferenceId uint16
	MeanCorrAge uint16
	SignalInfo  uint32
	AlertFlag   uint8
	_padding    [SBF_PVTCARTESIAN_2_0__PADDING_LENGTH]uint8
}

/*--PVTCartesian_2_1_t : ----------------------------------------------------*/
/* Position, velocity, and time in Cartesian coordinates */

/** PVTCartesian_2_1_t */
type PVTCartesian_2_1_t struct {
	Header BlockHeader_t

	/* Time Header */
	TOW uint32
	WNc uint16

	Mode        uint8
	Error       uint8
	X           SBFDOUBLE
	Y           SBFDOUBLE
	Z           SBFDOUBLE
	Undulation  float32
	Vx          float32
	Vy          float32
	Vz          float32
	COG         float32
	RxClkBias   SBFDOUBLE
	RxClkDrift  float32
	TimeSystem  uint8
	Datum       uint8
	NrSV        uint8
	WACorrInfo  uint8
	ReferenceId uint16
	MeanCorrAge uint16
	SignalInfo  uint32
	AlertFlag   uint8
	NrBases     uint8
	PPPInfo     uint16
}

/*--PVTCartesian_2_2_t : ----------------------------------------------------*/
/* Position, velocity, and time in Cartesian coordinates */

/** PVTCartesian_2_2_t */
type PVTCartesian_2_2_t struct {
	Header BlockHeader_t

	/* Time Header */
	TOW uint32
	WNc uint16

	Mode        uint8
	Error       uint8
	X           SBFDOUBLE
	Y           SBFDOUBLE
	Z           SBFDOUBLE
	Undulation  float32
	Vx          float32
	Vy          float32
	Vz          float32
	COG         float32
	RxClkBias   SBFDOUBLE
	RxClkDrift  float32
	TimeSystem  uint8
	Datum       uint8
	NrSV        uint8
	WACorrInfo  uint8
	ReferenceId uint16
	MeanCorrAge uint16
	SignalInfo  uint32
	AlertFlag   uint8
	NrBases     uint8
	PPPInfo     uint16
	Latency     uint16
	HAccuracy   uint16
	VAccuracy   uint16
	Misc        uint8
	_padding    [SBF_PVTCARTESIAN_2_2__PADDING_LENGTH]uint8
}

type PVTCartesian_2_t PVTCartesian_2_2_t

/*--PVTGeodetic_2_0_t : -----------------------------------------------------*/
/* Position, velocity, and time in geodetic coordinates */

/** PVTGeodetic_2_0_t */
type PVTGeodetic_2_0_t struct {
	Header BlockHeader_t

	/* Time Header */
	TOW uint32
	WNc uint16

	Mode        uint8
	Error       uint8
	Lat         SBFDOUBLE
	Lon         SBFDOUBLE
	Alt         SBFDOUBLE
	Undulation  float32
	Vn          float32
	Ve          float32
	Vu          float32
	COG         float32
	RxClkBias   SBFDOUBLE
	RxClkDrift  float32
	TimeSystem  uint8
	Datum       uint8
	NrSV        uint8
	WACorrInfo  uint8
	ReferenceId uint16
	MeanCorrAge uint16
	SignalInfo  uint32
	AlertFlag   uint8
	_padding    [SBF_PVTGEODETIC_2_0__PADDING_LENGTH]uint8
}

/*--PVTGeodetic_2_1_t : -----------------------------------------------------*/
/* Position, velocity, and time in geodetic coordinates */

/** PVTGeodetic_2_1_t */
type PVTGeodetic_2_1_t struct {
	Header BlockHeader_t

	/* Time Header */
	TOW uint32
	WNc uint16

	Mode        uint8
	Error       uint8
	Lat         SBFDOUBLE
	Lon         SBFDOUBLE
	Alt         SBFDOUBLE
	Undulation  float32
	Vn          float32
	Ve          float32
	Vu          float32
	COG         float32
	RxClkBias   SBFDOUBLE
	RxClkDrift  float32
	TimeSystem  uint8
	Datum       uint8
	NrSV        uint8
	WACorrInfo  uint8
	ReferenceId uint16
	MeanCorrAge uint16
	SignalInfo  uint32
	AlertFlag   uint8
	NrBases     uint8
	PPPInfo     uint16
}

/*--PVTGeodetic_2_2_t : -----------------------------------------------------*/
/* Position, velocity, and time in geodetic coordinates */

/** PVTGeodetic_2_2_t */
type PVTGeodetic_2_2_t struct {
	Header BlockHeader_t

	/* Time Header */
	TOW uint32
	WNc uint16

	Mode        uint8
	Error       uint8
	Lat         SBFDOUBLE
	Lon         SBFDOUBLE
	Alt         SBFDOUBLE
	Undulation  float32
	Vn          float32
	Ve          float32
	Vu          float32
	COG         float32
	RxClkBias   SBFDOUBLE
	RxClkDrift  float32
	TimeSystem  uint8
	Datum       uint8
	NrSV        uint8
	WACorrInfo  uint8
	ReferenceId uint16
	MeanCorrAge uint16
	SignalInfo  uint32
	AlertFlag   uint8
	NrBases     uint8
	PPPInfo     uint16
	Latency     uint16
	HAccuracy   uint16
	VAccuracy   uint16
	Misc        uint8
	_padding    [SBF_PVTGEODETIC_2_2__PADDING_LENGTH]uint8
}

type PVTGeodetic_2_t PVTGeodetic_2_2_t

/*--PVTGeodeticAuth_1_0_t : -------------------------------------------------*/
/* OSNMA-Authenticated Position, velocity, and time in geodetic coordinates */

/** PVTGeodeticAuth_1_0_t */
type PVTGeodeticAuth_1_0_t struct {
	Header BlockHeader_t

	/* Time Header */
	TOW uint32
	WNc uint16

	Mode        uint8
	Error       uint8
	Latitude    SBFDOUBLE
	Longitude   SBFDOUBLE
	Height      SBFDOUBLE
	Undulation  float32
	Vn          float32
	Ve          float32
	Vu          float32
	COG         float32
	RxClkBias   SBFDOUBLE
	RxClkDrift  float32
	TimeSystem  uint8
	Datum       uint8
	NrSV        uint8
	WACorrInfo  uint8
	ReferenceID uint16
	MeanCorrAge uint16
	SignalInfo  uint32
	AlertFlag   uint8
	_padding    [SBF_PVTGEODETICAUTH_1_0__PADDING_LENGTH]uint8
}

/*--PVTGeodeticAuth_1_1_t : -------------------------------------------------*/
/* OSNMA-Authenticated Position, velocity, and time in geodetic coordinates */

/** PVTGeodeticAuth_1_1_t */
type PVTGeodeticAuth_1_1_t struct {
	Header BlockHeader_t

	/* Time Header */
	TOW uint32
	WNc uint16

	Mode        uint8
	Error       uint8
	Latitude    SBFDOUBLE
	Longitude   SBFDOUBLE
	Height      SBFDOUBLE
	Undulation  float32
	Vn          float32
	Ve          float32
	Vu          float32
	COG         float32
	RxClkBias   SBFDOUBLE
	RxClkDrift  float32
	TimeSystem  uint8
	Datum       uint8
	NrSV        uint8
	WACorrInfo  uint8
	ReferenceID uint16
	MeanCorrAge uint16
	SignalInfo  uint32
	AlertFlag   uint8
	NrBases     uint8
	PPPInfo     uint16
}

/*--PVTGeodeticAuth_1_2_t : -------------------------------------------------*/
/* OSNMA-Authenticated Position, velocity, and time in geodetic coordinates */

/** PVTGeodeticAuth_1_2_t */
type PVTGeodeticAuth_1_2_t struct {
	Header BlockHeader_t

	/* Time Header */
	TOW uint32
	WNc uint16

	Mode        uint8
	Error       uint8
	Latitude    SBFDOUBLE
	Longitude   SBFDOUBLE
	Height      SBFDOUBLE
	Undulation  float32
	Vn          float32
	Ve          float32
	Vu          float32
	COG         float32
	RxClkBias   SBFDOUBLE
	RxClkDrift  float32
	TimeSystem  uint8
	Datum       uint8
	NrSV        uint8
	WACorrInfo  uint8
	ReferenceID uint16
	MeanCorrAge uint16
	SignalInfo  uint32
	AlertFlag   uint8
	NrBases     uint8
	PPPInfo     uint16
	Latency     uint16
	HAccuracy   uint16
	VAccuracy   uint16
	Misc        uint8
	_padding    [SBF_PVTGEODETICAUTH_1_2__PADDING_LENGTH]uint8
}

type PVTGeodeticAuth_1_t PVTGeodeticAuth_1_2_t

/*--PosCovCartesian_1_0_t : -------------------------------------------------*/
/* Position covariance matrix (X,Y, Z) */

/** PosCovCartesian_1_0_t */
type PosCovCartesian_1_0_t struct {
	Header BlockHeader_t

	/* Time Header */
	TOW uint32
	WNc uint16

	Mode   uint8
	Error  uint8
	Cov_xx float32
	Cov_yy float32
	Cov_zz float32
	Cov_tt float32
	Cov_xy float32
	Cov_xz float32
	Cov_xt float32
	Cov_yz float32
	Cov_yt float32
	Cov_zt float32
}

type PosCovCartesian_1_t PosCovCartesian_1_0_t

/*--PosCovGeodetic_1_0_t : --------------------------------------------------*/
/* Position covariance matrix (Lat, Lon, Alt) */

/** PosCovGeodetic_1_0_t */
type PosCovGeodetic_1_0_t struct {
	Header BlockHeader_t

	/* Time Header */
	TOW uint32
	WNc uint16

	Mode       uint8
	Error      uint8
	Cov_LatLat float32
	Cov_LonLon float32
	Cov_AltAlt float32
	Cov_tt     float32
	Cov_LatLon float32
	Cov_LatAlt float32
	Cov_Latt   float32
	Cov_LonAlt float32
	Cov_Lont   float32
	Cov_Altt   float32
}

type PosCovGeodetic_1_t PosCovGeodetic_1_0_t

/*--VelCovCartesian_1_0_t : -------------------------------------------------*/
/* Velocity covariance matrix (X, Y, Z) */

/** VelCovCartesian_1_0_t */
type VelCovCartesian_1_0_t struct {
	Header BlockHeader_t

	/* Time Header */
	TOW uint32
	WNc uint16

	Mode     uint8
	Error    uint8
	Cov_VxVx float32
	Cov_VyVy float32
	Cov_VzVz float32
	Cov_DtDt float32
	Cov_VxVy float32
	Cov_VxVz float32
	Cov_VxDt float32
	Cov_VyVz float32
	Cov_VyDt float32
	Cov_VzDt float32
}

type VelCovCartesian_1_t VelCovCartesian_1_0_t

/*--VelCovGeodetic_1_0_t : --------------------------------------------------*/
/* Velocity covariance matrix (North, East, Up) */

/** VelCovGeodetic_1_0_t */
type VelCovGeodetic_1_0_t struct {
	Header BlockHeader_t

	/* Time Header */
	TOW uint32
	WNc uint16

	Mode     uint8
	Error    uint8
	Cov_VnVn float32
	Cov_VeVe float32
	Cov_VuVu float32
	Cov_DtDt float32
	Cov_VnVe float32
	Cov_VnVu float32
	Cov_VnDt float32
	Cov_VeVu float32
	Cov_VeDt float32
	Cov_VuDt float32
}

type VelCovGeodetic_1_t VelCovGeodetic_1_0_t

/*--DOP_2_0_t : -------------------------------------------------------------*/
/* Dilution of precision */

/** DOP_2_0_t */
type DOP_2_0_t struct {
	Header BlockHeader_t

	/* Time Header */
	TOW uint32
	WNc uint16

	NrSV     uint8
	Reserved uint8
	PDOP     uint16
	TDOP     uint16
	HDOP     uint16
	VDOP     uint16
	HPL      float32
	VPL      float32
}

type DOP_2_t DOP_2_0_t

/*--PosCart_1_0_t : ---------------------------------------------------------*/
/* Position, variance and baseline in Cartesian coordinates */

/** PosCart_1_0_t */
type PosCart_1_0_t struct {
	Header BlockHeader_t

	/* Time Header */
	TOW uint32
	WNc uint16

	Mode        uint8
	Error       uint8
	X           SBFDOUBLE
	Y           SBFDOUBLE
	Z           SBFDOUBLE
	Base2RoverX SBFDOUBLE
	Base2RoverY SBFDOUBLE
	Base2RoverZ SBFDOUBLE
	Cov_xx      float32
	Cov_yy      float32
	Cov_zz      float32
	Cov_xy      float32
	Cov_xz      float32
	Cov_yz      float32
	PDOP        uint16
	HDOP        uint16
	VDOP        uint16
	Misc        uint8
	Reserved    uint8
	AlertFlag   uint8
	Datum       uint8
	NrSV        uint8
	WACorrInfo  uint8
	ReferenceId uint16
	MeanCorrAge uint16
	SignalInfo  uint32
}

type PosCart_1_t PosCart_1_0_t

/*--PosLocal_1_0_t : --------------------------------------------------------*/
/* Position in a local datum */

/** PosLocal_1_0_t */
type PosLocal_1_0_t struct {
	Header BlockHeader_t

	/* Time Header */
	TOW uint32
	WNc uint16

	Mode     uint8
	Error    uint8
	Lat      SBFDOUBLE
	Lon      SBFDOUBLE
	Alt      SBFDOUBLE
	Datum    uint8
	_padding [SBF_POSLOCAL_1_0__PADDING_LENGTH]uint8
}

type PosLocal_1_t PosLocal_1_0_t

/*--PosProjected_1_0_t : ----------------------------------------------------*/
/* Plane grid coordinates */

/** PosProjected_1_0_t */
type PosProjected_1_0_t struct {
	Header BlockHeader_t

	/* Time Header */
	TOW uint32
	WNc uint16

	Mode     uint8
	Error    uint8
	Northing SBFDOUBLE
	Easting  SBFDOUBLE
	Alt      SBFDOUBLE
	Datum    uint8
	_padding [SBF_POSPROJECTED_1_0__PADDING_LENGTH]uint8
}

type PosProjected_1_t PosProjected_1_0_t

/*--PVTSatCartesian_1_0_t : -------------------------------------------------*/
/* Satellite positions */

/** SatPos_1_0_t */
type SatPos_1_0_t struct {
	SVID   uint8
	FreqNr uint8
	IODE   uint16
	x      SBFDOUBLE
	y      SBFDOUBLE
	z      SBFDOUBLE
	Vx     float32
	Vy     float32
	Vz     float32
}

/** PVTSatCartesian_1_0_t */
type PVTSatCartesian_1_0_t struct {
	Header BlockHeader_t

	/* Time Header */
	TOW uint32
	WNc uint16

	N      uint8
	SBSize uint8
	SatPos [SBF_PVTSATCARTESIAN_1_0_SATPOS_LENGTH]SatPos_1_0_t
}

/*--PVTSatCartesian_1_1_t : -------------------------------------------------*/
/* Satellite positions */
/** SatPos_1_1_t */
type SatPos_1_1_t struct {
	SVID      uint8
	FreqNr    uint8
	IODE      uint16
	x         SBFDOUBLE
	y         SBFDOUBLE
	z         SBFDOUBLE
	Vx        float32
	Vy        float32
	Vz        float32
	IonoMSB   int16
	TropoMSB  int16
	IonoLSB   uint8
	TropoLSB  uint8
	IonoModel uint8
	_padding  [SBF_SATPOS_1_1__PADDING_LENGTH]uint8
}

/** PVTSatCartesian_1_1_t */
type PVTSatCartesian_1_1_t struct {
	Header BlockHeader_t

	/* Time Header */
	TOW uint32
	WNc uint16

	N      uint8
	SBSize uint8
	SatPos [SBF_PVTSATCARTESIAN_1_1_SATPOS_LENGTH]SatPos_1_1_t
}

type SatPos_1_t SatPos_1_1_t
type PVTSatCartesian_1_t PVTSatCartesian_1_1_t

/*--PVTResiduals_2_0_t : ----------------------------------------------------*/
/* Measurement residuals */

/** ResidualInfoCode_2_0_t */
type ResidualInfoCode_2_0_t struct {
	Residual float32
	W        uint16
	MDB      uint16
}

/** ResidualInfoPhase_2_0_t */
type ResidualInfoPhase_2_0_t struct {
	Residual float32
	W        uint16
	MDB      uint16
}

/** ResidualInfoDoppler_2_0_t */
type ResidualInfoDoppler_2_0_t struct {
	Residual float32
	W        uint16
	MDB      uint16
}

/** SatSignalInfo_2_0_t */
type SatSignalInfo_2_0_t struct {
	SVID      uint8
	FreqNr    uint8
	Type      uint8
	RefSVID   uint8
	RefFreqNr uint8
	MeasInfo  uint8
	IODE      uint16
	CorrAge   uint16
	Reserved  uint16 /* Reserved for future use */
}

/** PVTResiduals_2_0_t */
type PVTResiduals_2_0_t struct {
	Header BlockHeader_t

	/* Time Header */
	TOW uint32
	WNc uint16

	N        uint8
	SB1Size  uint8
	SB2Size  uint8
	Reserved [3]uint8
	Data     [SBF_PVTRESIDUALS_2_0_DATA_LENGTH]uint8
}

/*--PVTResiduals_2_1_t : ----------------------------------------------------*/
/* Measurement residuals */

/** ResidualInfoCode_2_1_t */
type ResidualInfoCode_2_1_t struct {
	Residual float32
	W        uint16
	MDB      uint16
}

/** ResidualInfoPhase_2_1_t */
type ResidualInfoPhase_2_1_t struct {
	Residual float32
	W        uint16
	MDB      uint16
}

/** ResidualInfoDoppler_2_1_t */
type ResidualInfoDoppler_2_1_t struct {
	Residual float32
	W        uint16
	MDB      uint16
}

/** SatSignalInfo_2_1_t */
type SatSignalInfo_2_1_t struct {
	SVID        uint8
	FreqNr      uint8
	Type        uint8
	RefSVID     uint8
	RefFreqNr   uint8
	MeasInfo    uint8
	IODE        uint16
	CorrAge     uint16
	ReferenceID uint16
}

/** PVTResiduals_2_1_t */
type PVTResiduals_2_1_t struct {
	Header BlockHeader_t

	/* Time Header */
	TOW uint32
	WNc uint16

	N        uint8
	SB1Size  uint8
	SB2Size  uint8
	Reserved [3]uint8
	Data     [SBF_PVTRESIDUALS_2_1_DATA_LENGTH]uint8
}

type ResidualInfoCode_2_t ResidualInfoCode_2_1_t
type ResidualInfoPhase_2_t ResidualInfoPhase_2_1_t
type ResidualInfoDoppler_2_t ResidualInfoDoppler_2_1_t
type SatSignalInfo_2_t SatSignalInfo_2_1_t
type PVTResiduals_2_t PVTResiduals_2_1_t

/*--RAIMStatistics_2_0_t : --------------------------------------------------*/
/* Integrity statistics */

/** RAIMStatistics_2_0_t */
type RAIMStatistics_2_0_t struct {
	Header BlockHeader_t

	/* Time Header */
	TOW uint32
	WNc uint16

	Integrity         uint8
	Reserved1         uint8
	PositionHERL      float32
	PositionVERL      float32
	VelocityHERL      float32
	VelocityVERL      float32
	UnityOverallModel uint16
	_padding          [SBF_RAIMSTATISTICS_2_0__PADDING_LENGTH]uint8
}

type RAIMStatistics_2_t RAIMStatistics_2_0_t

/*--GEOCorrections_1_0_t : --------------------------------------------------*/
/* Orbit, Clock and pseudoranges SBAS corrections */

/** SatCorr_1_0_t */
type SatCorr_1_0_t struct {
	prn         uint8
	IODE        uint8
	Reserved    [2]uint8
	PRC         float32
	CorrAgeFC   float32
	deltaX      float32
	deltaY      float32
	deltaZ      float32
	deltaClock  float32
	CorrAgeLT   float32
	IonoPPlat   float32
	IonoPPlon   float32
	SlantIono   float32
	CorrAgeIono float32
	VarFLT      float32
	VarUIRE     float32
	VarAir      float32
	VarTropo    float32
}

/** GEOCorrections_1_0_t */
type GEOCorrections_1_0_t struct {
	Header BlockHeader_t

	/* Time Header */
	TOW uint32
	WNc uint16

	N              uint8
	SBLength       uint8
	GeoCorrChannel [SBF_GEOCORRECTIONS_1_0_GEOCORRCHANNEL_LENGTH]SatCorr_1_0_t
}

type SatCorr_1_t SatCorr_1_0_t
type GEOCorrections_1_t GEOCorrections_1_0_t

/*--BaseVectorCart_1_0_t : --------------------------------------------------*/
/* XYZ relative position and velocity with respect to base(s) */

/** VectorInfoCart_1_0_t */
type VectorInfoCart_1_0_t struct {
	NrSV           uint8
	Error          uint8
	Mode           uint8
	Misc           uint8
	DeltaX         SBFDOUBLE
	DeltaY         SBFDOUBLE
	DeltaZ         SBFDOUBLE
	DeltaXVelocity float32
	DeltaYVelocity float32
	DeltaZVelocity float32
	Azimuth        uint16
	Elevation      int16
	ReferenceID    uint16
	CorrAge        uint16
	SignalInfo     uint32
}

/** BaseVectorCart_1_0_t */
type BaseVectorCart_1_0_t struct {
	Header BlockHeader_t

	/* Time Header */
	TOW uint32
	WNc uint16

	N          uint8
	SBLength   uint8
	VectorInfo [SBF_BASEVECTORCART_1_0_VECTORINFO_LENGTH]VectorInfoCart_1_0_t
}

type VectorInfoCart_1_t VectorInfoCart_1_0_t
type BaseVectorCart_1_t BaseVectorCart_1_0_t

/*--BaseVectorGeod_1_0_t : --------------------------------------------------*/
/* ENU relative position and velocity with respect to base(s) */

/** VectorInfoGeod_1_0_t */
type VectorInfoGeod_1_0_t struct {
	NrSV               uint8
	Error              uint8
	Mode               uint8
	Misc               uint8
	DeltaEast          SBFDOUBLE
	DeltaNorth         SBFDOUBLE
	DeltaUp            SBFDOUBLE
	DeltaEastVelocity  float32
	DeltaNorthVelocity float32
	DeltaUpVelocity    float32
	Azimuth            uint16
	Elevation          int16
	ReferenceID        uint16
	CorrAge            uint16
	SignalInfo         uint32
}

/** BaseVectorGeod_1_0_t */
type BaseVectorGeod_1_0_t struct {
	Header BlockHeader_t

	/* Time Header */
	TOW uint32
	WNc uint16

	N          uint8
	SBLength   uint8
	VectorInfo [SBF_BASEVECTORGEOD_1_0_VECTORINFO_LENGTH]VectorInfoGeod_1_0_t
}

type VectorInfoGeod_1_t VectorInfoGeod_1_0_t
type BaseVectorGeod_1_t BaseVectorGeod_1_0_t

/*--Ambiguities_1_0_t : -----------------------------------------------------*/
/* Carrier phase ambiguity states */

/** SatSigAmbiguity_1_0_t */
type SatSigAmbiguity_1_0_t struct {
	SVID            uint16
	RefSVID         uint16
	Type            uint8
	Reserved        uint8
	AmbiguityStdDev uint16
	Ambiguity       int32
}

/** Ambiguities_1_0_t */
type Ambiguities_1_0_t struct {
	Header BlockHeader_t

	/* Time Header */
	TOW uint32
	WNc uint16

	N           uint8
	SBLength    uint8
	Ambiguities [SBF_AMBIGUITIES_1_0_AMBIGUITIES_LENGTH]SatSigAmbiguity_1_0_t
}

type SatSigAmbiguity_1_t SatSigAmbiguity_1_0_t
type Ambiguities_1_t Ambiguities_1_0_t

/*--EndOfPVT_1_0_t : --------------------------------------------------------*/
/* PVT epoch marker */

/** EndOfPVT_1_0_t */
type EndOfPVT_1_0_t struct {
	Header BlockHeader_t

	/* Time Header */
	TOW uint32
	WNc uint16
}

type EndOfPVT_1_t EndOfPVT_1_0_t

/*--BaseLine_1_0_t : --------------------------------------------------------*/
/* Base-rover vector (deprecated block - not to be used) */

/** BaseLine_1_0_t */
type BaseLine_1_0_t struct {
	Header BlockHeader_t

	/* Time Header */
	TOW uint32
	WNc uint16

	BaseStationID uint16    /* 12 bits required RTCM3 */
	East          SBFDOUBLE /* 20km cannot be captured with F32_t in mm */
	North         SBFDOUBLE /* 20km cannot be captured with F32_t in mm */
	Up            SBFDOUBLE /* 20km cannot be captured with F32_t in mm */
}

type BaseLine_1_t BaseLine_1_0_t

/*==INS/GNSS Integrated Blocks===============================================*/

/*--IntPVCart_1_0_t : -------------------------------------------------------*/
/* Integrated PV in Cartesian coordinates */

/** IntPVCart_1_0_t */
type IntPVCart_1_0_t struct {
	Header BlockHeader_t

	/* Time Header */
	TOW uint32
	WNc uint16

	Mode        uint8
	Error       uint8
	Info        uint16
	NrSV        uint8
	NrAnt       uint8
	GNSSPVTMode uint8
	Datum       uint8
	GNSSage     uint16
	X           SBFDOUBLE
	Y           SBFDOUBLE
	Z           SBFDOUBLE
	Vx          float32
	Vy          float32
	Vz          float32
	COG         float32
}

type IntPVCart_1_t IntPVCart_1_0_t

/*--IntPVGeod_1_0_t : -------------------------------------------------------*/
/* Integrated PV in Geodetic coordinates */

/** IntPVGeod_1_0_t */
type IntPVGeod_1_0_t struct {
	Header BlockHeader_t

	/* Time Header */
	TOW uint32
	WNc uint16

	Mode        uint8
	Error       uint8
	Info        uint16
	NrSV        uint8
	NrAnt       uint8
	GNSSPVTMode uint8
	Datum       uint8
	GNSSage     uint16
	Lat         SBFDOUBLE
	Lon         SBFDOUBLE
	Alt         SBFDOUBLE
	Vn          float32
	Ve          float32
	Vu          float32
	COG         float32
}

type IntPVGeod_1_t IntPVGeod_1_0_t

/*--IntPosCovCart_1_0_t : ---------------------------------------------------*/
/* Integrated position covariance matrix (X, Y, Z) */

/** IntPosCovCart_1_0_t */
type IntPosCovCart_1_0_t struct {
	Header BlockHeader_t

	/* Time Header */
	TOW uint32
	WNc uint16

	Mode   uint8
	Error  uint8
	Cov_xx float32
	Cov_yy float32
	Cov_zz float32
	Cov_xy float32
	Cov_xz float32
	Cov_yz float32
}

type IntPosCovCart_1_t IntPosCovCart_1_0_t

/*--IntVelCovCart_1_0_t : ---------------------------------------------------*/
/* Integrated velocity covariance matrix (X, Y, Z) */

/** IntVelCovCart_1_0_t */
type IntVelCovCart_1_0_t struct {
	Header BlockHeader_t

	/* Time Header */
	TOW uint32
	WNc uint16

	Mode     uint8
	Error    uint8
	Cov_VxVx float32
	Cov_VyVy float32
	Cov_VzVz float32
	Cov_VxVy float32
	Cov_VxVz float32
	Cov_VyVz float32
}

type IntVelCovCart_1_t IntVelCovCart_1_0_t

/*--IntPosCovGeod_1_0_t : ---------------------------------------------------*/
/* Integrated position covariance matrix (Lat, Lon, Alt) */

/** IntPosCovGeod_1_0_t */
type IntPosCovGeod_1_0_t struct {
	Header BlockHeader_t

	/* Time Header */
	TOW uint32
	WNc uint16

	Mode       uint8
	Error      uint8
	Cov_LatLat float32
	Cov_LonLon float32
	Cov_AltAlt float32
	Cov_LatLon float32
	Cov_LatAlt float32
	Cov_LonAlt float32
}

type IntPosCovGeod_1_t IntPosCovGeod_1_0_t

/*--IntVelCovGeod_1_0_t : ---------------------------------------------------*/
/* Integrated velocity covariance matrix (North, East, Up) */

/** IntVelCovGeod_1_0_t */
type IntVelCovGeod_1_0_t struct {
	Header BlockHeader_t

	/* Time Header */
	TOW uint32
	WNc uint16

	Mode     uint8
	Error    uint8
	Cov_VnVn float32
	Cov_VeVe float32
	Cov_VuVu float32
	Cov_VnVe float32
	Cov_VnVu float32
	Cov_VeVu float32
}

type IntVelCovGeod_1_t IntVelCovGeod_1_0_t

/*--IntAttEuler_1_0_t : -----------------------------------------------------*/
/* Integrated attitude in Euler angles */

/** IntAttEuler_1_0_t */
type IntAttEuler_1_0_t struct {
	Header BlockHeader_t

	/* Time Header */
	TOW uint32
	WNc uint16

	Mode       uint8
	Error      uint8
	Info       uint16
	NrSV       uint8
	NrAnt      uint8
	Reserved   uint8 /* Reserved for future use */
	Datum      uint8
	GNSSage    uint16
	Heading    float32
	Pitch      float32
	Roll       float32
	PitchDot   float32
	RollDot    float32
	HeadingDot float32
}

/*--IntAttEuler_1_1_t : -----------------------------------------------------*/
/* Integrated attitude in Euler angles */

/** IntAttEuler_1_1_t */
type IntAttEuler_1_1_t struct {
	Header BlockHeader_t

	/* Time Header */
	TOW uint32
	WNc uint16

	Mode        uint8
	Error       uint8
	Info        uint16
	NrSV        uint8
	NrAnt       uint8
	GNSSPVTMode uint8
	Datum       uint8
	GNSSage     uint16
	Heading     float32
	Pitch       float32
	Roll        float32
	PitchDot    float32
	RollDot     float32
	HeadingDot  float32
}

type IntAttEuler_1_t IntAttEuler_1_1_t

/*--IntAttCovEuler_1_0_t : --------------------------------------------------*/
/* Integrated attitude covariance matrix of Euler angles */

/** IntAttCovEuler_1_0_t */
type IntAttCovEuler_1_0_t struct {
	Header BlockHeader_t

	/* Time Header */
	TOW uint32
	WNc uint16

	Mode           uint8
	Error          uint8
	Cov_HeadHead   float32
	Cov_PitchPitch float32
	Cov_RollRoll   float32
	Cov_HeadPitch  float32
	Cov_HeadRoll   float32
	Cov_PitchRoll  float32
}

type IntAttCovEuler_1_t IntAttCovEuler_1_0_t

/*--IntPVAAGeod_1_0_t : -----------------------------------------------------*/
/* Integrated position, velocity, acceleration and attitude */

/** IntPVAAGeod_1_0_t */
type IntPVAAGeod_1_0_t struct {
	Header BlockHeader_t

	/* Time Header */
	TOW uint32
	WNc uint16

	Mode        uint8
	Error       uint8
	Info        uint16
	GNSSPVTMode uint8
	Datum       uint8
	GNSSage     uint8
	NrSVAnt     uint8
	Reserved    uint8
	PosFine     uint8
	Lat         int32
	Lon         int32
	Alt         int32
	Vn          int32
	Ve          int32
	Vu          int32
	Ax          int16
	Ay          int16
	Az          int16
	Heading     uint16
	Pitch       int16
	Roll        int16
}

type IntPVAAGeod_1_t IntPVAAGeod_1_0_t

/*--INSNavCart_1_0_t : ------------------------------------------------------*/
/* INS solution in Cartesian coordinates */

/** INSNavCartPosStdDev_1_0_t */
type INSNavCartPosStdDev_1_0_t struct {
	XStdDev float32
	YStdDev float32
	ZStdDev float32
}

/** INSNavCartPosCov_1_0_t */
type INSNavCartPosCov_1_0_t struct {
	XYCov float32
	XZCov float32
	YZCov float32
}

/** INSNavCartAtt_1_0_t */
type INSNavCartAtt_1_0_t struct {
	Heading float32
	Pitch   float32
	Roll    float32
}

/** INSNavCartAttStdDev_1_0_t */
type INSNavCartAttStdDev_1_0_t struct {
	HeadingStdDev float32
	PitchStdDev   float32
	RollStdDev    float32
}

/** INSNavCartAttCov_1_0_t */
type INSNavCartAttCov_1_0_t struct {
	HeadingPitchCov float32
	HeadingRollCov  float32
	PitchRollCov    float32
}

/** INSNavCartVel_1_0_t */
type INSNavCartVel_1_0_t struct {
	Vx float32
	Vy float32
	Vz float32
}

/** INSNavCartVelStdDev_1_0_t */
type INSNavCartVelStdDev_1_0_t struct {
	VxStdDev float32
	VyStdDev float32
	VzStdDev float32
}

/** INSNavCartVelCov_1_0_t */
type INSNavCartVelCov_1_0_t struct {
	VxVyCov float32
	VxVzCov float32
	VyVzCov float32
}

/** INSNavCartData_1_0_t */

//
// TODO - Figure out how to model C++ union in go. Determine if we even need this
//

//  type union
//  {
// 	 INSNavCartPosStdDev_1_0_t PosStdDev;
// 	 INSNavCartPosCov_1_0_t PosCov;
// 	 INSNavCartAtt_1_0_t Att;
// 	 INSNavCartAttStdDev_1_0_t AttStdDev;
// 	 INSNavCartAttCov_1_0_t AttCov;
// 	 INSNavCartVel_1_0_t Vel;
// 	 INSNavCartVelStdDev_1_0_t VelStdDev;
// 	 INSNavCartVelCov_1_0_t VelCov;
//  } INSNavCartData_1_0_t;

type INSNavCartData_1_0_t struct {
	PosStdDev INSNavCartPosStdDev_1_0_t
	PosCov    INSNavCartPosCov_1_0_t
	Att       INSNavCartAtt_1_0_t
	AttStdDev INSNavCartAttStdDev_1_0_t
	AttCov    INSNavCartAttCov_1_0_t
	Vel       INSNavCartVel_1_0_t
	VelStdDev INSNavCartVelStdDev_1_0_t
	VelCov    INSNavCartVelCov_1_0_t
}

/** INSNavCart_1_0_t */
type INSNavCart_1_0_t struct {
	Header BlockHeader_t

	/* Time Header */
	TOW uint32
	WNc uint16

	GNSSMode       uint8
	Error          uint8
	Info           uint16
	GNSSAge        uint16
	X              SBFDOUBLE
	Y              SBFDOUBLE
	Z              SBFDOUBLE
	Accuracy       uint16
	Latency        uint16
	Datum          uint8
	Reserved       uint8
	SBList         uint16
	INSNavCartData [SBF_INSNAVCART_1_0_INSNAVCARTDATA_LENGTH]INSNavCartData_1_0_t
}

type INSNavCartPosStdDev_1_t INSNavCartPosStdDev_1_0_t
type INSNavCartPosCov_1_t INSNavCartPosCov_1_0_t
type INSNavCartAtt_1_t INSNavCartAtt_1_0_t
type INSNavCartAttStdDev_1_t INSNavCartAttStdDev_1_0_t
type INSNavCartAttCov_1_t INSNavCartAttCov_1_0_t
type INSNavCartVel_1_t INSNavCartVel_1_0_t
type INSNavCartVelStdDev_1_t INSNavCartVelStdDev_1_0_t
type INSNavCartVelCov_1_t INSNavCartVelCov_1_0_t
type INSNavCartData_1_t INSNavCartData_1_0_t
type INSNavCart_1_t INSNavCart_1_0_t

/*--INSNavGeod_1_0_t : ------------------------------------------------------*/
/* INS solution in Geodetic coordinates */

/** INSNavGeodPosStdDev_1_0_t */
type INSNavGeodPosStdDev_1_0_t struct {
	LatitudeStdDev  float32
	LongitudeStdDev float32
	HeightStdDev    float32
}

/** INSNavGeodPosCov_1_0_t */
type INSNavGeodPosCov_1_0_t struct {
	LatitudeLongitudeCov float32
	LatitudeHeightCov    float32
	LongitudeHeightCov   float32
}

/** INSNavGeodAtt_1_0_t */
type INSNavGeodAtt_1_0_t struct {
	Heading float32
	Pitch   float32
	Roll    float32
}

/** INSNavGeodAttStdDev_1_0_t */
type INSNavGeodAttStdDev_1_0_t struct {
	HeadingStdDev float32
	PitchStdDev   float32
	RollStdDev    float32
}

/** INSNavGeodAttCov_1_0_t */
type INSNavGeodAttCov_1_0_t struct {
	HeadingPitchCov float32
	HeadingRollCov  float32
	PitchRollCov    float32
}

/** INSNavGeodVel_1_0_t */
type INSNavGeodVel_1_0_t struct {
	Ve float32
	Vn float32
	Vu float32
}

/** INSNavGeodVelStdDev_1_0_t */
type INSNavGeodVelStdDev_1_0_t struct {
	VeStdDev float32
	VnStdDev float32
	VuStdDev float32
}

/** INSNavGeodVelCov_1_0_t */
type INSNavGeodVelCov_1_0_t struct {
	VeVnCov float32
	VeVuCov float32
	VnVuCov float32
}

/** INSNavGeodData_1_0_t */

//
// TODO - Figure out how to model C++ union in go. Determine if we even need this
//
//  type union
//  {
// 	 INSNavGeodPosStdDev_1_0_t PosStdDev;
// 	 INSNavGeodPosCov_1_0_t PosCov;
// 	 INSNavGeodAtt_1_0_t Att;
// 	 INSNavGeodAttStdDev_1_0_t AttStdDev;
// 	 INSNavGeodAttCov_1_0_t AttCov;
// 	 INSNavGeodVel_1_0_t Vel;
// 	 INSNavGeodVelStdDev_1_0_t VelStdDev;
// 	 INSNavGeodVelCov_1_0_t VelCov;
//  } INSNavGeodData_1_0_t;

type INSNavGeodData_1_0_t struct {
	PosStdDev INSNavGeodPosStdDev_1_0_t
	PosCov    INSNavGeodPosCov_1_0_t
	Att       INSNavGeodAtt_1_0_t
	AttStdDev INSNavGeodAttStdDev_1_0_t
	AttCov    INSNavGeodAttCov_1_0_t
	Vel       INSNavGeodVel_1_0_t
	VelStdDev INSNavGeodVelStdDev_1_0_t
	VelCov    INSNavGeodVelCov_1_0_t
}

/** INSNavGeod_1_0_t */
type INSNavGeod_1_0_t struct {
	Header BlockHeader_t

	/* Time Header */
	TOW uint32
	WNc uint16

	GNSSMode       uint8
	Error          uint8
	Info           uint16
	GNSSAge        uint16
	Latitude       SBFDOUBLE
	Longitude      SBFDOUBLE
	Height         SBFDOUBLE
	Undulation     float32
	Accuracy       uint16
	Latency        uint16
	Datum          uint8
	Reserved       uint8
	SBList         uint16
	INSNavGeodData [SBF_INSNAVGEOD_1_0_INSNAVGEODDATA_LENGTH]INSNavGeodData_1_0_t
}

type INSNavGeodPosStdDev_1_t INSNavGeodPosStdDev_1_0_t
type INSNavGeodPosCov_1_t INSNavGeodPosCov_1_0_t
type INSNavGeodAtt_1_t INSNavGeodAtt_1_0_t
type INSNavGeodAttStdDev_1_t INSNavGeodAttStdDev_1_0_t
type INSNavGeodAttCov_1_t INSNavGeodAttCov_1_0_t
type INSNavGeodVel_1_t INSNavGeodVel_1_0_t
type INSNavGeodVelStdDev_1_t INSNavGeodVelStdDev_1_0_t
type INSNavGeodVelCov_1_t INSNavGeodVelCov_1_0_t
type INSNavGeodData_1_t INSNavGeodData_1_0_t
type INSNavGeod_1_t INSNavGeod_1_0_t

/*--IMUBias_1_0_t : ---------------------------------------------------------*/
/* Estimated parameters of the IMU, such as the IMU biases and their standard deviation */

/** IMUBias_1_0_t */
type IMUBias_1_0_t struct {
	Header BlockHeader_t

	/* Time Header */
	TOW uint32
	WNc uint16

	Error        uint8
	Reserved1    uint8
	GNSSAge      uint16
	Reserved2    uint16
	XAccBias     float32
	YAccBias     float32
	ZAccBias     float32
	XGyroBias    float32
	YGyroBias    float32
	ZGyroBias    float32
	StdXAccBias  float32
	StdYAccBias  float32
	StdZAccBias  float32
	StdXGyroBias float32
	StdYGyroBias float32
	StdZGyroBias float32
}

type IMUBias_1_t IMUBias_1_0_t

/*==GNSS Attitude Blocks=====================================================*/

/*--AttEuler_1_0_t : --------------------------------------------------------*/
/* GNSS attitude expressed as Euler angles */

/** AttEuler_1_0_t */
type AttEuler_1_0_t struct {
	Header BlockHeader_t

	/* Time Header */
	TOW uint32
	WNc uint16

	NRSV       uint8
	Error      uint8
	Mode       uint16
	Reserved   uint16
	Heading    float32
	Pitch      float32
	Roll       float32
	PitchDot   float32
	RollDot    float32
	HeadingDot float32
}

type AttEuler_1_t AttEuler_1_0_t

/*--AttCovEuler_1_0_t : -----------------------------------------------------*/
/* Covariance matrix of attitude */

/** AttCovEuler_1_0_t */
type AttCovEuler_1_0_t struct {
	Header BlockHeader_t

	/* Time Header */
	TOW uint32
	WNc uint16

	Reserved       uint8
	Error          uint8
	Cov_HeadHead   float32
	Cov_PitchPitch float32
	Cov_RollRoll   float32
	Cov_HeadPitch  float32
	Cov_HeadRoll   float32
	Cov_PitchRoll  float32
}

type AttCovEuler_1_t AttCovEuler_1_0_t

/*--AuxAntPositions_1_0_t : -------------------------------------------------*/
/* Relative position and velocity estimates of auxiliary antennas */

/** AuxAntPositionSub_1_0_t */
type AuxAntPositionSub_1_0_t struct {
	NRSV          uint8
	Error         uint8
	AmbiguityType uint8
	AuxAntID      uint8
	DeltaEast     SBFDOUBLE
	DeltaNorth    SBFDOUBLE
	DeltaUp       SBFDOUBLE
	EastVelocity  SBFDOUBLE
	NorthVelocity SBFDOUBLE
	UpVelocity    SBFDOUBLE
}

/** AuxAntPositions_1_0_t */
type AuxAntPositions_1_0_t struct {
	Header BlockHeader_t

	/* Time Header */
	TOW uint32
	WNc uint16

	NbrAuxAntennas  uint8
	SBSize          uint8
	AuxAntPositions [SBF_AUXANTPOSITIONS_1_0_AUXANTPOSITIONS_LENGTH]AuxAntPositionSub_1_0_t
}

type AuxAntPositionSub_1_t AuxAntPositionSub_1_0_t
type AuxAntPositions_1_t AuxAntPositions_1_0_t

/*--EndOfAtt_1_0_t : --------------------------------------------------------*/
/* GNSS attitude epoch marker */

/** EndOfAtt_1_0_t */
type EndOfAtt_1_0_t struct {
	Header BlockHeader_t

	/* Time Header */
	TOW uint32
	WNc uint16
}

type EndOfAtt_1_t EndOfAtt_1_0_t

/*--AttQuat_1_0_t : ---------------------------------------------------------*/

/** AttQuat_1_0_t */
type AttQuat_1_0_t struct {
	Header BlockHeader_t

	/* Time Header */
	TOW uint32
	WNc uint16

	NRSV       uint8
	Error      uint8
	Mode       uint16
	Reserved   uint16
	q1         float32
	q2         float32
	q3         float32
	q4         float32
	PitchDot   float32
	RollDot    float32
	HeadingDot float32
}

type AttQuat_1_t AttQuat_1_0_t

/** AttCovQuat_1_0_t */
type AttCovQuat_1_0_t struct {
	Header BlockHeader_t

	/* Time Header */
	TOW uint32
	WNc uint16

	Reserved uint8
	Error    uint8
	Cov_q1q1 float32
	Cov_q2q2 float32
	Cov_q3q3 float32
	Cov_q4q4 float32
	Cov_q1q2 float32
	Cov_q1q3 float32
	Cov_q1q4 float32
	Cov_q2q3 float32
	Cov_q2q4 float32
	Cov_q3q4 float32
}

type AttCovQuat_1_t AttCovQuat_1_0_t

/*==Receiver Time Blocks=====================================================*/

/*--ReceiverTime_1_0_t : ----------------------------------------------------*/
/* Current receiver and UTC time */

/** ReceiverTime_1_0_t */
type ReceiverTime_1_0_t struct {
	Header BlockHeader_t

	/* Time Header */
	TOW uint32
	WNc uint16

	UTCYear   int8
	UTCMonth  int8
	UTCDay    int8
	UTCHour   int8
	UTCMin    int8
	UTCSec    int8
	DeltaLS   int8
	SyncLevel uint8
	_padding  [SBF_RECEIVERTIME_1_0__PADDING_LENGTH]uint8
}

type ReceiverTime_1_t ReceiverTime_1_0_t

/*--xPPSOffset_1_0_t : ------------------------------------------------------*/
/* Offset of the xPPS pulse with respect to GNSS time */

/** PPSData_1_0_t */
type PPSData_1_0_t struct {
	/* Time Header */
	TOW uint32
	WNc uint16

	SyncAge uint8
	Source  uint8
	Offset  float32
}

/** xPPSOffset_1_0_t */
type xPPSOffset_1_0_t struct {
	Header  BlockHeader_t
	PPSData PPSData_1_0_t
}

type PPSData_1_t PPSData_1_0_t
type xPPSOffset_1_t xPPSOffset_1_0_t

/*--SysTimeOffset_1_0_t : ---------------------------------------------------*/
/* Time offset between different constellations */

/** TimeOffsetSub_1_0_t */
type TimeOffsetSub_1_0_t struct {
	TimeSystem uint8
	Offset     float32
	_padding   [SBF_TIMEOFFSETSUB_1_0__PADDING_LENGTH]uint8
}

/** SysTimeOffset_1_0_t */
type SysTimeOffset_1_0_t struct {
	Header BlockHeader_t

	/* Time Header */
	TOW uint32
	WNc uint16

	RefTimeSystem uint8
	N             uint8
	SBLength      uint8
	Reserved2     uint8    /* Reserved for future use */
	Reserved3     [2]uint8 /* Reserved for future use */
	TimeOffset    [SBF_SYSTIMEOFFSET_1_0_TIMEOFFSET_LENGTH]TimeOffsetSub_1_0_t
}

/*--SysTimeOffset_1_1_t : ---------------------------------------------------*/
/* Time offset between different constellations */

/** TimeOffsetSub_1_1_t */
type TimeOffsetSub_1_1_t struct {
	TimeSystem uint8
	Mode       uint8
	Reserved   [2]uint8
	Offset     float32
}

/** SysTimeOffset_1_1_t */
type SysTimeOffset_1_1_t struct {
	Header BlockHeader_t

	/* Time Header */
	TOW uint32
	WNc uint16

	RefTimeSystem uint8
	N             uint8
	SBLength      uint8
	Alert_flags   uint8
	Reserved      [2]uint8
	TimeOffset    [SBF_SYSTIMEOFFSET_1_1_TIMEOFFSET_LENGTH]TimeOffsetSub_1_1_t
}

type TimeOffsetSub_1_t TimeOffsetSub_1_1_t
type SysTimeOffset_1_t SysTimeOffset_1_1_t

/*==External Event Blocks====================================================*/

/*--ExtEvent_1_0_t : --------------------------------------------------------*/
/* Time at the instant of an external event */

/** TimerData_1_0_t */
type TimerData_1_0_t struct {
	/* Time Header */
	TOW uint32
	WNc uint16

	Source    uint8
	Polarity  uint8
	Offset    float32
	RxClkBias SBFDOUBLE
}

/** ExtEvent_1_0_t */
type ExtEvent_1_0_t struct {
	Header    BlockHeader_t
	TimerData TimerData_1_0_t
}

/*--ExtEvent_1_1_t : --------------------------------------------------------*/
/* Time at the instant of an external event */

/** TimerData_1_1_t */
type TimerData_1_1_t struct {
	/* Time Header */
	TOW uint32
	WNc uint16

	Source    uint8
	Polarity  uint8
	Offset    float32
	RxClkBias SBFDOUBLE
	PVTAge    uint16
}

/** ExtEvent_1_1_t */
type ExtEvent_1_1_t struct {
	Header BlockHeader_t

	TimerData TimerData_1_1_t
}

type TimerData_1_t TimerData_1_1_t
type ExtEvent_1_t ExtEvent_1_1_t

/*--ExtEventPVTCartesian_1_0_t : --------------------------------------------*/
/* Cartesian position at the instant of an event */

/** ExtEventPVTCartesian_1_0_t */
type ExtEventPVTCartesian_1_0_t struct {
	Header BlockHeader_t

	/* Time Header */
	TOW uint32
	WNc uint16

	Mode        uint8
	Error       uint8
	X           SBFDOUBLE
	Y           SBFDOUBLE
	Z           SBFDOUBLE
	Undulation  float32
	Vx          float32
	Vy          float32
	Vz          float32
	COG         float32
	RxClkBias   SBFDOUBLE
	RxClkDrift  float32
	TimeSystem  uint8
	Datum       uint8
	NrSV        uint8
	WACorrInfo  uint8
	ReferenceId uint16
	MeanCorrAge uint16
	SignalInfo  uint32
	AlertFlag   uint8
	_padding    [SBF_EXTEVENTPVTCARTESIAN_1_0__PADDING_LENGTH]uint8
}

/*--ExtEventPVTCartesian_1_1_t : --------------------------------------------*/
/* Cartesian position at the instant of an event */

/** ExtEventPVTCartesian_1_1_t */
type ExtEventPVTCartesian_1_1_t struct {
	Header BlockHeader_t

	/* Time Header */
	TOW uint32
	WNc uint16

	Mode        uint8
	Error       uint8
	X           SBFDOUBLE
	Y           SBFDOUBLE
	Z           SBFDOUBLE
	Undulation  float32
	Vx          float32
	Vy          float32
	Vz          float32
	COG         float32
	RxClkBias   SBFDOUBLE
	RxClkDrift  float32
	TimeSystem  uint8
	Datum       uint8
	NrSV        uint8
	WACorrInfo  uint8
	ReferenceId uint16
	MeanCorrAge uint16
	SignalInfo  uint32
	AlertFlag   uint8
	NrBases     uint8
	PPPInfo     uint16
}

/*--ExtEventPVTCartesian_1_2_t : --------------------------------------------*/
/* Cartesian position at the instant of an event */
/** ExtEventPVTCartesian_1_2_t */
type ExtEventPVTCartesian_1_2_t struct {
	Header BlockHeader_t

	/* Time Header */
	TOW uint32
	WNc uint16

	Mode        uint8
	Error       uint8
	X           SBFDOUBLE
	Y           SBFDOUBLE
	Z           SBFDOUBLE
	Undulation  float32
	Vx          float32
	Vy          float32
	Vz          float32
	COG         float32
	RxClkBias   SBFDOUBLE
	RxClkDrift  float32
	TimeSystem  uint8
	Datum       uint8
	NrSV        uint8
	WACorrInfo  uint8
	ReferenceId uint16
	MeanCorrAge uint16
	SignalInfo  uint32
	AlertFlag   uint8
	NrBases     uint8
	PPPInfo     uint16
	Latency     uint16
	HAccuracy   uint16
	VAccuracy   uint16
	Misc        uint8
	_padding    [SBF_EXTEVENTPVTCARTESIAN_1_2__PADDING_LENGTH]uint8
}

type ExtEventPVTCartesian_1_t ExtEventPVTCartesian_1_2_t

/*--ExtEventPVTGeodetic_1_0_t : ---------------------------------------------*/
/* Geodetic position at the instant of an event */

/** ExtEventPVTGeodetic_1_0_t */
type ExtEventPVTGeodetic_1_0_t struct {
	Header BlockHeader_t

	/* Time Header */
	TOW uint32
	WNc uint16

	Mode        uint8
	Error       uint8
	Lat         SBFDOUBLE
	Lon         SBFDOUBLE
	Alt         SBFDOUBLE
	Undulation  float32
	Vn          float32
	Ve          float32
	Vu          float32
	COG         float32
	RxClkBias   SBFDOUBLE
	RxClkDrift  float32
	TimeSystem  uint8
	Datum       uint8
	NrSV        uint8
	WACorrInfo  uint8
	ReferenceId uint16
	MeanCorrAge uint16
	SignalInfo  uint32
	AlertFlag   uint8
	_padding    [SBF_EXTEVENTPVTGEODETIC_1_0__PADDING_LENGTH]uint8
}

/*--ExtEventPVTGeodetic_1_1_t : ---------------------------------------------*/
/* Geodetic position at the instant of an event */

/** ExtEventPVTGeodetic_1_1_t */
type ExtEventPVTGeodetic_1_1_t struct {
	Header BlockHeader_t

	/* Time Header */
	TOW uint32
	WNc uint16

	Mode        uint8
	Error       uint8
	Lat         SBFDOUBLE
	Lon         SBFDOUBLE
	Alt         SBFDOUBLE
	Undulation  float32
	Vn          float32
	Ve          float32
	Vu          float32
	COG         float32
	RxClkBias   SBFDOUBLE
	RxClkDrift  float32
	TimeSystem  uint8
	Datum       uint8
	NrSV        uint8
	WACorrInfo  uint8
	ReferenceId uint16
	MeanCorrAge uint16
	SignalInfo  uint32
	AlertFlag   uint8
	NrBases     uint8
	PPPInfo     uint16
}

/*--ExtEventPVTGeodetic_1_2_t : ---------------------------------------------*/
/* Geodetic position at the instant of an event */

/** ExtEventPVTGeodetic_1_2_t */
type ExtEventPVTGeodetic_1_2_t struct {
	Header BlockHeader_t

	/* Time Header */
	TOW uint32
	WNc uint16

	Mode        uint8
	Error       uint8
	Lat         SBFDOUBLE
	Lon         SBFDOUBLE
	Alt         SBFDOUBLE
	Undulation  float32
	Vn          float32
	Ve          float32
	Vu          float32
	COG         float32
	RxClkBias   SBFDOUBLE
	RxClkDrift  float32
	TimeSystem  uint8
	Datum       uint8
	NrSV        uint8
	WACorrInfo  uint8
	ReferenceId uint16
	MeanCorrAge uint16
	SignalInfo  uint32
	AlertFlag   uint8
	NrBases     uint8
	PPPInfo     uint16
	Latency     uint16
	HAccuracy   uint16
	VAccuracy   uint16
	Misc        uint8
	_padding    [SBF_EXTEVENTPVTGEODETIC_1_2__PADDING_LENGTH]uint8
}

type ExtEventPVTGeodetic_1_t ExtEventPVTGeodetic_1_2_t

/*--ExtEventBaseVectCart_1_0_t : --------------------------------------------*/
/* XYZ relative position with respect to base(s) at the instant of an event */

/** ExtEventVectorInfoCart_1_0_t */
type ExtEventVectorInfoCart_1_0_t struct {
	NrSV        uint8
	Error       uint8
	Mode        uint8
	Misc        uint8
	DeltaX      SBFDOUBLE
	DeltaY      SBFDOUBLE
	DeltaZ      SBFDOUBLE
	DeltaVx     float32
	DeltaVy     float32
	DeltaVz     float32
	Azimuth     uint16
	Elevation   int16
	ReferenceID uint16
	CorrAge     uint16
	SignalInfo  uint32
}

/** ExtEventBaseVectCart_1_0_t */
type ExtEventBaseVectCart_1_0_t struct {
	Header BlockHeader_t

	/* Time Header */
	TOW uint32
	WNc uint16

	N                      uint8
	SBLength               uint8
	ExtEventVectorInfoCart [SBF_EXTEVENTBASEVECTCART_1_0_EXTEVENTVECTORINFOCART_LENGTH]ExtEventVectorInfoCart_1_0_t
}

type ExtEventVectorInfoCart_1_t ExtEventVectorInfoCart_1_0_t
type ExtEventBaseVectCart_1_t ExtEventBaseVectCart_1_0_t

/*--ExtEventBaseVectGeod_1_0_t : --------------------------------------------*/
/* ENU relative position with respect to base(s) at the instant of an event */

/** ExtEventVectorInfoGeod_1_0_t */
type ExtEventVectorInfoGeod_1_0_t struct {
	NrSV        uint8
	Error       uint8
	Mode        uint8
	Misc        uint8
	DeltaEast   SBFDOUBLE
	DeltaNorth  SBFDOUBLE
	DeltaUp     SBFDOUBLE
	DeltaVe     float32
	DeltaVn     float32
	DeltaVu     float32
	Azimuth     uint16
	Elevation   int16
	ReferenceID uint16
	CorrAge     uint16
	SignalInfo  uint32
}

/** ExtEventBaseVectGeod_1_0_t */
type ExtEventBaseVectGeod_1_0_t struct {
	Header BlockHeader_t

	/* Time Header */
	TOW uint32
	WNc uint16

	N                      uint8
	SBLength               uint8
	ExtEventVectorInfoGeod [SBF_EXTEVENTBASEVECTGEOD_1_0_EXTEVENTVECTORINFOGEOD_LENGTH]ExtEventVectorInfoGeod_1_0_t
}

type ExtEventVectorInfoGeod_1_t ExtEventVectorInfoGeod_1_0_t
type ExtEventBaseVectGeod_1_t ExtEventBaseVectGeod_1_0_t

/*--ExtEventINSNavCart_1_0_t : ----------------------------------------------*/
/* INS solution in Cartesian coordinates at the instant of an event */

/** ExtEventINSNavCartPosStdDev_1_0_t */
type ExtEventINSNavCartPosStdDev_1_0_t struct {
	XStdDev float32
	YStdDev float32
	ZStdDev float32
}

/** ExtEventINSNavCartAtt_1_0_t */
type ExtEventINSNavCartAtt_1_0_t struct {
	Heading float32
	Pitch   float32
	Roll    float32
}

/** ExtEventINSNavCartAttStdDev_1_0_t */
type ExtEventINSNavCartAttStdDev_1_0_t struct {
	HeadingStdDev float32
	PitchStdDev   float32
	RollStdDev    float32
}

/** ExtEventINSNavCartVel_1_0_t */
type ExtEventINSNavCartVel_1_0_t struct {
	Vx float32
	Vy float32
	Vz float32
}

/** ExtEventINSNavCartVelStdDev_1_0_t */
type ExtEventINSNavCartVelStdDev_1_0_t struct {
	VxStdDev float32
	VyStdDev float32
	VzStdDev float32
}

/** ExtEventINSNavCartData_1_0_t */

//
// TODO - Figure out how to model C++ union in go. Determine if we even need this
//
// type union
// {
// 	ExtEventINSNavCartPosStdDev_1_0_t PosStdDev;
// 	ExtEventINSNavCartAtt_1_0_t Att;
// 	ExtEventINSNavCartAttStdDev_1_0_t AttStdDev;
// 	ExtEventINSNavCartVel_1_0_t Vel;
// 	ExtEventINSNavCartVelStdDev_1_0_t VelStdDev;
// } ExtEventINSNavCartData_1_0_t;

/** ExtEventINSNavCartData_1_0_t */
type ExtEventINSNavCartData_1_0_t struct {
	PosStdDev ExtEventINSNavCartPosStdDev_1_0_t
	Att       ExtEventINSNavCartAtt_1_0_t
	AttStdDev ExtEventINSNavCartAttStdDev_1_0_t
	Vel       ExtEventINSNavCartVel_1_0_t
	VelStdDev ExtEventINSNavCartVelStdDev_1_0_t
}

/** ExtEventINSNavCart_1_0_t */
type ExtEventINSNavCart_1_0_t struct {
	Header BlockHeader_t

	/* Time Header */
	TOW uint32
	WNc uint16

	GNSSMode               uint8
	Error                  uint8
	Info                   uint16
	GNSSAge                uint16
	X                      SBFDOUBLE
	Y                      SBFDOUBLE
	Z                      SBFDOUBLE
	Accuracy               uint16
	Reserved2              uint16
	Datum                  uint8
	Reserved               uint8
	SBList                 uint16
	ExtEventINSNavCartData [SBF_EXTEVENTINSNAVCART_1_0_EXTEVENTINSNAVCARTDATA_LENGTH]ExtEventINSNavCartData_1_0_t
}

type ExtEventINSNavCartPosStdDev_1_t ExtEventINSNavCartPosStdDev_1_0_t
type ExtEventINSNavCartAtt_1_t ExtEventINSNavCartAtt_1_0_t
type ExtEventINSNavCartAttStdDev_1_t ExtEventINSNavCartAttStdDev_1_0_t
type ExtEventINSNavCartVel_1_t ExtEventINSNavCartVel_1_0_t
type ExtEventINSNavCartVelStdDev_1_t ExtEventINSNavCartVelStdDev_1_0_t
type ExtEventINSNavCartData_1_t ExtEventINSNavCartData_1_0_t
type ExtEventINSNavCart_1_t ExtEventINSNavCart_1_0_t

/*--ExtEventINSNavGeod_1_0_t : ----------------------------------------------*/
/* INS solution in Geodetic coordinates at the instant of an event */

/** ExtEventINSNavGeodPosStdDev_1_0_t */
type ExtEventINSNavGeodPosStdDev_1_0_t struct {
	LatitudeStdDev  float32
	LongitudeStdDev float32
	HeightStdDev    float32
}

/** ExtEventINSNavGeodAtt_1_0_t */
type ExtEventINSNavGeodAtt_1_0_t struct {
	Heading float32
	Pitch   float32
	Roll    float32
}

/** ExtEventINSNavGeodAttStdDev_1_0_t */
type ExtEventINSNavGeodAttStdDev_1_0_t struct {
	HeadingStdDev float32
	PitchStdDev   float32
	RollStdDev    float32
}

/** ExtEventINSNavGeodVel_1_0_t */
type ExtEventINSNavGeodVel_1_0_t struct {
	Ve float32
	Vn float32
	Vu float32
}

/** ExtEventINSNavGeodVelStdDev_1_0_t */
type ExtEventINSNavGeodVelStdDev_1_0_t struct {
	VeStdDev float32
	VnStdDev float32
	VuStdDev float32
}

/** ExtEventINSNavGeodData_1_0_t */

//
// TODO - Figure out how to model C++ union in go. Determine if we even need this
//
//  type union
//  {
// 	 ExtEventINSNavGeodPosStdDev_1_0_t PosStdDev;
// 	 ExtEventINSNavGeodAtt_1_0_t Att;
// 	 ExtEventINSNavGeodAttStdDev_1_0_t AttStdDev;
// 	 ExtEventINSNavGeodVel_1_0_t Vel;
// 	 ExtEventINSNavGeodVelStdDev_1_0_t VelStdDev;
//  } ExtEventINSNavGeodData_1_0_t;

type ExtEventINSNavGeodData_1_0_t struct {
	PosStdDev ExtEventINSNavGeodPosStdDev_1_0_t
	Att       ExtEventINSNavGeodAtt_1_0_t
	AttStdDev ExtEventINSNavGeodAttStdDev_1_0_t
	Vel       ExtEventINSNavGeodVel_1_0_t
	VelStdDev ExtEventINSNavGeodVelStdDev_1_0_t
}

/** ExtEventINSNavGeod_1_0_t */
type ExtEventINSNavGeod_1_0_t struct {
	Header BlockHeader_t

	/* Time Header */
	TOW uint32
	WNc uint16

	GNSSMode               uint8
	Error                  uint8
	Info                   uint16
	GNSSAge                uint16
	Latitude               SBFDOUBLE
	Longitude              SBFDOUBLE
	Height                 SBFDOUBLE
	Undulation             float32
	Accuracy               uint16
	Reserved2              uint16
	Datum                  uint8
	Reserved               uint8
	SBList                 uint16
	ExtEventINSNavGeodData [SBF_EXTEVENTINSNAVGEOD_1_0_EXTEVENTINSNAVGEODDATA_LENGTH]ExtEventINSNavGeodData_1_0_t
}

type ExtEventINSNavGeodPosStdDev_1_t ExtEventINSNavGeodPosStdDev_1_0_t
type ExtEventINSNavGeodAtt_1_t ExtEventINSNavGeodAtt_1_0_t
type ExtEventINSNavGeodAttStdDev_1_t ExtEventINSNavGeodAttStdDev_1_0_t
type ExtEventINSNavGeodVel_1_t ExtEventINSNavGeodVel_1_0_t
type ExtEventINSNavGeodVelStdDev_1_t ExtEventINSNavGeodVelStdDev_1_0_t
type ExtEventINSNavGeodData_1_t ExtEventINSNavGeodData_1_0_t
type ExtEventINSNavGeod_1_t ExtEventINSNavGeod_1_0_t

/*--ExtEventAttEuler_1_0_t : ------------------------------------------------*/
/* GNSS attitude expressed as Euler angles at the instant of an event */

/** ExtEventAttEuler_1_0_t */
type ExtEventAttEuler_1_0_t struct {
	Header BlockHeader_t

	/* Time Header */
	TOW uint32
	WNc uint16

	NrSV       uint8
	Error      uint8
	Mode       uint16
	Reserved   uint16
	Heading    float32
	Pitch      float32
	Roll       float32
	PitchDot   float32
	RollDot    float32
	HeadingDot float32
}

type ExtEventAttEuler_1_t ExtEventAttEuler_1_0_t

/*==Differential Correction Blocks===========================================*/

/*--DiffCorrIn_1_0_t : ------------------------------------------------------*/
/* Incoming RTCM or CMR message */

/** DiffCorrIn_1_0_t */
type DiffCorrIn_1_0_t struct {
	Header BlockHeader_t

	/* Time Header */
	TOW uint32
	WNc uint16

	Mode   uint8
	Source uint8
	Frame  [SBF_DIFFCORRIN_1_0_FRAME_LENGTH]uint8
}
type DiffCorrIn_1_t DiffCorrIn_1_0_t

/*--BaseStation_1_0_t : -----------------------------------------------------*/
/* Base station coordinates */

/** BaseStation_1_0_t */
type BaseStation_1_0_t struct {
	Header BlockHeader_t

	/* Time Header */
	TOW uint32
	WNc uint16

	BaseStationID   uint16
	BaseType        uint8
	Source          uint8
	Datum           uint8
	Reserved        uint8
	X_L1PhaseCenter SBFDOUBLE
	Y_L1PhaseCenter SBFDOUBLE
	Z_L1PhaseCenter SBFDOUBLE
}

type BaseStation_1_t BaseStation_1_0_t

/*--RTCMDatum_1_0_t : -------------------------------------------------------*/
/* Datum information from the RTK service provider */

/** RTCMDatum_1_0_t */
type RTCMDatum_1_0_t struct {
	Header BlockHeader_t

	/* Time Header */
	TOW uint32
	WNc uint16

	SourceCRS  [SBF_RTCMDATUM_1_0_SOURCECRS_LENGTH]byte
	TargetCRS  [SBF_RTCMDATUM_1_0_TARGETCRS_LENGTH]byte
	Datum      uint8
	HeightType uint8
	QualityInd uint8
}
type RTCMDatum_1_t RTCMDatum_1_0_t

/*--BaseLink_1_0_t : --------------------------------------------------------*/

/** BaseLink_1_0_t */
type BaseLink_1_0_t struct {
	Header BlockHeader_t

	/* Time Header */
	TOW uint32
	WNc uint16

	CorrAvailable      uint8
	Reserved           uint8
	NrBytesReceived    uint32
	NrBytesAccepted    uint32
	NrMessagesReceived uint32
	NrMessagesAccepted uint32
	AgeOfLastMsg       float32
}

type BaseLink_1_t BaseLink_1_0_t

/*==L-Band Demodulator Blocks================================================*/

/*--LBandReceiverStatus_1_0_t : ---------------------------------------------*/

/** LBandReceiverStatus_1_0_t */
type LBandReceiverStatus_1_0_t struct {
	Header BlockHeader_t

	/* Time Header */
	TOW uint32
	WNc uint16

	CPULoad  uint8
	Reserved uint8
	UpTime   uint32
	RxStatus uint32
	RxError  uint32
}

type LBandReceiverStatus_1_t LBandReceiverStatus_1_0_t

/*--LBandTrackerStatus_1_0_t : ----------------------------------------------*/
/* Status of the L-band signal tracking */

/** TrackData_1_0_t */
type TrackData_1_0_t struct {
	Frequency  uint32
	BaudRate   uint16
	ServiceID  uint16
	FreqOffset float32
	CN0        uint16
	AvgPower   int16
	AGCGain    int8
	Mode       uint8
	Status     uint8
	Reserved   uint8 /* Reserved for future use */
}

/** LBandTrackerStatus_1_0_t */
type LBandTrackerStatus_1_0_t struct {
	Header BlockHeader_t

	/* Time Header */
	TOW uint32
	WNc uint16

	N         uint8
	SBSize    uint8
	TrackData [SBF_LBANDTRACKERSTATUS_1_0_TRACKDATA_LENGTH]TrackData_1_0_t
}

/*--LBandTrackerStatus_1_1_t : ----------------------------------------------*/
/* Status of the L-band signal tracking */

/** TrackData_1_1_t */
type TrackData_1_1_t struct {
	Frequency  uint32
	BaudRate   uint16
	ServiceID  uint16
	FreqOffset float32
	CN0        uint16
	AvgPower   int16
	AGCGain    int8
	Mode       uint8
	Status     uint8
	Reserved   uint8 /* Reserved for future use */
	LockTime   uint16
	_padding   [SBF_TRACKDATA_1_1__PADDING_LENGTH]uint8
}

/** LBandTrackerStatus_1_1_t */
type LBandTrackerStatus_1_1_t struct {
	Header BlockHeader_t

	/* Time Header */
	TOW uint32
	WNc uint16

	N         uint8
	SBSize    uint8
	TrackData [SBF_LBANDTRACKERSTATUS_1_1_TRACKDATA_LENGTH]TrackData_1_1_t
}

/*--LBandTrackerStatus_1_2_t : ----------------------------------------------*/
/* Status of the L-band signal tracking */

/** TrackData_1_2_t */
type TrackData_1_2_t struct {
	Frequency  uint32
	BaudRate   uint16
	ServiceID  uint16
	FreqOffset float32
	CN0        uint16
	AvgPower   int16
	AGCGain    int8
	Mode       uint8
	Status     uint8
	SVID       uint8
	LockTime   uint16
	_padding   [SBF_TRACKDATA_1_2__PADDING_LENGTH]uint8
}

/** LBandTrackerStatus_1_2_t */
type LBandTrackerStatus_1_2_t struct {
	Header BlockHeader_t

	/* Time Header */
	TOW uint32
	WNc uint16

	N         uint8
	SBSize    uint8
	TrackData [SBF_LBANDTRACKERSTATUS_1_2_TRACKDATA_LENGTH]TrackData_1_2_t
}

/*--LBandTrackerStatus_1_3_t : ----------------------------------------------*/
/* Status of the L-band signal tracking */
/** TrackData_1_3_t */
type TrackData_1_3_t struct {
	Frequency  uint32
	BaudRate   uint16
	ServiceID  uint16
	FreqOffset float32
	CN0        uint16
	AvgPower   int16
	AGCGain    int8
	Mode       uint8
	Status     uint8
	SVID       uint8
	LockTime   uint16
	Source     uint8
	_padding   [SBF_TRACKDATA_1_3__PADDING_LENGTH]uint8
}

/** LBandTrackerStatus_1_3_t */
type LBandTrackerStatus_1_3_t struct {
	Header BlockHeader_t

	/* Time Header */
	TOW uint32
	WNc uint16

	N         uint8
	SBSize    uint8
	TrackData [SBF_LBANDTRACKERSTATUS_1_3_TRACKDATA_LENGTH]TrackData_1_3_t
}

type TrackData_1_t TrackData_1_3_t
type LBandTrackerStatus_1_t LBandTrackerStatus_1_3_t

/*--LBAS1DecoderStatus_1_0_t : ----------------------------------------------*/
/* Status of the LBAS1 L-band service */

/** LBAS1DecoderStatus_1_0_t */
type LBAS1DecoderStatus_1_0_t struct {
	Header BlockHeader_t

	/* Time Header */
	TOW uint32
	WNc uint16

	Reserved        [2]uint8
	Status          uint8
	Access          uint8
	GeoGatingMode   uint8
	GeoGatingStatus uint8
	Event           uint32
}

/*--LBAS1DecoderStatus_1_1_t : ----------------------------------------------*/
/* Status of the LBAS1 L-band service */

/** LBAS1DecoderStatus_1_1_t */
type LBAS1DecoderStatus_1_1_t struct {
	Header BlockHeader_t

	/* Time Header */
	TOW uint32
	WNc uint16

	Reserved        [2]uint8
	Status          uint8
	Access          uint8
	GeoGatingMode   uint8
	GeoGatingStatus uint8
	Event           uint32
	LeaseTime       uint32
	LeaseRemaining  uint32
	LocalAreaLat    int32
	LocalAreaLon    int32
	LocalAreaRadius uint16
	LocalAreaStatus uint8
	Reserved1       uint8
	SubscrEndYear   int8
	SubscrEndMonth  int8
	SubscrEndDay    int8
	SubscrEndHour   int8
	PAC             [SBF_LBAS1DECODERSTATUS_1_1_PAC_LENGTH]byte
}

/*--LBAS1DecoderStatus_1_2_t : ----------------------------------------------*/
/* Status of the LBAS1 L-band service */

/** LBAS1DecoderStatus_1_2_t */
type LBAS1DecoderStatus_1_2_t struct {
	Header BlockHeader_t

	/* Time Header */
	TOW uint32
	WNc uint16

	Reserved          [2]uint8
	Status            uint8
	Access            uint8
	GeoGatingMode     uint8
	GeoGatingStatus   uint8
	Event             uint32
	LeaseTime         uint32
	LeaseRemaining    uint32
	LocalAreaLat      int32
	LocalAreaLon      int32
	LocalAreaRadius   uint16
	LocalAreaStatus   uint8
	Reserved1         uint8
	SubscrEndYear     int8
	SubscrEndMonth    int8
	SubscrEndDay      int8
	SubscrEndHour     int8
	PAC               [SBF_LBAS1DECODERSTATUS_1_2_PAC_LENGTH]byte
	VelocityLimit     uint8
	SpeedGatingStatus uint8
	_padding          [SBF_LBAS1DECODERSTATUS_1_2__PADDING_LENGTH]uint8
}

type LBAS1DecoderStatus_1_t LBAS1DecoderStatus_1_2_t

/*--LBAS1Messages_1_0_t : ---------------------------------------------------*/
/* LBAS1over-the-air message */

/** LBAS1Messages_1_0_t */
type LBAS1Messages_1_0_t struct {
	Header BlockHeader_t

	/* Time Header */
	TOW uint32
	WNc uint16

	MessageLen uint16
	Message    [SBF_LBAS1MESSAGES_1_0_MESSAGE_LENGTH]byte
}

type LBAS1Messages_1_t LBAS1Messages_1_0_t
type LBAS1Message_1_0_t LBAS1Messages_1_0_t
type LBAS1Message_1_t LBAS1Messages_1_t

/*--LBandBeams_1_0_t : ------------------------------------------------------*/
/* L-band satellite/beam information */

/** BeamInfo_1_0_t */
type BeamInfo_1_0_t struct {
	SVID         uint8
	SatName      [SBF_BEAMINFO_1_0_SATNAME_LENGTH]byte
	SatLongitude int16
	BeamFreq     uint32
}

/** LBandBeams_1_0_t */
type LBandBeams_1_0_t struct {
	Header BlockHeader_t

	/* Time Header */
	TOW uint32
	WNc uint16

	N        uint8
	SBLength uint8
	BeamData [SBF_LBANDBEAMS_1_0_BEAMDATA_LENGTH]BeamInfo_1_0_t
}

type BeamInfo_1_t BeamInfo_1_0_t
type LBandBeams_1_t LBandBeams_1_0_t
type LBandBeamInfo_1_0_t BeamInfo_1_0_t
type LBandBeamInfo_1_t BeamInfo_1_t

/*--LBandRaw_1_0_t : --------------------------------------------------------*/
/* L-Band raw user data */

/** LBandRaw_1_0_t */
type LBandRaw_1_0_t struct {
	Header BlockHeader_t

	/* Time Header */
	TOW uint32
	WNc uint16

	N         uint16
	Frequency uint32
	UserData  [SBF_LBANDRAW_1_0_USERDATA_LENGTH]uint8
}
type LBandRaw_1_t LBandRaw_1_0_t

/*--FugroStatus_1_0_t : -----------------------------------------------------*/
/* Fugro Status Information */

/** FugroStatus_1_0_t */
type FugroStatus_1_0_t struct {
	Header BlockHeader_t

	/* Time Header */
	TOW uint32
	WNc uint16

	Reserved          [2]uint8
	Status            uint32
	SubStartingTime   int32
	SubExpirationTime int32
	SubHourGlass      int32
	SubscribedMode    uint32
	SubCurrentMode    uint32
	SubLinkVector     uint32
	CRCGoodCount      uint32
	CRCBadCount       uint32
}

type FugroStatus_1_t FugroStatus_1_0_t

/*==External Sensor Blocks===================================================*/

/*--ExtSensorMeas_1_0_t : ---------------------------------------------------*/
/* Measurement set of external sensors of one epoch */

/** ExtSensorMeasAcceleration_1_0_t */
type ExtSensorMeasAcceleration_1_0_t struct {
	AccelerationX SBFDOUBLE
	AccelerationY SBFDOUBLE
	AccelerationZ SBFDOUBLE
}

/** ExtSensorMeasAngularRate_1_0_t */
type ExtSensorMeasAngularRate_1_0_t struct {
	AngularRateX SBFDOUBLE
	AngularRateY SBFDOUBLE
	AngularRateZ SBFDOUBLE
}

/** ExtSensorMeasInfo_1_0_t */
type ExtSensorMeasInfo_1_0_t struct {
	SensorTemperature int16
	Reserved          [22]uint8
}

/** ExtSensorMeasVelocity_1_0_t */
type ExtSensorMeasVelocity_1_0_t struct {
	VelocityX float32
	VelocityY float32
	VelocityZ float32
	StdDevX   float32
	StdDevY   float32
	StdDevZ   float32
}

/** ExtSensorMeasZeroVelocityFlag_1_0_t */
type ExtSensorMeasZeroVelocityFlag_1_0_t struct {
	ZeroVelocityFlag SBFDOUBLE
	Reserved         [16]uint8
}

/** ExtSensorMeasData_1_0_t */

//
// TODO - Figure out how to model C++ union in go. Determine if we even need this
//
//  type union
//  {
// 	 ExtSensorMeasAcceleration_1_0_t Acceleration;
// 	 ExtSensorMeasAngularRate_1_0_t AngularRate;
// 	 ExtSensorMeasInfo_1_0_t Info;
// 	 ExtSensorMeasVelocity_1_0_t Velocity;
// 	 ExtSensorMeasZeroVelocityFlag_1_0_t ZeroVelocityFlag;
//  } ExtSensorMeasData_1_0_t;

type ExtSensorMeasData_1_0_t struct {
	Acceleration     ExtSensorMeasAcceleration_1_0_t
	AngularRate      ExtSensorMeasAngularRate_1_0_t
	Info             ExtSensorMeasInfo_1_0_t
	Velocity         ExtSensorMeasVelocity_1_0_t
	ZeroVelocityFlag ExtSensorMeasZeroVelocityFlag_1_0_t
}

/** ExtSensorMeasSet_1_0_t */
type ExtSensorMeasSet_1_0_t struct {
	Source            uint8
	SensorModel       uint8
	Type              uint8
	ObsInfo           uint8
	ExtSensorMeasData ExtSensorMeasData_1_0_t
}

/** ExtSensorMeas_1_0_t */
type ExtSensorMeas_1_0_t struct {
	Header BlockHeader_t

	/* Time Header */
	TOW uint32
	WNc uint16

	N             uint8
	SBSize        uint8
	ExtSensorMeas [SBF_EXTSENSORMEAS_1_0_EXTSENSORMEAS_LENGTH]ExtSensorMeasSet_1_0_t
}

type ExtSensorMeasAcceleration_1_t ExtSensorMeasAcceleration_1_0_t
type ExtSensorMeasAngularRate_1_t ExtSensorMeasAngularRate_1_0_t
type ExtSensorMeasInfo_1_t ExtSensorMeasInfo_1_0_t
type ExtSensorMeasVelocity_1_t ExtSensorMeasVelocity_1_0_t
type ExtSensorMeasZeroVelocityFlag_1_t ExtSensorMeasZeroVelocityFlag_1_0_t
type ExtSensorMeasData_1_t ExtSensorMeasData_1_0_t
type ExtSensorMeasSet_1_t ExtSensorMeasSet_1_0_t
type ExtSensorMeas_1_t ExtSensorMeas_1_0_t

type ExtSensorMeasSB_1_0_t ExtSensorMeasSet_1_0_t
type ExtSensorMeasSB_1_t ExtSensorMeasSet_1_t

/*--ExtSensorStatus_1_0_t : -------------------------------------------------*/
/* Overall status of external sensors */

/** ExtSensorStatus_1_0_t */
type ExtSensorStatus_1_0_t struct {
	Header BlockHeader_t

	/* Time Header */
	TOW uint32
	WNc uint16

	Source      uint8
	SensorModel uint8
	StatusType  uint8
	Reserved    [3]uint8
	StatusBits  [SBF_EXTSENSORSTATUS_1_0_STATUSBITS_LENGTH]uint8
}
type ExtSensorStatus_1_t ExtSensorStatus_1_0_t

/*--ExtSensorSetup_1_0_t : --------------------------------------------------*/
/* General information about the setup of external sensors */

/** OneSensor_1_0_t */
type OneSensor_1_0_t struct {
	Source      uint8
	SensorModel uint8
	MeasTypes   uint16
}

/** ExtSensorSetup_1_0_t */
type ExtSensorSetup_1_0_t struct {
	Header BlockHeader_t

	/* Time Header */
	TOW uint32
	WNc uint16

	N              uint8
	SBLength       uint8
	ExtSensorSetup [SBF_EXTSENSORSETUP_1_0_EXTSENSORSETUP_LENGTH]OneSensor_1_0_t
}

/*--ExtSensorSetup_1_1_t : --------------------------------------------------*/
/* General information about the setup of external sensors */
/** OneSensor_1_1_t */
type OneSensor_1_1_t struct {
	Source         uint8
	SensorModel    uint8
	MeasTypes      uint16
	LeverArmSource uint8
	Reserved1      uint8

	_padding [SBF_ONESENSOR_1_1__PADDING_LENGTH]uint8
}

/** ExtSensorSetup_1_1_t */
type ExtSensorSetup_1_1_t struct {
	Header BlockHeader_t

	/* Time Header */
	TOW uint32
	WNc uint16

	N              uint8
	SBLength       uint8
	ExtSensorSetup [SBF_EXTSENSORSETUP_1_1_EXTSENSORSETUP_LENGTH]OneSensor_1_1_t
}

/*--ExtSensorSetup_1_2_t : --------------------------------------------------*/
/* General information about the setup of external sensors */

/** OneSensor_1_2_t */
type OneSensor_1_2_t struct {
	Source         uint8
	SensorModel    uint8
	MeasTypes      uint16
	LeverArmSource uint8
	Reserved1      uint8
	LeverArmX      int16
	LeverArmY      int16
	LeverArmZ      int16
	LeverArmStdX   uint16
	LeverArmStdY   uint16
	LeverArmStdZ   uint16
	Reserved2      uint16
}

/** ExtSensorSetup_1_2_t */
type ExtSensorSetup_1_2_t struct {
	Header BlockHeader_t

	/* Time Header */
	TOW uint32
	WNc uint16

	N              uint8
	SBLength       uint8
	ExtSensorSetup [SBF_EXTSENSORSETUP_1_2_EXTSENSORSETUP_LENGTH]OneSensor_1_2_t
}

type OneSensor_1_t OneSensor_1_2_t
type ExtSensorSetup_1_t ExtSensorSetup_1_2_t

/*--ExtSensorStatus_2_0_t : -------------------------------------------------*/
/* Overall status of external sensors */

/** ExtSensorStatus_2_0_t */
type ExtSensorStatus_2_0_t struct {
	Header BlockHeader_t

	/* Time Header */
	TOW uint32
	WNc uint16

	Source      uint8
	SensorModel uint8
	Data        [SBF_EXTSENSORSTATUS_2_0_DATA_LENGTH]uint8
}
type ExtSensorStatus_2_t ExtSensorStatus_2_0_t

/*--ExtSensorInfo_1_0_t : ---------------------------------------------------*/
/* Configuration information of external sensors */

/** ExtSensorInfo_1_0_t */
type ExtSensorInfo_1_0_t struct {
	Header BlockHeader_t

	/* Time Header */
	TOW uint32
	WNc uint16

	Source      uint8
	SensorModel uint8
	Data        [SBF_EXTSENSORINFO_1_0_DATA_LENGTH]uint8
}
type ExtSensorInfo_1_t ExtSensorInfo_1_0_t

/*--IMUSetup_1_0_t : --------------------------------------------------------*/
/* General information about the setup of the IMU */

/** IMUSetup_1_0_t */
type IMUSetup_1_0_t struct {
	Header BlockHeader_t

	/* Time Header */
	TOW uint32
	WNc uint16

	Reserved     uint8
	SerialPort   uint8
	AntLeverArmX float32
	AntLeverArmY float32
	AntLeverArmZ float32
	ThetaX       float32
	ThetaY       float32
	ThetaZ       float32
}

type IMUSetup_1_t IMUSetup_1_0_t

/*==Status Blocks============================================================*/

/*--ReceiverStatus_1_0_t : --------------------------------------------------*/
/* Overall status information of the receiver */

/** AGCData_1_0_t */
type AGCData_1_0_t struct {
	L1Magn             uint8
	L2Magn             uint8
	AGCSettings        uint8
	AttenuatorSettings uint8
}

/** ReceiverStatus_1_0_t */
type ReceiverStatus_1_0_t struct {
	Header BlockHeader_t

	/* Time Header */
	TOW uint32
	WNc uint16

	CPULoad  uint8
	Reserved uint8
	UpTime   uint32
	RxStatus uint32
	AGCData  AGCData_1_0_t
}

type AGCData_1_t AGCData_1_0_t
type ReceiverStatus_1_t ReceiverStatus_1_0_t

/*--TrackingStatus_1_0_t : --------------------------------------------------*/
/* Status of the tracking for all receiver channels */

/** TrackingStatusChannel_1_0_t */
type TrackingStatusChannel_1_0_t struct {
	RxChannel      uint8
	SVID           uint8
	AttitudeStatus uint8
	Status         uint8
	Azimuth        int16
	Elevation      int8
	Health         uint8
	ElevChange     int8
	_padding       [SBF_TRACKINGSTATUSCHANNEL_1_0__PADDING_LENGTH]uint8
}

/** TrackingStatus_1_0_t */
type TrackingStatus_1_0_t struct {
	Header BlockHeader_t

	/* Time Header */
	TOW uint32
	WNc uint16

	N           uint8
	SBSize      uint8
	ChannelData [SBF_TRACKINGSTATUS_1_0_CHANNELDATA_LENGTH]TrackingStatusChannel_1_0_t
}

type TrackingStatusChannel_1_t TrackingStatusChannel_1_0_t
type TrackingStatus_1_t TrackingStatus_1_0_t

/*--ChannelStatus_1_0_t : ---------------------------------------------------*/
/* Status of the tracking for all receiver channels */

/** ChannelStateInfo_1_0_t */
type ChannelStateInfo_1_0_t struct {
	Antenna        uint8
	ReservedA      uint8
	TrackingStatus uint16
	PVTStatus      uint16
	PVTInfo        uint16
}

/** ChannelSatInfo_1_0_t */
type ChannelSatInfo_1_0_t struct {
	SVID         uint8
	FreqNr       uint8
	Reserved1    [2]uint8
	Az_RiseSet   uint16
	HealthStatus uint16
	Elev         int8
	N2           uint8
	Channel      uint8
	Reserved2    uint8
}

/** ChannelStatus_1_0_t */
type ChannelStatus_1_0_t struct {
	Header BlockHeader_t

	/* Time Header */
	TOW uint32
	WNc uint16

	N        uint8
	SB1Size  uint8
	SB2Size  uint8
	Reserved [3]uint8
	Data     [SBF_CHANNELSTATUS_1_0_DATA_LENGTH]uint8
}

type ChannelStateInfo_1_t ChannelStateInfo_1_0_t
type ChannelSatInfo_1_t ChannelSatInfo_1_0_t
type ChannelStatus_1_t ChannelStatus_1_0_t

/*--ReceiverStatus_2_0_t : --------------------------------------------------*/
/* Overall status information of the receiver */

/** AGCState_2_0_t */
type AGCState_2_0_t struct {
	FrontendID   uint8
	Gain         int8
	SampleVar    uint8
	BlankingStat uint8
}

/** ReceiverStatus_2_0_t */
type ReceiverStatus_2_0_t struct {
	Header BlockHeader_t

	/* Time Header */
	TOW uint32
	WNc uint16

	CPULoad     uint8
	ExtError    uint8
	UpTime      uint32
	RxStatus    uint32
	RxError     uint32
	N           uint8
	SBSize      uint8
	CmdCount    uint8
	Temperature uint8
	AGCState    [SBF_RECEIVERSTATUS_2_0_AGCSTATE_LENGTH]AGCState_2_0_t
}

/*--ReceiverStatus_2_1_t : --------------------------------------------------*/
/* Overall status information of the receiver */

/** AGCState_2_1_t */
type AGCState_2_1_t struct {
	FrontendID   uint8
	Gain         int8
	SampleVar    uint8
	BlankingStat uint8
}

/** ReceiverStatus_2_1_t */
type ReceiverStatus_2_1_t struct {
	Header BlockHeader_t

	/* Time Header */
	TOW uint32
	WNc uint16

	CPULoad     uint8
	ExtError    uint8
	UpTime      uint32
	RxStatus    uint32
	RxError     uint32
	N           uint8
	SBSize      uint8
	CmdCount    uint8
	Temperature uint8
	AGCState    [SBF_RECEIVERSTATUS_2_1_AGCSTATE_LENGTH]AGCState_2_1_t
}

type AGCState_2_t AGCState_2_1_t
type ReceiverStatus_2_t ReceiverStatus_2_1_t

/*--SatVisibility_1_0_t : ---------------------------------------------------*/
/* Azimuth/elevation of visible satellites */

/** SatInfo_1_0_t */
type SatInfo_1_0_t struct {
	SVID          uint8
	FreqNr        uint8
	Azimuth       uint16
	Elevation     int16
	RiseSet       uint8
	SatelliteInfo uint8
}

/** SatVisibility_1_0_t */
type SatVisibility_1_0_t struct {
	Header BlockHeader_t

	/* Time Header */
	TOW uint32
	WNc uint16

	N        uint8
	SBLength uint8
	SatInfo  [SBF_SATVISIBILITY_1_0_SATINFO_LENGTH]SatInfo_1_0_t
}
type SatInfo_1_t SatInfo_1_0_t
type SatVisibility_1_t SatVisibility_1_0_t

/*--InputLink_1_0_t : -------------------------------------------------------*/
/* Statistics on input streams */

/** InputStatsSub_1_0_t */
type InputStatsSub_1_0_t struct {
	CD               uint8
	Type             uint8
	AgeOfLastMessage uint16
	NrBytesReceived  uint32
	NrBytesAccepted  uint32
	NrMsgReceived    uint32
	NrMsgAccepted    uint32
}

/** InputLink_1_0_t */
type InputLink_1_0_t struct {
	Header BlockHeader_t

	/* Time Header */
	TOW uint32
	WNc uint16

	N        uint8
	SBLength uint8
	Data     [SBF_INPUTLINK_1_0_DATA_LENGTH]uint8
}

type InputStatsSub_1_t InputStatsSub_1_0_t
type InputLink_1_t InputLink_1_0_t

/*--OutputLink_1_0_t : ------------------------------------------------------*/
/* Statistics on output streams */

/** OutputTypeSub_1_0_t */
type OutputTypeSub_1_0_t struct {
	Type       uint8
	Percentage uint8
	_padding   [SBF_OUTPUTTYPESUB_1_0__PADDING_LENGTH]uint8
}

/** OutputStatsSub_1_0_t */
type OutputStatsSub_1_0_t struct {
	CD              uint8
	N2              uint8
	AllowedRate     uint16
	NrBytesProduced uint32
	NrBytesSent     uint32
}

/** OutputLink_1_0_t */
type OutputLink_1_0_t struct {
	Header BlockHeader_t

	/* Time Header */
	TOW uint32
	WNc uint16

	N1        uint8
	SB1Length uint8
	SB2Length uint8
	Reserved  [3]uint8
	Data      [SBF_OUTPUTLINK_1_0_DATA_LENGTH]uint8
}

/*--OutputLink_1_1_t : ------------------------------------------------------*/
/* Statistics on output streams */

/** OutputTypeSub_1_1_t */
type OutputTypeSub_1_1_t struct {
	Type       uint8
	Percentage uint8
}

/** OutputStatsSub_1_1_t */
type OutputStatsSub_1_1_t struct {
	CD              uint8
	N2              uint8
	AllowedRate     uint16
	NrBytesProduced uint32
	NrBytesSent     uint32
	NrClients       uint8
	Reserved        [3]uint8
}

/** OutputLink_1_1_t */
type OutputLink_1_1_t struct {
	Header BlockHeader_t

	/* Time Header */
	TOW uint32
	WNc uint16

	N1        uint8
	SB1Length uint8
	SB2Length uint8
	Reserved  [3]uint8
	Data      [SBF_OUTPUTLINK_1_1_DATA_LENGTH]uint8
}

type OutputTypeSub_1_t OutputTypeSub_1_1_t
type OutputStatsSub_1_t OutputStatsSub_1_1_t
type OutputLink_1_t OutputLink_1_1_t

/*--NTRIPClientStatus_1_0_t : -----------------------------------------------*/
/* NTRIP client connection status */

/** NTRIPClientConnection_1_0_t */
type NTRIPClientConnection_1_0_t struct {
	CDIndex   uint8
	Status    uint8
	ErrorCode uint8
	Info      uint8
}

/** NTRIPClientStatus_1_0_t */
type NTRIPClientStatus_1_0_t struct {
	Header BlockHeader_t

	/* Time Header */
	TOW uint32
	WNc uint16

	N                     uint8
	SBLength              uint8
	NTRIPClientConnection [SBF_NTRIPCLIENTSTATUS_1_0_NTRIPCLIENTCONNECTION_LENGTH]NTRIPClientConnection_1_0_t
}

type NTRIPClientConnection_1_t NTRIPClientConnection_1_0_t
type NTRIPClientStatus_1_t NTRIPClientStatus_1_0_t

/*--NTRIPServerStatus_1_0_t : -----------------------------------------------*/
/* NTRIP server connection status */

/** NTRIPServerConnection_1_0_t */
type NTRIPServerConnection_1_0_t struct {
	CDIndex   uint8
	Status    uint8
	ErrorCode uint8
	Info      uint8
}

/** NTRIPServerStatus_1_0_t */
type NTRIPServerStatus_1_0_t struct {
	Header BlockHeader_t

	/* Time Header */
	TOW uint32
	WNc uint16

	N                     uint8
	SBLength              uint8
	NTRIPServerConnection [SBF_NTRIPSERVERSTATUS_1_0_NTRIPSERVERCONNECTION_LENGTH]NTRIPServerConnection_1_0_t
}

type NTRIPServerConnection_1_t NTRIPServerConnection_1_0_t
type NTRIPServerStatus_1_t NTRIPServerStatus_1_0_t

/*--IPStatus_1_0_t : --------------------------------------------------------*/
/* IP address, gateway and MAC address of Ethernet interface */

/** IPStatus_1_0_t */
type IPStatus_1_0_t struct {
	Header BlockHeader_t

	/* Time Header */
	TOW uint32
	WNc uint16

	MACAddress [SBF_IPSTATUS_1_0_MACADDRESS_LENGTH]uint8
	IPAddress  [SBF_IPSTATUS_1_0_IPADDRESS_LENGTH]uint8
	Gateway    [SBF_IPSTATUS_1_0_GATEWAY_LENGTH]uint8
	Netmask    uint8
	Reserved   [3]uint8
}

/*--IPStatus_1_1_t : --------------------------------------------------------*/
/* IP address, gateway and MAC address of Ethernet interface */

/** IPStatus_1_1_t */
type IPStatus_1_1_t struct {
	Header BlockHeader_t

	/* Time Header */
	TOW uint32
	WNc uint16

	MACAddress [SBF_IPSTATUS_1_1_MACADDRESS_LENGTH]uint8
	IPAddress  [SBF_IPSTATUS_1_1_IPADDRESS_LENGTH]uint8
	Gateway    [SBF_IPSTATUS_1_1_GATEWAY_LENGTH]uint8
	Netmask    uint8
	Reserved   [3]uint8
	HostName   [SBF_IPSTATUS_1_1_HOSTNAME_LENGTH]byte
}
type IPStatus_1_t IPStatus_1_1_t

/*--WiFiAPStatus_1_0_t : ----------------------------------------------------*/
/* WiFi status in access point mode */

/** WiFiClient_1_0_t */
type WiFiClient_1_0_t struct {
	ClientHostName   [SBF_WIFICLIENT_1_0_CLIENTHOSTNAME_LENGTH]byte
	ClientMACAddress [SBF_WIFICLIENT_1_0_CLIENTMACADDRESS_LENGTH]uint8
	ClientIPAddress  [SBF_WIFICLIENT_1_0_CLIENTIPADDRESS_LENGTH]uint8
	_padding         [SBF_WIFICLIENT_1_0__PADDING_LENGTH]uint8
}

/** WiFiAPStatus_1_0_t */
type WiFiAPStatus_1_0_t struct {
	Header BlockHeader_t

	/* Time Header */
	TOW uint32
	WNc uint16

	N           uint8
	SBLength    uint8
	APIPAddress [SBF_WIFIAPSTATUS_1_0_APIPADDRESS_LENGTH]uint8
	Mode        uint8
	Hotspot     uint8
	Reserved    [2]uint8
	WiFiClient  [SBF_WIFIAPSTATUS_1_0_WIFICLIENT_LENGTH]WiFiClient_1_0_t
}
type WiFiClient_1_t WiFiClient_1_0_t
type WiFiAPStatus_1_t WiFiAPStatus_1_0_t

/*--WiFiClientStatus_1_0_t : ------------------------------------------------*/
/* WiFi status in client mode */

/** WiFiClientStatus_1_0_t */
type WiFiClientStatus_1_0_t struct {
	Header BlockHeader_t

	/* Time Header */
	TOW uint32
	WNc uint16

	SSID_AP   [SBF_WIFICLIENTSTATUS_1_0_SSID_AP_LENGTH]byte
	IPAddress [SBF_WIFICLIENTSTATUS_1_0_IPADDRESS_LENGTH]uint8
	Reserved  [1]uint8
	SigLevel  int8
	Status    uint8
	ErrorCode uint8
	_padding  [SBF_WIFICLIENTSTATUS_1_0__PADDING_LENGTH]uint8
}
type WiFiClientStatus_1_t WiFiClientStatus_1_0_t

/*--CellularStatus_1_0_t : --------------------------------------------------*/
/* Cellular status */

/** CellularStatus_1_0_t */
type CellularStatus_1_0_t struct {
	Header BlockHeader_t

	/* Time Header */
	TOW uint32
	WNc uint16

	ConnectionType uint8
	RSSI           int8
	OperatorName   [SBF_CELLULARSTATUS_1_0_OPERATORNAME_LENGTH]byte
	Status         uint8
	ErrorCode      uint8
	_padding       [SBF_CELLULARSTATUS_1_0__PADDING_LENGTH]uint8
}

/*--CellularStatus_1_1_t : --------------------------------------------------*/
/* Cellular status */

/** CellularStatus_1_1_t */
type CellularStatus_1_1_t struct {
	Header BlockHeader_t

	/* Time Header */
	TOW uint32
	WNc uint16

	ConnectionType uint8
	RSSI           int8
	OperatorName   [SBF_CELLULARSTATUS_1_1_OPERATORNAME_LENGTH]byte
	Status         uint8
	ErrorCode      uint8
	DataCall       uint8
	Info           uint8
}
type CellularStatus_1_t CellularStatus_1_1_t

/*--BluetoothStatus_1_0_t : -------------------------------------------------*/
/* Bluetooth status */

/** BTDevice_1_0_t */
type BTDevice_1_0_t struct {
	DeviceName [SBF_BTDEVICE_1_0_DEVICENAME_LENGTH]byte
	Flags      uint8
	_padding   [SBF_BTDEVICE_1_0__PADDING_LENGTH]uint8
}

/** BluetoothStatus_1_0_t */
type BluetoothStatus_1_0_t struct {
	Header BlockHeader_t

	/* Time Header */
	TOW uint32
	WNc uint16

	N        uint8
	SBLength uint8
	Mode     uint8
	Reserved [3]uint8
	BTDevice [SBF_BLUETOOTHSTATUS_1_0_BTDEVICE_LENGTH]BTDevice_1_0_t
}

type BTDevice_1_t BTDevice_1_0_t
type BluetoothStatus_1_t BluetoothStatus_1_0_t

/*--DynDNSStatus_1_0_t : ----------------------------------------------------*/
/* DynDNS status */

/** DynDNSStatus_1_0_t */
type DynDNSStatus_1_0_t struct {
	Header BlockHeader_t

	/* Time Header */
	TOW uint32
	WNc uint16

	Status    uint8
	ErrorCode uint8
}

/*--DynDNSStatus_1_1_t : ----------------------------------------------------*/
/* DynDNS status */

/** DynDNSStatus_1_1_t */
type DynDNSStatus_1_1_t struct {
	Header BlockHeader_t

	/* Time Header */
	TOW uint32
	WNc uint16

	Status    uint8
	ErrorCode uint8
	IPAddress [SBF_DYNDNSSTATUS_1_1_IPADDRESS_LENGTH]uint8
}
type DynDNSStatus_1_t DynDNSStatus_1_1_t

/*--BatteryStatus_1_0_t : ---------------------------------------------------*/
/* Battery status */

/** Battery_1_0_t */
type Battery_1_0_t struct {
	ChargeLevel   uint8
	Status        uint8
	RemainingTime uint16
}

/** BatteryStatus_1_0_t */
type BatteryStatus_1_0_t struct {
	Header BlockHeader_t

	/* Time Header */
	TOW uint32
	WNc uint16

	N         uint8
	SBLength  uint8
	ExtSupply uint8
	Reserved  [3]uint8
	Battery   [SBF_BATTERYSTATUS_1_0_BATTERY_LENGTH]Battery_1_0_t
}

/*--BatteryStatus_1_1_t : ---------------------------------------------------*/
/* Battery status */

/** Battery_1_1_t */
type Battery_1_1_t struct {
	ChargeLevel   uint8
	Status        uint8
	RemainingTime uint16
	Voltage       uint16
	Current       int16
}

/** BatteryStatus_1_1_t */
type BatteryStatus_1_1_t struct {
	Header BlockHeader_t

	/* Time Header */
	TOW uint32
	WNc uint16

	N         uint8
	SBLength  uint8
	ExtSupply uint8
	Reserved  [3]uint8
	Battery   [SBF_BATTERYSTATUS_1_1_BATTERY_LENGTH]Battery_1_1_t
}

/*--BatteryStatus_1_2_t : ---------------------------------------------------*/
/* Battery status */

/** Battery_1_2_t */
type Battery_1_2_t struct {
	ChargeLevel   uint8
	Status        uint8
	RemainingTime uint16
	Voltage       uint16
	Current       int16
	Temperature   int8
	Reserved      [3]uint8
}

/** BatteryStatus_1_2_t */
type BatteryStatus_1_2_t struct {
	Header BlockHeader_t

	/* Time Header */
	TOW uint32
	WNc uint16

	N         uint8
	SBLength  uint8
	ExtSupply uint8
	Reserved  [3]uint8
	Battery   [SBF_BATTERYSTATUS_1_2_BATTERY_LENGTH]Battery_1_2_t
}

type Battery_1_t Battery_1_2_t
type BatteryStatus_1_t BatteryStatus_1_2_t

/*--PowerStatus_1_0_t : -----------------------------------------------------*/
/* Power supply source and voltage */

/** PowerStatus_1_0_t */
type PowerStatus_1_0_t struct {
	Header BlockHeader_t

	/* Time Header */
	TOW uint32
	WNc uint16

	PowerInfo uint16
}

type PowerStatus_1_t PowerStatus_1_0_t

/*--QualityInd_1_0_t : ------------------------------------------------------*/
/* Quality indicators */

/** QualityInd_1_0_t */
type QualityInd_1_0_t struct {
	Header BlockHeader_t

	/* Time Header */
	TOW uint32
	WNc uint16

	N          uint8
	Reserved   uint8
	Indicators [SBF_QUALITYIND_1_0_INDICATORS_LENGTH]uint16
}

type QualityInd_1_t QualityInd_1_0_t

/*--DiskStatus_1_0_t : ------------------------------------------------------*/
/* Internal logging status */

/** DiskData_1_0_t */
type DiskData_1_0_t struct {
	DiskID            uint8
	Status            uint8
	DiskUsageMSB      uint16
	DiskUsageLSB      uint32
	DiskSize          uint32
	CreateDeleteCount uint8

	_padding [SBF_DISKDATA_1_0__PADDING_LENGTH]uint8
}

/** DiskStatus_1_0_t */
type DiskStatus_1_0_t struct {
	Header BlockHeader_t

	/* Time Header */
	TOW uint32
	WNc uint16

	N        uint8
	SBLength uint8
	Reserved [4]uint8
	DiskData [SBF_DISKSTATUS_1_0_DISKDATA_LENGTH]DiskData_1_0_t
}

/** DiskData_1_1_t */
type DiskData_1_1_t struct {
	DiskID            uint8
	Status            uint8
	DiskUsageMSB      uint16
	DiskUsageLSB      uint32
	DiskSize          uint32
	CreateDeleteCount uint8
	Error             uint8

	_padding [SBF_DISKDATA_1_1__PADDING_LENGTH]uint8
}

/** DiskStatus_1_1_t */
type DiskStatus_1_1_t struct {
	Header BlockHeader_t

	/* Time Header */
	TOW uint32
	WNc uint16

	N        uint8
	SBLength uint8
	Reserved [4]uint8
	DiskData [SBF_DISKSTATUS_1_1_DISKDATA_LENGTH]DiskData_1_1_t
}

type DiskData_1_t DiskData_1_1_t
type DiskStatus_1_t DiskStatus_1_1_t

/*--LogStatus_1_0_t : -------------------------------------------------------*/
/* Log sessions status */

/** FileUploadStatus_1_0_t */
type FileUploadStatus_1_0_t struct {
	Type              uint8
	ErrorCode         uint8
	RetryQueueSize    uint8
	NrFailedTransfers uint8
}

/** LogSession_1_0_t */
type LogSession_1_0_t struct {
	SessionID        uint8
	SessionStatus    uint8
	N2               uint8
	Reserved         uint8
	FileUploadStatus [SBF_LOGSESSION_1_0_FILEUPLOADSTATUS_LENGTH]FileUploadStatus_1_0_t
}

/** LogStatus_1_0_t */
type LogStatus_1_0_t struct {
	Header BlockHeader_t

	/* Time Header */
	TOW uint32
	WNc uint16

	N1          uint8
	SB1Length   uint8
	SB2Length   uint8
	Reserved    [3]uint8
	LogSessions [SBF_LOGSTATUS_1_0_LOGSESSIONS_LENGTH]LogSession_1_0_t
}

type FileUploadStatus_1_t FileUploadStatus_1_0_t
type LogSession_1_t LogSession_1_0_t
type LogStatus_1_t LogStatus_1_0_t

/*--UHFStatus_1_0_t : -------------------------------------------------------*/
/* UHF status */

/** UHFData_1_0_t */
type UHFData_1_0_t struct {
	Frequency uint32
	Channel   uint8
	Bandwidth uint8
	RSSI      int8

	_padding [SBF_UHFDATA_1_0__PADDING_LENGTH]uint8
}

/** UHFStatus_1_0_t */
type UHFStatus_1_0_t struct {
	Header BlockHeader_t

	/* Time Header */
	TOW uint32
	WNc uint16

	N         uint8
	SBLength  uint8
	UHFID     uint8
	Status    uint8
	ErrorCode uint8
	Reserved  uint8
	UHFData   [SBF_UHFSTATUS_1_0_UHFDATA_LENGTH]UHFData_1_0_t
}

type UHFData_1_t UHFData_1_0_t
type UHFStatus_1_t UHFStatus_1_0_t

/*--RFStatus_1_0_t : --------------------------------------------------------*/
/* Radio-frequency interference mitigation status */

/** RFBand_1_0_t */
type RFBand_1_0_t struct {
	Frequency uint32
	Bandwidth uint16
	Info      uint8

	_padding [SBF_RFBAND_1_0__PADDING_LENGTH]uint8
}

/** RFStatus_1_0_t */
type RFStatus_1_0_t struct {
	Header BlockHeader_t

	/* Time Header */
	TOW uint32
	WNc uint16

	N        uint8
	SBLength uint8
	Flags    uint8
	Reserved [3]uint8
	RFBand   [SBF_RFSTATUS_1_0_RFBAND_LENGTH]RFBand_1_0_t
}

type RFBand_1_t RFBand_1_0_t
type RFStatus_1_t RFStatus_1_0_t

/*--RIMSHealth_1_0_t : ------------------------------------------------------*/
/* Health status of the receiver */

/** RIMSHealth_1_0_t */
type RIMSHealth_1_0_t struct {
	Header BlockHeader_t

	/* Time Header */
	TOW uint32
	WNc uint16

	Reserved     [2]uint8
	RxError      uint32
	Degraded     uint32
	RxInVoltage  uint16
	RxInCurrent  uint16
	AntInVoltage uint16
	AntInCurrent uint16
	Temperature  int16
	MemoryLoad   uint8

	_padding [SBF_RIMSHEALTH_1_0__PADDING_LENGTH]uint8
}

type RIMSHealth_1_t RIMSHealth_1_0_t

/*--OSNMAStatus_1_0_t : -----------------------------------------------------*/
/* OSNMA status information */

/** MACKStatus_1_0_t */
type MACKStatus_1_0_t struct {
	SVID               uint8
	MACInfo            uint8
	VerificationStatus uint8
	MacPosition        uint8
}

/** OSNMAStatus_1_0_t */
type OSNMAStatus_1_0_t struct {
	Header BlockHeader_t

	/* Time Header */
	TOW uint32
	WNc uint16

	N                    uint8
	ValidData            uint8
	Chain1_Status        uint16
	Chain2_Status        uint16
	Chain3_Status        uint16
	Chain4_Status        uint16
	PKR_Status           uint16
	NMAStatus            uint8
	ChainID              uint8
	CPKS                 uint8
	Reserved             uint8
	KRootInfo            uint16
	KeySize              uint16
	MacSize              uint8
	MacLookupTable       uint8
	SVID                 uint8
	KRootTOWH            uint8
	KRootWN              uint16
	EmergencyMsgReceived uint8
	Reserved1            [SBF_OSNMASTATUS_1_0_RESERVED1_LENGTH]uint8
	MACKStatus           [SBF_OSNMASTATUS_1_0_MACKSTATUS_LENGTH]MACKStatus_1_0_t
}

type MACKStatus_1_t MACKStatus_1_0_t
type OSNMAStatus_1_t OSNMAStatus_1_0_t

/*--GALNavMonitor_1_0_t : ---------------------------------------------------*/
/* Monitoring navigation data per Galileo satellite. */

/** GALNavMonitor_1_0_t */
type GALNavMonitor_1_0_t struct {
	Header BlockHeader_t

	/* Time Header */
	TOW uint32
	WNc uint16

	SVID         uint8
	Source       uint8
	PositionDiff float32
	TimeCorrDiff float32
	IODNavCurr   uint16
	IODNavPrev   uint16
	Flags        uint8

	_padding [SBF_GALNAVMONITOR_1_0__PADDING_LENGTH]uint8
}
type GALNavMonitor_1_t GALNavMonitor_1_0_t

/*--INAVmonitor_1_0_t : -----------------------------------------------------*/
/* Reed-Solomon and SSP status information */

/** INAVmonitor_1_0_t */
type INAVmonitor_1_0_t struct {
	Header BlockHeader_t

	/* Time Header */
	TOW uint32
	WNc uint16

	SVID      uint8
	RS_status uint8
	SSPstatus uint8

	_padding [SBF_INAVMONITOR_1_0__PADDING_LENGTH]uint8
}
type INAVmonitor_1_t INAVmonitor_1_0_t

/*--P2PPStatus_1_0_t : ------------------------------------------------------*/
/* P2PP client/server status */

/** P2PPSession_1_0_t */
type P2PPSession_1_0_t struct {
	SessionID uint8
	Port      uint8
	Status    uint8
	ErrorCode uint8
}

/** P2PPStatus_1_0_t */
type P2PPStatus_1_0_t struct {
	Header BlockHeader_t

	/* Time Header */
	TOW uint32
	WNc uint16

	N           uint8
	SBLength    uint8
	P2PPSession [SBF_P2PPSTATUS_1_0_P2PPSESSION_LENGTH]P2PPSession_1_0_t
}

type P2PPSession_1_t P2PPSession_1_0_t
type P2PPStatus_1_t P2PPStatus_1_0_t

/*--AuthenticationStatus_1_0_t : --------------------------------------------*/

/** AuthenticationStatus_1_0_t */
type AuthenticationStatus_1_0_t struct {
	Header BlockHeader_t

	/* Time Header */
	TOW uint32
	WNc uint16

	AdkdMask      uint16
	window        uint32
	M             uint32
	GalActiveMask uint64
	GalEphMask    uint64
	GalAlmMask    uint64
	GalIonMask    uint64
	GalUtcMask    uint64
	GpsActiveMask uint64
	GpsEphMask    uint64
	GpsIonMask    uint64
}

type AuthenticationStatus_1_t AuthenticationStatus_1_0_t

/*--CosmosStatus_1_0_t : ----------------------------------------------------*/
/* Cosmos receiver service status */

/** CosmosStatus_1_0_t */
type CosmosStatus_1_0_t struct {
	Header BlockHeader_t

	/* Time Header */
	TOW uint32
	WNc uint16

	Status   uint8
	_padding [SBF_COSMOSSTATUS_1_0__PADDING_LENGTH]uint8
}

type CosmosStatus_1_t CosmosStatus_1_0_t

/*==Miscellaneous Blocks=====================================================*/

/*--ReceiverSetup_1_0_t : ---------------------------------------------------*/
/* General information about the receiver installation */

/** ReceiverSetup_1_0_t */
type ReceiverSetup_1_0_t struct {
	Header BlockHeader_t

	/* Time Header */
	TOW uint32
	WNc uint16

	Reserved     [2]uint8
	MarkerName   [SBF_RECEIVERSETUP_1_0_MARKERNAME_LENGTH]byte
	MarkerNumber [SBF_RECEIVERSETUP_1_0_MARKERNUMBER_LENGTH]byte
	Observer     [SBF_RECEIVERSETUP_1_0_OBSERVER_LENGTH]byte
	Agency       [SBF_RECEIVERSETUP_1_0_AGENCY_LENGTH]byte
	RxSerialNbr  [SBF_RECEIVERSETUP_1_0_RXSERIALNBR_LENGTH]byte
	RxName       [SBF_RECEIVERSETUP_1_0_RXNAME_LENGTH]byte
	RxVersion    [SBF_RECEIVERSETUP_1_0_RXVERSION_LENGTH]byte
	AntSerialNbr [SBF_RECEIVERSETUP_1_0_ANTSERIALNBR_LENGTH]byte
	AntType      [SBF_RECEIVERSETUP_1_0_ANTTYPE_LENGTH]byte
	deltaH       float32 /* [m] */
	deltaE       float32 /* [m] */
	deltaN       float32 /* [m] */
}

/*--ReceiverSetup_1_1_t : ---------------------------------------------------*/
/* General information about the receiver installation */

/** ReceiverSetup_1_1_t */
type ReceiverSetup_1_1_t struct {
	Header BlockHeader_t

	/* Time Header */
	TOW uint32
	WNc uint16

	Reserved     [2]uint8
	MarkerName   [SBF_RECEIVERSETUP_1_1_MARKERNAME_LENGTH]byte
	MarkerNumber [SBF_RECEIVERSETUP_1_1_MARKERNUMBER_LENGTH]byte
	Observer     [SBF_RECEIVERSETUP_1_1_OBSERVER_LENGTH]byte
	Agency       [SBF_RECEIVERSETUP_1_1_AGENCY_LENGTH]byte
	RxSerialNbr  [SBF_RECEIVERSETUP_1_1_RXSERIALNBR_LENGTH]byte
	RxName       [SBF_RECEIVERSETUP_1_1_RXNAME_LENGTH]byte
	RxVersion    [SBF_RECEIVERSETUP_1_1_RXVERSION_LENGTH]byte
	AntSerialNbr [SBF_RECEIVERSETUP_1_1_ANTSERIALNBR_LENGTH]byte
	AntType      [SBF_RECEIVERSETUP_1_1_ANTTYPE_LENGTH]byte
	deltaH       float32 /* [m] */
	deltaE       float32 /* [m] */
	deltaN       float32 /* [m] */
	MarkerType   [SBF_RECEIVERSETUP_1_1_MARKERTYPE_LENGTH]byte
}

/*--ReceiverSetup_1_2_t : ---------------------------------------------------*/
/* General information about the receiver installation */

/** ReceiverSetup_1_2_t */
type ReceiverSetup_1_2_t struct {
	Header BlockHeader_t

	/* Time Header */
	TOW uint32
	WNc uint16

	Reserved      [2]uint8
	MarkerName    [SBF_RECEIVERSETUP_1_2_MARKERNAME_LENGTH]byte
	MarkerNumber  [SBF_RECEIVERSETUP_1_2_MARKERNUMBER_LENGTH]byte
	Observer      [SBF_RECEIVERSETUP_1_2_OBSERVER_LENGTH]byte
	Agency        [SBF_RECEIVERSETUP_1_2_AGENCY_LENGTH]byte
	RxSerialNbr   [SBF_RECEIVERSETUP_1_2_RXSERIALNBR_LENGTH]byte
	RxName        [SBF_RECEIVERSETUP_1_2_RXNAME_LENGTH]byte
	RxVersion     [SBF_RECEIVERSETUP_1_2_RXVERSION_LENGTH]byte
	AntSerialNbr  [SBF_RECEIVERSETUP_1_2_ANTSERIALNBR_LENGTH]byte
	AntType       [SBF_RECEIVERSETUP_1_2_ANTTYPE_LENGTH]byte
	deltaH        float32 /* [m] */
	deltaE        float32 /* [m] */
	deltaN        float32 /* [m] */
	MarkerType    [SBF_RECEIVERSETUP_1_2_MARKERTYPE_LENGTH]byte
	GNSSFWVersion [SBF_RECEIVERSETUP_1_2_GNSSFWVERSION_LENGTH]byte
}

/*--ReceiverSetup_1_3_t : ---------------------------------------------------*/
/* General information about the receiver installation */

/** ReceiverSetup_1_3_t */
type ReceiverSetup_1_3_t struct {
	Header BlockHeader_t

	/* Time Header */
	TOW uint32
	WNc uint16

	Reserved      [2]uint8
	MarkerName    [SBF_RECEIVERSETUP_1_3_MARKERNAME_LENGTH]byte
	MarkerNumber  [SBF_RECEIVERSETUP_1_3_MARKERNUMBER_LENGTH]byte
	Observer      [SBF_RECEIVERSETUP_1_3_OBSERVER_LENGTH]byte
	Agency        [SBF_RECEIVERSETUP_1_3_AGENCY_LENGTH]byte
	RxSerialNbr   [SBF_RECEIVERSETUP_1_3_RXSERIALNBR_LENGTH]byte
	RxName        [SBF_RECEIVERSETUP_1_3_RXNAME_LENGTH]byte
	RxVersion     [SBF_RECEIVERSETUP_1_3_RXVERSION_LENGTH]byte
	AntSerialNbr  [SBF_RECEIVERSETUP_1_3_ANTSERIALNBR_LENGTH]byte
	AntType       [SBF_RECEIVERSETUP_1_3_ANTTYPE_LENGTH]byte
	deltaH        float32 /* [m] */
	deltaE        float32 /* [m] */
	deltaN        float32 /* [m] */
	MarkerType    [SBF_RECEIVERSETUP_1_3_MARKERTYPE_LENGTH]byte
	GNSSFWVersion [SBF_RECEIVERSETUP_1_3_GNSSFWVERSION_LENGTH]byte
	ProductName   [SBF_RECEIVERSETUP_1_3_PRODUCTNAME_LENGTH]byte
}

/*--ReceiverSetup_1_4_t : ---------------------------------------------------*/
/* General information about the receiver installation */

/** ReceiverSetup_1_4_t */
type ReceiverSetup_1_4_t struct {
	Header BlockHeader_t

	/* Time Header */
	TOW uint32
	WNc uint16

	Reserved      [2]uint8
	MarkerName    [SBF_RECEIVERSETUP_1_4_MARKERNAME_LENGTH]byte
	MarkerNumber  [SBF_RECEIVERSETUP_1_4_MARKERNUMBER_LENGTH]byte
	Observer      [SBF_RECEIVERSETUP_1_4_OBSERVER_LENGTH]byte
	Agency        [SBF_RECEIVERSETUP_1_4_AGENCY_LENGTH]byte
	RxSerialNbr   [SBF_RECEIVERSETUP_1_4_RXSERIALNBR_LENGTH]byte
	RxName        [SBF_RECEIVERSETUP_1_4_RXNAME_LENGTH]byte
	RxVersion     [SBF_RECEIVERSETUP_1_4_RXVERSION_LENGTH]byte
	AntSerialNbr  [SBF_RECEIVERSETUP_1_4_ANTSERIALNBR_LENGTH]byte
	AntType       [SBF_RECEIVERSETUP_1_4_ANTTYPE_LENGTH]byte
	deltaH        float32 /* [m] */
	deltaE        float32 /* [m] */
	deltaN        float32 /* [m] */
	MarkerType    [SBF_RECEIVERSETUP_1_4_MARKERTYPE_LENGTH]byte
	GNSSFWVersion [SBF_RECEIVERSETUP_1_4_GNSSFWVERSION_LENGTH]byte
	ProductName   [SBF_RECEIVERSETUP_1_4_PRODUCTNAME_LENGTH]byte
	Latitude      SBFDOUBLE
	Longitude     SBFDOUBLE
	Height        float32
	StationCode   [SBF_RECEIVERSETUP_1_4_STATIONCODE_LENGTH]byte
	MonumentIdx   uint8
	ReceiverIdx   uint8
	CountryCode   [SBF_RECEIVERSETUP_1_4_COUNTRYCODE_LENGTH]byte
	Reserved1     [SBF_RECEIVERSETUP_1_4_RESERVED1_LENGTH]byte
}
type ReceiverSetup_1_t ReceiverSetup_1_4_t

/*--RxComponents_1_0_t : ----------------------------------------------------*/
/* Information on various receiver components */

/** Component_1_0_t */
type Component_1_0_t struct {
	Type         uint8
	CPULoad      uint8
	Reserved     [2]uint8
	Name         [SBF_COMPONENT_1_0_NAME_LENGTH]byte
	SerialNumber [SBF_COMPONENT_1_0_SERIALNUMBER_LENGTH]byte
	FWVersion    [SBF_COMPONENT_1_0_FWVERSION_LENGTH]byte
	MACAddress   [SBF_COMPONENT_1_0_MACADDRESS_LENGTH]uint8

	_padding [SBF_COMPONENT_1_0__PADDING_LENGTH]uint8
}

/** RxComponents_1_0_t */
type RxComponents_1_0_t struct {
	Header BlockHeader_t

	/* Time Header */
	TOW uint32
	WNc uint16

	N         uint8
	SBLength  uint8
	Reserved  [4]uint8
	Component [SBF_RXCOMPONENTS_1_0_COMPONENT_LENGTH]Component_1_0_t
}

type Component_1_t Component_1_0_t
type RxComponents_1_t RxComponents_1_0_t

/*--RxMessage_1_0_t : -------------------------------------------------------*/
/* Receiver message */

/** RxMessage_1_0_t */
type RxMessage_1_0_t struct {
	Header BlockHeader_t

	/* Time Header */
	TOW uint32
	WNc uint16

	Type      uint8
	Severity  uint8
	MessageID uint32
	StringLn  uint16
	Reserved2 [SBF_RXMESSAGE_1_0_RESERVED2_LENGTH]uint8
	Message   [SBF_RXMESSAGE_1_0_MESSAGE_LENGTH]byte
}
type RxMessage_1_t RxMessage_1_0_t

/*--Commands_1_0_t : --------------------------------------------------------*/
/* Commands entered by the user */

/** Commands_1_0_t */
type Commands_1_0_t struct {
	Header BlockHeader_t

	/* Time Header */
	TOW uint32
	WNc uint16

	Reserved [2]uint8
	CmdData  [SBF_COMMANDS_1_0_CMDDATA_LENGTH]uint8
}
type Commands_1_t Commands_1_0_t

/*--Comment_1_0_t : ---------------------------------------------------------*/
/* Comment entered by the user */

/** Comment_1_0_t */
type Comment_1_0_t struct {
	Header BlockHeader_t

	/* Time Header */
	TOW uint32
	WNc uint16

	CommentLn uint16
	Comment   [SBF_COMMENT_1_0_COMMENT_LENGTH]byte
}
type Comment_1_t Comment_1_0_t

/*--BBSamples_1_0_t : -------------------------------------------------------*/
/* Baseband samples */

/** BBSamplesData_1_0_t */
type BBSamplesData_1_0_t struct {
	/* Time Header */
	TOW uint32
	WNc uint16

	N          uint16
	Info       uint8
	Reserved   [3]uint8
	SampleFreq uint32
	LOFreq     uint32
	Samples    [SBF_BBSAMPLESDATA_1_0_SAMPLES_LENGTH]uint16
}

/** BBSamples_1_0_t */
type BBSamples_1_0_t struct {
	Header BlockHeader_t

	BBSamplesData BBSamplesData_1_0_t
}

type BBSamplesData_1_t BBSamplesData_1_0_t

type BBSamples_1_t BBSamples_1_0_t

/*--ASCIIIn_1_0_t : ---------------------------------------------------------*/
/* ASCII input from external sensor */

/** ASCIIIn_1_0_t */
type ASCIIIn_1_0_t struct {
	Header BlockHeader_t

	/* Time Header */
	TOW uint32
	WNc uint16

	CD          uint8
	Reserved1   [3]uint8
	StringLn    uint16
	SensorModel [SBF_ASCIIIN_1_0_SENSORMODEL_LENGTH]byte
	SensorType  [SBF_ASCIIIN_1_0_SENSORTYPE_LENGTH]byte
	Reserved2   [20]uint8
	ASCIIString [SBF_ASCIIIN_1_0_ASCIISTRING_LENGTH]byte
}
type ASCIIIn_1_t ASCIIIn_1_0_t

/*--EncapsulatedOutput_1_0_t : ----------------------------------------------*/
/* SBF encapsulation of non-SBF messages */
/** EncapsulatedOutput_1_0_t */
type EncapsulatedOutput_1_0_t struct {
	Header BlockHeader_t

	/* Time Header */
	TOW uint32
	WNc uint16

	Mode       uint8
	Reserved   uint8
	N          uint16
	ReservedId uint16
	Payload    [SBF_ENCAPSULATEDOUTPUT_1_0_PAYLOAD_LENGTH]uint8
}
type EncapsulatedOutput_1_t EncapsulatedOutput_1_0_t

/*--RawDataIn_1_0_t : -------------------------------------------------------*/
/* Incoming raw data message */

/** RawDataIn_1_0_t */
type RawDataIn_1_0_t struct {
	Header BlockHeader_t

	/* Time Header */
	TOW uint32
	WNc uint16

	Mode   uint8
	Source uint8
	Bytes  [SBF_RAWDATAIN_1_0_BYTES_LENGTH]uint8
}
type RawDataIn_1_t RawDataIn_1_0_t

/*--TURPVTSatCorrections_1_0_t : --------------------------------------------*/
/* Satellite range corrections */

/** SatClkCorrInfo_1_0_t */
type SatClkCorrInfo_1_0_t struct {
	Type       uint8
	Reserved   [3]uint8
	ClockBias  float32
	ClockDrift float32
}

/** SatCorrInfo_1_0_t */
type SatCorrInfo_1_0_t struct {
	SVID       uint8
	FreqNr     uint8
	N2         uint8
	Reserved   uint8
	TropoDelay float32
	IonoDelay  float32
	RelCorr    float32
	PRC        float32
}

/** TURPVTSatCorrections_1_0_t */
type TURPVTSatCorrections_1_0_t struct {
	Header BlockHeader_t

	/* Time Header */
	TOW uint32
	WNc uint16

	N         uint8
	SB1Length uint8
	SB2Length uint8
	Reserved  [3]uint8
	Data      [SBF_TURPVTSATCORRECTIONS_1_0_DATA_LENGTH]uint8
}

/*--TURPVTSatCorrections_1_1_t : --------------------------------------------*/
/* Satellite range corrections */

/** SatClkCorrInfo_1_1_t */
type SatClkCorrInfo_1_1_t struct {
	Type       uint8
	Reserved   [3]uint8
	ClockBias  float32
	ClockDrift float32
}

/** SatCorrInfo_1_1_t */
type SatCorrInfo_1_1_t struct {
	SVID       uint8
	FreqNr     uint8
	N2         uint8
	Reserved   uint8
	TropoDelay float32
	IonoDelay  float32
	RelCorr    float32
	PRC        float32
	EarthRot_X float32
	EarthRot_Y float32
}

/** TURPVTSatCorrections_1_1_t */
type TURPVTSatCorrections_1_1_t struct {
	Header BlockHeader_t

	/* Time Header */
	TOW uint32
	WNc uint16

	N         uint8
	SB1Length uint8
	SB2Length uint8
	Reserved  [3]uint8
	Data      [SBF_TURPVTSATCORRECTIONS_1_1_DATA_LENGTH]uint8
}

type SatClkCorrInfo_1_t SatClkCorrInfo_1_1_t

type SatCorrInfo_1_t SatCorrInfo_1_1_t
type TURPVTSatCorrections_1_t TURPVTSatCorrections_1_1_t

/*--TURHPCAInfo_1_0_t : -----------------------------------------------------*/
/* HMI Probability information */

/** SatHpca_1_0_t */
type SatHpca_1_0_t struct {
	SVID    uint8
	info    uint8
	SISA    uint8
	SISMA   uint8
	HorDev  float32
	VerDev  float32
	HMIProb SBFDOUBLE
	NavAge  uint16
	FreqNr  uint8

	_padding [SBF_SATHPCA_1_0__PADDING_LENGTH]uint8
}

/** TURHPCAInfo_1_0_t */
type TURHPCAInfo_1_0_t struct {
	Header BlockHeader_t

	/* Time Header */
	TOW uint32
	WNc uint16

	N        uint8
	SBSize   uint8
	NWA_I    uint8
	NrSVc    uint8
	NrLink   uint8
	Reserved uint8
	HMIProb  SBFDOUBLE
	HorDev   float32
	VerDev   float32
	HPCAData [SBF_TURHPCAINFO_1_0_HPCADATA_LENGTH]SatHpca_1_0_t
}

type SatHpca_1_t SatHpca_1_0_t
type TURHPCAInfo_1_t TURHPCAInfo_1_0_t

/*--CorrPeakSample_1_0_t : --------------------------------------------------*/
/* Real-time samples of the correlation peak function */

/** CorrSample_1_0_t */
type CorrSample_1_0_t struct {
	Offset int16
	I      int8
	Q      int8
}

/** CorrPeakSample_1_0_t */
type CorrPeakSample_1_0_t struct {
	Header BlockHeader_t

	/* Time Header */
	TOW uint32
	WNc uint16

	N              uint8
	SBSize         uint8
	RxChannel      uint8
	Type           uint8
	SVID           uint8
	Reserved       uint8
	CorrSampleData [SBF_CORRPEAKSAMPLE_1_0_CORRSAMPLEDATA_LENGTH]CorrSample_1_0_t
}

type CorrSample_1_t CorrSample_1_0_t
type CorrPeakSample_1_t CorrPeakSample_1_0_t

/*--CorrValues_1_0_t : ------------------------------------------------------*/
/* Raw correlation values */

/** CorrValue_1_0_t */
type CorrValue_1_0_t struct {
	dT       uint8
	Type     uint8
	CorrInfo uint16
	I        int16
	Q        int16
}

/** CorrValues_1_0_t */
type CorrValues_1_0_t struct {
	Header BlockHeader_t

	/* Time Header */
	TOW uint32
	WNc uint16

	N                 uint16
	SBLength          uint8
	RxChannel         uint8
	SVID              uint8
	Reserved          [3]uint8
	CorrelationValues [SBF_CORRVALUES_1_0_CORRELATIONVALUES_LENGTH]CorrValue_1_0_t

	_padding [SBF_CORRVALUES_1_0__PADDING_LENGTH]uint8
}

type CorrValue_1_t CorrValue_1_0_t
type CorrValues_1_t CorrValues_1_0_t

/*--TURStatus_1_0_t : -------------------------------------------------------*/
/* TUR-specific status information */

/** TURStatus_1_0_t */
type TURStatus_1_0_t struct {
	Header BlockHeader_t

	/* Time Header */
	TOW uint32
	WNc uint16

	RxTemp  int8
	AntTemp int8
}

/*--TURStatus_1_1_t : -------------------------------------------------------*/
/* TUR-specific status information */

/** TURStatus_1_1_t */
type TURStatus_1_1_t struct {
	Header BlockHeader_t

	/* Time Header */
	TOW uint32
	WNc uint16

	RxTemp          int8
	AntTemp         int8
	EventCountTB1   uint32
	EventCountTB2   uint32
	EventCountTB3   uint32
	StartIFL_TOW    uint32
	StartIFL_WNc    uint16
	StartIFL_Status uint8
	Reserved1       uint8
	StopIFL_TOW     uint32
	StopIFL_WNc     uint16
	StopIFL_Status  uint8
	Reserved2       uint8
}

/*--TURStatus_1_2_t : -------------------------------------------------------*/
/* TUR-specific status information */

/** TURStatus_1_2_t */
type TURStatus_1_2_t struct {
	Header BlockHeader_t

	/* Time Header */
	TOW uint32
	WNc uint16

	RxTemp          int8
	AntTemp         int8
	EventCountTB1   uint32
	EventCountTB2   uint32
	EventCountTB3   uint32
	StartIFL_TOW    uint32
	StartIFL_WNc    uint16
	StartIFL_Status uint8
	Reserved1       uint8
	StopIFL_TOW     uint32
	StopIFL_WNc     uint16
	StopIFL_Status  uint8
	Reserved2       uint8
	I_E5            int16
	V_E5            int16
	I_E5a           int16
	V_E5a           int16
	I_E5b           int16
	V_E5b           int16
	I_E6            int16
	V_E6            int16
	I_L1            int16
	V_L1            int16
	I_12V_Ant       int16
	V_12V_Ant       int16
	I_12V           int16
	V_12V           int16
	I_5V            int16
	V_5V            int16
	I_5V_OCXO       int16
	V_5V_OCXO       int16
}

type TURStatus_1_t TURStatus_1_2_t

/*--GALIntegrity_1_0_t : ----------------------------------------------------*/
/* Galileo integrity data */

/** gaIntegrity_1_0_t */
type gaIntegrity_1_0_t struct {
	/* Time Header */
	TOW uint32
	WNc uint16

	SVID          uint8  /* SBF range 71-102 SIS range 1-64 [F1/I4/G1] */
	Source        uint8  /* bitfield according to sigType */
	ToC           uint32 /* Time of Computation in sec of week */
	RegionStatus  uint8  /* 8-bit region status */
	Authenticated uint8  /* 1 : succeeded, 0 otherwise */
	CleverWord    uint8
	UpdateEvent   uint8
	IntegritySVi  [SBF_GAINTEGRITY_1_0_INTEGRITYSVI_LENGTH]uint8 /* Bit 0-3 : IF bits
	Bit 4-6 : SNF info
	Bit   7 : reserved */
}

/** GALIntegrity_1_0_t */
type GALIntegrity_1_0_t struct {
	Header BlockHeader_t

	galIntegrity gaIntegrity_1_0_t
}

type gaIntegrity_1_t gaIntegrity_1_0_t

type GALIntegrity_1_t GALIntegrity_1_0_t

/*--TURFormat_1_0_t : -------------------------------------------------------*/

/** TURFormatMode1RGCrequest_1_0_t */
type TURFormatMode1RGCrequest_1_0_t struct {
	uCodeGen1 uint8
	uCodeGen2 uint8
	uCodeGen3 uint8
	uSatNo    uint8
	uFreq     uint8
	uChipRate uint8
	uWeekNr   uint16
	uTOW      uint16
	uPeriod   uint16
	uLatency  uint16

	_padding [SBF_TURFORMATMODE1RGCREQUEST_1_0__PADDING_LENGTH]uint8
}

/** TURFormatRangingCodeReady_1_0_t */
type TURFormatRangingCodeReady_1_0_t struct {
	uCodeGen uint8
	uPadding uint8

	_padding [SBF_TURFORMATRANGINGCODEREADY_1_0__PADDING_LENGTH]uint8
}

/** TURFormatMode2RGCrequest_1_0_t */
type TURFormatMode2RGCrequest_1_0_t struct {
	uCodeGen     uint8
	uStartParam  uint8
	uStep        uint8
	uLength      uint8
	uSatNo       uint8
	uFreq        uint8
	uChipRate    uint8
	uWeekNr      uint16
	uTOW         uint16
	uM1parameter uint8
}

/** TURFormatGetMode2Time_1_0_t */
type TURFormatGetMode2Time_1_0_t struct {
	uCodeGen  uint8
	uM2hit    uint8
	uSatNo    uint8
	uFreq     uint8
	uChipRate uint8
	uWeekNr   uint16
	uTOW      uint16
	uPadding  uint8

	_padding [SBF_TURFORMATGETMODE2TIME_1_0__PADDING_LENGTH]uint8
}

/** TURFormatMode2TimeReply_1_0_t */
type TURFormatMode2TimeReply_1_0_t struct {
	uRefCount uint16
	uIndex    uint32

	_padding [SBF_TURFORMATMODE2TIMEREPLY_1_0__PADDING_LENGTH]uint8
}

/** TURFormatRangingCodeRequest_1_0_t */
type TURFormatRangingCodeRequest_1_0_t struct {
	uCodeGen     uint8
	uSatNo       uint8
	uFreq        uint8
	uChipRate    uint8
	uWeekNr      uint16
	uTOW         uint16
	uBlock       uint32
	uM1Parameter uint8

	_padding [SBF_TURFORMATRANGINGCODEREQUEST_1_0__PADDING_LENGTH]uint8
}

/** TURFormatEncryptedData_1_0_t */
type TURFormatEncryptedData_1_0_t struct {
	uSatNo   uint8
	uFreq    uint8
	uWeekNr  uint16
	uTOW     uint32
	uData    [SBF_TURFORMATENCRYPTEDDATA_1_0_UDATA_LENGTH]uint8
	uPadding uint8

	_padding [SBF_TURFORMATENCRYPTEDDATA_1_0__PADDING_LENGTH]uint8
}

/** TURFormatDecryptedData_1_0_t */
type TURFormatDecryptedData_1_0_t struct {
	uRefCount uint8
	uData     [SBF_TURFORMATDECRYPTEDDATA_1_0_UDATA_LENGTH]uint8
	uPadding  uint8

	_padding [SBF_TURFORMATDECRYPTEDDATA_1_0__PADDING_LENGTH]uint8
}

/** TURFormatPVT_1_0_t */
type TURFormatPVT_1_0_t struct {
	iLatitude    int32
	iLongitude   int32
	uVelocity    uint32
	iBearing     int32
	iAltitude    int32
	uWeekNumber  uint16
	uTimeOfWeek  uint32
	uMillisecond uint16
	uUTCOffeset  uint16
	uTimeType    uint8
	uPadding     uint8
}

/** TURFormatDenial_1_0_t */
type TURFormatDenial_1_0_t struct {
	uDenialFlag uint8
	uPadding    uint8

	_padding [SBF_TURFORMATDENIAL_1_0__PADDING_LENGTH]uint8
}

/** TURFormatEventLog_1_0_t */
type TURFormatEventLog_1_0_t struct {
	Format    uint8
	EventType uint8
	WN        uint16
	TOW       uint32
}

/** TURFormatVersionResponse_1_0_t */
type TURFormatVersionResponse_1_0_t struct {
	uSerialNumberMSB uint16
	uSerialNumberLSB uint32
	uSMVersionMajor  uint16
	uSMVersionMinor1 uint16
	uSMVersionMinor2 uint16
}

/** TURFormatStatusResponse_1_0_t */
type TURFormatStatusResponse_1_0_t struct {
	uStatus  uint8
	uPadding uint8

	_padding [SBF_TURFORMATSTATUSRESPONSE_1_0__PADDING_LENGTH]uint8
}

/** TURFormatOptionA_1_0_t */
type TURFormatOptionA_1_0_t struct {
	uWN1       uint16
	uTOW1      uint32
	uOptionA_1 uint8
	uWN2       uint16
	uTOW2      uint32
	uOptionA_2 uint8

	_padding [SBF_TURFORMATOPTIONA_1_0__PADDING_LENGTH]uint8
}

/** TURFormatOptionB_1_0_t */
type TURFormatOptionB_1_0_t struct {
	uMask1     [SBF_TURFORMATOPTIONB_1_0_UMASK1_LENGTH]uint8
	uWN1       uint16
	uTOW1      uint32
	uOptionB_1 uint8
	uMask2     [SBF_TURFORMATOPTIONB_1_0_UMASK2_LENGTH]uint8
	uWN2       uint16
	uTOW2      uint32
	uOptionB_2 uint8
}

/** TURFormat_1_0_t */
type TURFormat_1_0_t struct {
	Header BlockHeader_t

	/* Time Header */
	TOW uint32
	WNc uint16

	Reserved                     [2]uint8
	Type                         uint8
	Payload_Length               uint8
	Message_Count                uint16
	Message_Mode_1_RCG_Request   TURFormatMode1RGCrequest_1_0_t
	Message_Ranging_Code_Ready   TURFormatRangingCodeReady_1_0_t
	Message_Mode_2_RCG_Request   TURFormatMode2RGCrequest_1_0_t
	Message_Get_Mode_2_Time      TURFormatGetMode2Time_1_0_t
	Message_Mode_2_Time_Reply    TURFormatMode2TimeReply_1_0_t
	Message_Ranging_Code_Request TURFormatRangingCodeRequest_1_0_t
	Message_Encrypted_Data       TURFormatEncryptedData_1_0_t
	Message_Decrypted_Data       TURFormatDecryptedData_1_0_t
	Message_PVT                  TURFormatPVT_1_0_t
	Message_Denial               TURFormatDenial_1_0_t
	Message_Event_Log            TURFormatEventLog_1_0_t
	Message_Version_Response     TURFormatVersionResponse_1_0_t
	Message_Status_Response      TURFormatStatusResponse_1_0_t
	OptionA                      TURFormatOptionA_1_0_t
	OptionB                      TURFormatOptionB_1_0_t
}

type TURFormatMode1RGCrequest_1_t TURFormatMode1RGCrequest_1_0_t
type TURFormatRangingCodeReady_1_t TURFormatRangingCodeReady_1_0_t
type TURFormatMode2RGCrequest_1_t TURFormatMode2RGCrequest_1_0_t
type TURFormatGetMode2Time_1_t TURFormatGetMode2Time_1_0_t
type TURFormatMode2TimeReply_1_t TURFormatMode2TimeReply_1_0_t
type TURFormatRangingCodeRequest_1_t TURFormatRangingCodeRequest_1_0_t
type TURFormatEncryptedData_1_t TURFormatEncryptedData_1_0_t
type TURFormatDecryptedData_1_t TURFormatDecryptedData_1_0_t
type TURFormatPVT_1_t TURFormatPVT_1_0_t
type TURFormatDenial_1_t TURFormatDenial_1_0_t
type TURFormatEventLog_1_t TURFormatEventLog_1_0_t
type TURFormatVersionResponse_1_t TURFormatVersionResponse_1_0_t
type TURFormatStatusResponse_1_t TURFormatStatusResponse_1_0_t
type TURFormatOptionA_1_t TURFormatOptionA_1_0_t
type TURFormatOptionB_1_t TURFormatOptionB_1_0_t
type TURFormat_1_t TURFormat_1_0_t

/*--CalibrationValues_1_0_t : -----------------------------------------------*/
/* Calibration values */

/** CalibrationInfoType_1_0_t */
type CalibrationInfoType_1_0_t struct {
	SignalType    uint8
	Reserved      [3]uint8
	ManualValue   int32
	MeasuredValue int32
}

/** CalibrationValues_1_0_t */
type CalibrationValues_1_0_t struct {
	Header BlockHeader_t

	/* Time Header */
	TOW uint32
	WNc uint16

	Flags           uint8
	Reserved        uint8
	TOW_last        uint32
	WNc_last        uint16
	N               uint8
	SBLength        uint8
	CalibrationInfo [SBF_CALIBRATIONVALUES_1_0_CALIBRATIONINFO_LENGTH]CalibrationInfoType_1_0_t
}

type CalibrationInfoType_1_t CalibrationInfoType_1_0_t
type CalibrationValues_1_t CalibrationValues_1_0_t

/*--MultipathMonitor_1_0_t : ------------------------------------------------*/
/* MultipathMonitorSub blocks */

/** MultipathMonitorSubType_1_0_t */
type MultipathMonitorSubType_1_0_t struct {
	RxChannel  uint8
	Type       uint8
	SVID       uint8
	Reserved   uint8
	MPEstimate float32
}

/** MultipathMonitor_1_0_t */
type MultipathMonitor_1_0_t struct {
	Header BlockHeader_t

	/* Time Header */
	TOW uint32
	WNc uint16

	N                   uint8
	SBLength            uint8
	MultipathMonitorSub [SBF_MULTIPATHMONITOR_1_0_MULTIPATHMONITORSUB_LENGTH]MultipathMonitorSubType_1_0_t
}

type MultipathMonitorSubType_1_t MultipathMonitorSubType_1_0_t
type MultipathMonitor_1_t MultipathMonitor_1_0_t

/*--FOCTURNStatus_1_0_t : ---------------------------------------------------*/

/** FOCTURNStatus_1_0_t */
type FOCTURNStatus_1_0_t struct {
	Header BlockHeader_t

	/* Time Header */
	TOW uint32
	WNc uint16

	DIG_FpgaTemp       int16
	DIG_VCCINT         uint16
	DIG_VCCAUX         uint16
	DIG_VCCBRAM        uint16
	DIG_VCC_12V_DIG_SW uint16
	DIG_VCC_5V_SM      uint16
	DIG_VCC_FPGA_1V5   uint16
	DIG_VCC_3V7        uint16
	DIG_VCC_12V_RF_SW  uint16
	DIG_VCC_5V1        uint16
	DIG_VCC_2V5        uint16
	RF_FpgaTemp        int16
	RF_VCCINT          uint16
	RF_VCCAUX          uint16
	RF_VCCBRAM         uint16
	RF_VCC_12V         uint16
	RF_VCC_5V5         uint16
	RF_VCC_3V6_RF      uint16
	Reserved           uint16
	TTFF               uint32
	TTFAF              uint32
}

type FOCTURNStatus_1_t FOCTURNStatus_1_0_t

/*--TGVFXStatus_1_0_t : -----------------------------------------------------*/

/** TGVFXStatus_1_0_t */
type TGVFXStatus_1_0_t struct {
	Header BlockHeader_t

	/* Time Header */
	TOW uint32
	WNc uint16

	Reserved uint16
	TTFF     uint32
	TTFAF    uint32
}

type TGVFXStatus_1_t TGVFXStatus_1_0_t

/*==PinPoint-GIS RX==========================================================*/

/*--GISAction_1_0_t : -------------------------------------------------------*/
/* PinPoint-GIS RX Action */

/** GISAction_1_0_t */
type GISAction_1_0_t struct {
	Header BlockHeader_t

	/* Time Header */
	TOW uint32
	WNc uint16

	CommentLn uint16
	ItemIDMSB uint32
	ItemIDLSB uint32
	Action    uint8
	Trigger   uint8
	Database  uint8
	Reserved  uint8
	Comment   [SBF_GISACTION_1_0_COMMENT_LENGTH]byte
	_padding  [SBF_GISACTION_1_0__PADDING_LENGTH]uint8
}

type GISAction_1_t GISAction_1_0_t

/*--GISStatus_1_0_t : -------------------------------------------------------*/
/* Status of the different PinPoint-GIS collection databases */

/** DatabaseStatus_1_0_t */
type DatabaseStatus_1_0_t struct {
	Database     uint8
	OnlineStatus uint8
	Error        uint8
	Reserved     uint8
	NrItems      uint32
	NrNotSync    uint32
}

/** GISStatus_1_0_t */
type GISStatus_1_0_t struct {
	Header BlockHeader_t

	/* Time Header */
	TOW uint32
	WNc uint16

	N              uint8
	SBLength       uint8
	DatabaseStatus [SBF_GISSTATUS_1_0_DATABASESTATUS_LENGTH]DatabaseStatus_1_0_t
}

type DatabaseStatus_1_t DatabaseStatus_1_0_t
type GISStatus_1_t GISStatus_1_0_t

/*==SBF-BLOCKS definition====================================================*/

/*--MeasExtra_1_0_t : -------------------------------------------------------*/
type MeasExtraChannel_1_0_t MeasExtraChannelSub_1_0_t

/*--MeasExtra_1_1_t : -------------------------------------------------------*/
type MeasExtraChannel_1_1_t MeasExtraChannelSub_1_1_t

/*--MeasExtra_1_2_t : -------------------------------------------------------*/
type MeasExtraChannel_1_2_t MeasExtraChannelSub_1_2_t

/*--MeasExtra_1_3_t : -------------------------------------------------------*/
type MeasExtraChannel_1_3_t MeasExtraChannelSub_1_3_t
type MeasExtraChannel_1_t MeasExtraChannelSub_1_t

/*--IQCorr_1_0_t : ----------------------------------------------------------*/
type CorrChannel_1_0_t IQCorrChannelSub_1_0_t

/*--IQCorr_1_1_t : ----------------------------------------------------------*/
type CorrChannel_1_1_t IQCorrChannelSub_1_1_t
type CorrChannel_1_t IQCorrChannelSub_1_t

/*==Navigation Page Blocks===================================================*/

/*--GPSRawCA_1_0_t : --------------------------------------------------------*/
type NavBits_1_0_t GPSRawCA_1_0_t
type NavBits_1_t NavBits_1_0_t

/* For backwards compatibility after renaming CMP to BDS */
type CMPRaw_1_0_t BDSRaw_1_0_t
type CMPRaw_1_t BDSRaw_1_t
type CMPNav_1_0_t BDSNav_1_0_t
type CMPNav_1_t BDSNav_1_t

/* For backwards compatibility after renaming IRNSS to NAVIC */
type IRNSSRaw_1_0_t NAVICRaw_1_0_t
type IRNSSRaw_1_t NAVICRaw_1_t

/*==SBAS Decoded Message Blocks==============================================*/

/*--GEOServiceLevel_1_0_t : -------------------------------------------------*/
type raServiceRegion_1_0_t ServiceRegion_1_0_t
type raServiceRegion_1_t ServiceRegion_1_t

/*==Position, Velocity and Time Blocks=======================================*/
/*--PVTResiduals_1_0_t : ----------------------------------------------------*/
type SatResidual_1_0_t PVTResidual_1_0_t
type SatResidual_1_t PVTResidual_1_t

/*--PVTResiduals_2_0_t : ----------------------------------------------------*/
type SatResInfo_2_0_t ResidualInfoCode_2_0_t

// same as ResidualInfoPhase_2_0_t and ResidualInfoDoppler_2_0_t

/*--PVTResiduals_2_1_t : ----------------------------------------------------*/
type SatResInfo_2_1_t ResidualInfoCode_2_1_t

// same as ResidualInfoPhase_2_1_t and ResidualInfoDoppler_2_1_t
type SatResInfo_2_t SatResInfo_2_1_t

/*--RAIMStatistics_1_0_t : --------------------------------------------------*/
type RAIMStatChannel_1_0_t RAIMSatData_1_0_t
type RAIMStatChannel_1_t RAIMSatData_1_t

/*--GEOCorrections_1_0_t : --------------------------------------------------*/
type GeoCorrChannel_1_0_t SatCorr_1_0_t
type GeoCorrChannel_1_t SatCorr_1_t
type GeoCorrections_1_0_t GEOCorrections_1_0_t
type GeoCorrections_1_t GEOCorrections_1_t

/*==INS/GNSS Integrated Blocks===============================================*/

/*--IntPVCart_1_0_t : -------------------------------------------------------*/
type intPVCart_1_0_t IntPVCart_1_0_t
type intPVCart_1_t IntPVCart_1_t

/*--IntPVGeod_1_0_t : -------------------------------------------------------*/
type intPVGeod_1_0_t IntPVGeod_1_0_t
type intPVGeod_1_t IntPVGeod_1_t

/*--IntPosCovCart_1_0_t : ---------------------------------------------------*/
type intPosCovCart_1_0_t IntPosCovCart_1_0_t
type intPosCovCart_1_t IntPosCovCart_1_t

/*--IntVelCovCart_1_0_t : ---------------------------------------------------*/
type intVelCovCart_1_0_t IntVelCovCart_1_0_t
type intVelCovCart_1_t IntVelCovCart_1_t

/*--IntPosCovGeod_1_0_t : ---------------------------------------------------*/
type intPosCovGeod_1_0_t IntPosCovGeod_1_0_t
type intPosCovGeod_1_t IntPosCovGeod_1_t

/*--IntVelCovGeod_1_0_t : ---------------------------------------------------*/
type intVelCovGeod_1_0_t IntVelCovGeod_1_0_t
type intVelCovGeod_1_t IntVelCovGeod_1_t

/*--IntAttEuler_1_0_t : -----------------------------------------------------*/
type intAttEuler_1_0_t IntAttEuler_1_0_t
type intAttEuler_1_t IntAttEuler_1_t

/*--IntAttCovEuler_1_0_t : --------------------------------------------------*/
type intAttCovEuler_1_0_t IntAttCovEuler_1_0_t
type intAttCovEuler_1_t IntAttCovEuler_1_t

/*--IntPVAAGeod_1_0_t : -----------------------------------------------------*/
type intPVAAGeod_1_0_t IntPVAAGeod_1_0_t
type intPVAAGeod_1_t IntPVAAGeod_1_t

/*==GNSS Attitude Blocks=====================================================*/

/*--AuxAntPositions_1_0_t : -------------------------------------------------*/
type AuxAntPosData_1_0_t AuxAntPositionSub_1_0_t
type AuxAntPosData_1_t AuxAntPositionSub_1_t

/*==L-Band Demodulator Blocks================================================*/

/*--LBandTrackerStatus_1_0_t : ----------------------------------------------*/
type LBandTrackData_1_0_t TrackData_1_0_t

/*--LBandTrackerStatus_1_1_t : ----------------------------------------------*/
type LBandTrackData_1_1_t TrackData_1_1_t

/*--LBandTrackerStatus_1_1_2 : ----------------------------------------------*/
type LBandTrackData_1_2_t TrackData_1_2_t

/*--LBandTrackerStatus_1_1_3 : ----------------------------------------------*/
type LBandTrackData_1_3_t TrackData_1_3_t
type LBandTrackData_1_t TrackData_1_t

/*==External Sensor Blocks===================================================*/

/*--ExtSensorInfo_1_0_t : ---------------------------------------------------*/
type ExtSensorSetupData_1_0_t OneSensor_1_0_t

/*--ExtSensorSetup_1_1_t : --------------------------------------------------*/
type ExtSensorSetupData_1_1_t OneSensor_1_1_t
type ExtSensorSetupData_1_t OneSensor_1_t

/*--ExtSensorSetup_1_2_t : --------------------------------------------------*/
type ExtSensorSetupData_1_2_t OneSensor_1_2_t

/*==Status Blocks============================================================*/

/*--InputLink_1_0_t : -------------------------------------------------------*/
type InputStats_1_0_t InputStatsSub_1_0_t
type InputStats_1_t InputStatsSub_1_t

/*--OutputLink_1_0_t : ------------------------------------------------------*/
type OutputTypes_1_0_t OutputTypeSub_1_0_t
type OutputStats_1_0_t OutputStatsSub_1_0_t

/*--OutputLink_1_1_t : ------------------------------------------------------*/
type OutputTypes_1_1_t OutputTypeSub_1_1_t
type OutputStats_1_1_t OutputStatsSub_1_1_t
type OutputTypes_1_t OutputTypeSub_1_t
type OutputStats_1_t OutputStatsSub_1_t

/*--TURHPCAInfo_1_0_t : -----------------------------------------------------*/
type SatHPCA_1_0_t SatHpca_1_0_t
type SatHPCA_1_t SatHpca_1_t

/*--SysTimeOffset_1_0_t : ---------------------------------------------------*/
type TimeOffset_1_0_t TimeOffsetSub_1_0_t

/*--SysTimeOffset_1_1_t : ---------------------------------------------------*/
type TimeOffset_1_1_t TimeOffsetSub_1_1_t
type TimeOffset_1_t TimeOffsetSub_1_t

/*==Other Blocks=============================================================*/

/*--sbsAlm_1_0_t : -----------------------------------------------------*/
type SbasL5Alm_t sbsAlm_1_0_t

/* These types are added for backwards compatibility reasons.
Please do not use them in new development.
*/
type MeasEpochChannelType2_t MeasEpochChannelType2_1_0_t
type MeasEpochChannelType1_t MeasEpochChannelType1_1_0_t
type MeasExtraChannel_t MeasExtraChannel_1_0_t
type gpEphLNAV_t gpEph_1_0_t
type qzEphLNAV_t qzEph_1_0_t
type gpEphCNAV_t gpEphCNAV_1_0_t
type gpAlm_t gpAlm_1_0_t
type qzAlm_t qzAlm_1_0_t
type gpIon_t gpIon_1_0_t
type gpUtc_t gpUtc_1_0_t
type glEph_t glEph_1_0_t
type glAlm_t glAlm_1_0_t
type glTime_t glTime_1_0_t
type galEph_t gaEph_1_0_t
type galAlm_t gaAlm_1_0_t
type galIon_t gaIon_1_0_t
type galUtc_t gaUtc_1_0_t
type galGstGps_t gaGstGps_1_0_t
type cmpEph_t cmpEph_1_0_t
type cmpAlm_t cmpAlm_1_0_t
type cmpIon_t cmpIon_1_0_t
type cmpUtc_t cmpUtc_1_0_t
type GeoMT00_t raMT00_1_0_t
type GeoPRNMask_t raPRNMask_1_0_t
type FastCorr_t FastCorr_1_0_t
type GeoFastCorr_t raFastCorr_1_0_t
type GeoIntegrity_t raIntegrity_1_0_t
type GeoFastCorrDegr_t raFastCorrDegr_1_0_t
type raEph_t raEph_1_0_t
type raDF_t raDF_1_0_t
type GeoDegrFactors_t raDegrFactors_1_0_t
type raNetworkTime_t NetworkTimeMsg_1_0_t
type GeoNetworkTime_t raNetworkTime_1_0_t
type raAlm_t raAlm_1_0_t
type GeoIGPMask_t raIGPMask_1_0_t
type LTCorr_t LTCorr_1_0_t
type GeoLongTermCorr_t raLongTermCorr_1_0_t
type IDC_t IDC_1_0_t
type GeoIonoDelay_t raIonoDelay_1_0_t
type raServiceRegion_t raServiceRegion_1_0_t
type raServiceLevel_t raServiceMsg_1_0_t
type GeoServiceLevel_t raServiceLevel_1_0_t
type CovMatrix_t CovMatrix_1_0_t
type GeoClockEphCovMatrix_t raClockEphCovMatrix_1_0_t
type SatPos_t SatPos_1_0_t
type SatResidual_t SatResidual_1_0_t
type SatSignalInfo_t SatSignalInfo_2_0_t
type SatResInfo_t SatResInfo_2_0_t
type RAIMStatChannel_t RAIMStatChannel_1_0_t
type GeoCorrChannel_t GeoCorrChannel_1_0_t
type PPSData_t PPSData_1_0_t
type TimerData_t TimerData_1_0_t
type TrackingStatusChannel_t TrackingStatusChannel_1_0_t
type ChannelSatInfo_t ChannelSatInfo_1_0_t
type ChannelStateInfo_t ChannelStateInfo_1_0_t
type AGCData_t AGCData_1_0_t
type AGCState_t AGCState_2_0_t
type SatInfo_t SatInfo_1_0_t
type SatCorrInfo_tv SatCorrInfo_1_1_t
type SatClkCorrInfo_t SatClkCorrInfo_1_0_t
type SatHPCA_t SatHPCA_1_0_t
type CorrSample_t CorrSample_1_0_t
type galIntegrity_t gaIntegrity_1_0_t
type galSARRLM_t gaSARRLM_1_0_t
type TimeOffset_t TimeOffset_1_1_t
type ExtSensorMeasSB_t ExtSensorMeasSB_1_0_t
