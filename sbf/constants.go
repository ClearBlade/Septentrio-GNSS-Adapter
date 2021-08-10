package sbf

//Payload type declarations

/**
 * Declarations and definitions of SBF (Septentrio Binary Format) blocks.
 *
 * Adapted from sbfdef.h in the sdk provided in the RxTools suite
 */

const RTCM16MSG_LENGTH = 90
const CMR2SHORTID_LENGTH = 8
const CMR2COGO_LENGTH = 16
const CMR2LONGID_LENGTH = 51
const EXTSENSORSERIALNBR_LENGTH = 20
const EXTSENSORVERSION_LENGTH = 20

/* Using maximum value of NR_OF_LOGICALCHANNELS for SBF definitions */
const NR_OF_LOGICALCHANNELS = 80
const MAX_NB_INMARSATCHANNELS = 1

/* Using maximum value of MAX_NR_OF_SIGNALS_PER_SATELLITE for SBF definitions */
const MAX_NR_OF_SIGNALS_PER_SATELLITE = 7

/* Using maximum value of NR_OF_ANTENNAS for SBF definitions */
const NR_OF_ANTENNAS = 3

/* Using maximum value of MAX_NR_SATSYSTEMS for SBF definitions */
const MAX_NR_SATSYSTEMS = 5

/* Max Nr of Subblocks for different SBF messages */
const MAXSB_MEASEPOCH_T1 = NR_OF_LOGICALCHANNELS + MAX_NB_INMARSATCHANNELS
const MAXSB_MEASEPOCH_T2 = MAXSB_MEASEPOCH_T1 * ((MAX_NR_OF_SIGNALS_PER_SATELLITE * NR_OF_ANTENNAS) - 1)
const MAXSB_NBRANTENNA = 4
const MAXSB_CHANNELSATINFO = NR_OF_LOGICALCHANNELS + MAX_NB_INMARSATCHANNELS
const MAXSB_CHANNELSTATEINFO = MAXSB_CHANNELSATINFO * MAXSB_NBRANTENNA
const MAXSB_SATOBS = NR_OF_LOGICALCHANNELS * NR_OF_ANTENNAS * MAX_NR_OF_SIGNALS_PER_SATELLITE
const MAXSB_SATRES = 3
const MAXSB_CONNDESCR = 16
const MAXSB_DATATYPES = 16

const I8_NOTVALID uint8 = 0x80
const UI8_NOTVALID uint8 = 0xFF
const I16_NOTVALID uint16 = 0x8000
const U16_NOTVALID uint16 = 0xFFFF
const I32_NOTVALID uint32 = 0x80000000
const U32_NOTVALID uint32 = 0xFFFFFFFF
const I64_NOTVALID uint64 = 0x8000000000000000
const U64_NOTVALID uint64 = 0xFFFFFFFFFFFFFFFF
const F32_NOTVALID = -2e10
const F64_NOTVALID = -2e10

/*==SBF-IDs==============================================================*/
/* minor version are indicated in first 3 bits */
const SBF_ID_REVISION_0 = 0x0000
const SBF_ID_REVISION_1 = 0x2000
const SBF_ID_REVISION_2 = 0x4000
const SBF_ID_REVISION_3 = 0x6000
const SBF_ID_REVISION_4 = 0x8000
const SBF_ID_REVISION_5 = 0xA000
const SBF_ID_REVISION_6 = 0xC000
const SBF_ID_REVISION_7 = 0xE000

func SBF_ID_TO_NUMBER(id uint16) uint16 { return id & 0x1fff }
func SBF_ID_TO_REV(id uint16) uint16    { return id >> 13 }

const MIN_SBFSIZE = 8     /*!< minimum size of a SBF block in bytes */
const MAX_SBFSIZE = 65535 /*!< maximum size of a SBF block in bytes */

/*==CONNECTION DESCRIPTOR definition=========================================*/
const CD_TYPE_COM = 0x00
const CD_TYPE_USB = 0x20
const CD_TYPE_IP = 0x40
const CD_TYPE_DSK = 0x60
const CD_TYPE_NTR = 0x80
const CD_TYPE_IPS = 0xA0
const CD_TYPE_RESERVED3 = 0xC0
const CD_TYPE_RESERVED4 = 0xE0

/*==DATA TYPE definition=====================================================*/
const DATA_TYPE_UNKNOWN = 0x00
const DATA_TYPE_SEPT = 0x20
const DATA_TYPE_NMEA = 0x40
const DATA_TYPE_DIFFCORR = 0x60
const DATA_TYPE_IMU = 0x80
const DATA_TYPE_RESERVED1 = 0xA0 /* Proper naming to be found */
const DATA_TYPE_RESERVED2 = 0xC0
const DATA_TYPE_RESERVED3 = 0xE0

/*==Measurement Blocks=======================================================*/

/*--GenMeasEpoch_1_0_t : ----------------------------------------------------*/
// const SBF_GENMEASEPOCH_1_0_DATA_LENGTH = (MAXSB_MEASEPOCH_T1*binary.Size(MeasEpochChannelType1_1_0_t) +
// 	MAXSB_MEASEPOCH_T2*sizeof(MeasEpochChannelType2_1_0_t))
const SBF_GENMEASEPOCH_1_0_DATA_LENGTH = (MAXSB_MEASEPOCH_T1*20 + MAXSB_MEASEPOCH_T2*12)

/*--MeasEpoch_2_0_t : -------------------------------------------------------*/
const SBF_MEASEPOCH_2_0_DATA_LENGTH = SBF_GENMEASEPOCH_1_0_DATA_LENGTH

/*--MeasEpoch_2_1_t : -------------------------------------------------------*/
const SBF_MEASEPOCH_2_1_DATA_LENGTH = SBF_GENMEASEPOCH_1_0_DATA_LENGTH

const SBF_MEASEPOCH_2_COMMONFLAG_MULTIPATHMITIGATION uint8 = 1 << 0 // = COMMONFLAG_MULTIPATHMITIGATION (same as in tracker)
const SBF_MEASEPOCH_2_COMMONFLAG_ATLEASTONESMOOTHING uint8 = 1 << 1 // = COMMONFLAG_ATLEASTONESMOOTHING (same as in tracker)
const SBF_MEASEPOCH_2_COMMONFLAG_CARRIERPHASEALIGN uint8 = 1 << 2
const SBF_MEASEPOCH_2_COMMONFLAG_CLOCKSTEERINGACTIVE uint8 = 1 << 3 // = COMMONFLAG_CLOCKSTEERINGACTIVE (same as in tracker)
const SBF_MEASEPOCH_2_COMMONFLAG_HIGHDYNAMICSMODE uint8 = 1 << 5

/*--MeasExtra_1_0_t : -------------------------------------------------------*/
const SBF_MEASEXTRA_1_0_CHANNELSUB_LENGTH = MAXSB_MEASEPOCH_T1 + MAXSB_MEASEPOCH_T2

/*--MeasExtra_1_1_t : -------------------------------------------------------*/
const SBF_MEASEXTRA_1_1_CHANNELSUB_LENGTH = MAXSB_MEASEPOCH_T1 + MAXSB_MEASEPOCH_T2

/*--MeasExtra_1_2_t : -------------------------------------------------------*/
const SBF_MEASEXTRA_1_2_CHANNELSUB_LENGTH = MAXSB_MEASEPOCH_T1 + MAXSB_MEASEPOCH_T2

/*--MeasExtra_1_3_t : -------------------------------------------------------*/
const SBF_MEASEXTRA_1_3_CHANNELSUB_LENGTH = MAXSB_MEASEPOCH_T1 + MAXSB_MEASEPOCH_T2

/*--IQCorr_1_0_t : ----------------------------------------------------------*/
const SBF_IQCORR_1_0_CHANNELSUB_LENGTH = MAXSB_MEASEPOCH_T1 + MAXSB_MEASEPOCH_T2

/*--IQCorr_1_1_t : ----------------------------------------------------------*/
const SBF_IQCORR_1_1_CHANNELSUB_LENGTH = SBF_IQCORR_1_0_CHANNELSUB_LENGTH

/*--SQMSamples_1_0_t : ------------------------------------------------------*/
const SBF_SQMSAMPLES_1_0_DATA_LENGTH = 8172

/*==Galileo Decoded Message Blocks===========================================*/
const SBF_GAL_SOURCE_INAV = 2
const SBF_GAL_SOURCE_FNAV = 16

/*==Position, Velocity and Time Blocks=======================================*/
const SBF_PVTERR_NONE = 0            /*!< no error */
const SBF_PVTERR_NOTENOUGHMEAS = 1   /*!< not enough measurements */
const SBF_PVTERR_NOTENOUGHEPH = 2    /*!< not enough ephemerides */
const SBF_PVTERR_DOPTOOHIGH = 3      /*!< DOP too large (larger than 15) */
const SBF_PVTERR_ETETOOLARGE = 4     /*!< sum of sqrd residuals too big */
const SBF_PVTERR_NOCONVERGENCE = 5   /*!< no convergence  */
const SBF_PVTERR_TOOMANYOUTLIERS = 6 /*!< not enough measurements after
outlier rejection  */
const SBF_PVTERR_EXCEEDNATOLIMITS = 7 /*!< position output prohibited
due to export laws  */
const SBF_PVTERR_NOTENOUGHDIFFCORR = 8 /*!< not enough differential
corrections available  */
const SBF_PVTERR_NOBASEAVL = 9 /*!< base station coordinates
unavailable */
const SBF_PVTERR_AMBIGUITIESNOTFIXED = 10 /*!< Error flag only used when
RTKfixed only is requested
and receiver is still in RTK
float32 mode */
const SBF_PVTERR_NOTRANSFOPARAMS = 17 /*!< Datum transformation parameters
unknown */
const SBF_PVTERR_INSNOTREQUESTEDBYUSER = 20 /*!< Error flag only used for integrated
loosely-coupled GNSS/IMU PVT. This
error is provided when the user has
selected 'ExtSensorIntegration off'
in the SPM command. */
const SBF_PVTERR_NOTENOUGHEXTSENSORMEAS = 21 /*!< Not enough external sensor measurements
available for integrated PVA */
const SBF_PVTERR_CALIBRATIONNOTREADY = 22 /*!< External sensor calibration (automatic
orientation and lever-arm detection)
is still ongoing */
const SBF_PVTERR_ALIGNMENTNOTREADY = 23 /*!< External sensor static alignment (coarse
alignment) is still ongoing */
const SBF_PVTERR_WAITINGFORGNSSPVT = 24         /*!< Waiting for GNSS PVT solution */
const SBF_PVTERR_FINETIMENOTREACHED = 27        /*!< FINETIME mode not reached yet */
const SBF_PVTERR_INMOTIONALIGNMENTNOTREADY = 28 /*!< External sensor in motion alignment is
still ongoing */
const SBF_PVTERR_WAITINGFORGNSSHEAD = 29 /*!< Waiting for GNSS heading
solution */
const SBF_PVTERR_WAITINGFORPPSSYNC = 30 /*!< External sensor waiting to
sync with PPS from
receiver */
const SBF_PVTERR_STDLIMITEXCEEDED = 31 /*!< Error flag if standard deviation of
position/attitude solution exceeds
user limit */
const SBF_PVTERR_UNSUPPORTEDSETTINGSINS = 32 /*!< user provided settings are not supported
in INS */

/* Defines for GNSS PVT mode */
const MODE_NO_PVT_AVAILABLE = 0x00
const MODE_STAND_ALONE_PVT = 0x01
const MODE_DIFFERENTIAL_PVT = 0x02
const MODE_FIXED_LOCATION = 0x03
const MODE_RTK_FIXED_AMBIGUITIES = 0x04
const MODE_RTK_FLOAT_AMBIGUITIES = 0x05
const MODE_SBAS_AIDED_PVT = 0x06
const MODE_MOVBASERTK_FIXED_AMBIGUITIES = 0x07
const MODE_MOVBASERTK_FLOAT_AMBIGUITIES = 0x08
const MODE_PPP_FIXED_AMBIGUITIES = 0x09
const MODE_PPP_FLOAT_AMBIGUITIES = 0x0A
const MODE_RTK_FLOAT_WL_FIXED_AMBIGUITIES = 0x0B
const MODE_LOCATA_PVT = 0x0C

const MODE_STATICAUTO_LOOKING = 0x40
const MODE_2D_PVT = 0x80

/* Defines for GNSS/INS Integrated PVT mode */
const MODEINTPVA_NO_PVA_AVAILABLE = 0x00
const MODEINTPVA_EXTRAP_EXTSENSORS = 0x01
const MODEINTPVA_LOOSELY_INTEGRATED = 0x02
const MODEINTPVA_TIGHTLY_INTEGRATED = 0x03
const MODEINTPVA_GNSSONLY_EXTRAPOLATED = 0x04

/* Defines for Attitude mode */
const SBF_ATT_MODE_NOT_AVAILABLE = 0
const SBF_ATT_MODE_HEADING_PITCH_FLOAT = 1
const SBF_ATT_MODE_HEADING_PITCH_FIXED = 2
const SBF_ATT_MODE_HEADING_PITCH_ROLL_FLOAT = 3
const SBF_ATT_MODE_HEADING_PITCH_ROLL_FIXED = 4

/*--PVTSatCartesian_1_0_t : -------------------------------------------------*/
//const SBF_PVTSATCARTESIAN_1_0_SATPOS_LENGTH = NR_OF_LOGICALCHANNELS

/*--PVTSatCartesian_1_1_t : -------------------------------------------------*/
//const SBF_PVTSATCARTESIAN_1_1_SATPOS_LENGTH = SBF_PVTSATCARTESIAN_1_0_SATPOS_LENGTH

const SATPOS_IONOMODEL_NOT_APPLICABLE = 0
const SATPOS_IONOMODEL_KLOBUCHAR = 1
const SATPOS_IONOMODEL_DO229 = 2
const SATPOS_IONOMODEL_NEQUICK = 3
const SATPOS_IONOMODEL_MEASURED = 4
const SATPOS_IONOMODEL_ESTIMATED = 5
const SATPOS_IONOMODEL_KLOBUCHARBEIDOU = 6

/*--PVTResiduals_1_0_t : ----------------------------------------------------*/
//const SBF_PVTRESIDUALS_1_0_SATRESIDUAL_LENGTH = NR_OF_LOGICALCHANNELS

/*--PVTResiduals_2_0_t : ----------------------------------------------------*/
//#define SBF_PVTRESIDUALS_2_0_DATA_LENGTH  (MAX_SBFSIZE - sizeof(BlockHeader_t) - 12)
const SBF_PVTRESIDUALS_2_0_DATA_LENGTH = MAX_SBFSIZE - 8 - 12

/*--PVTResiduals_2_1_t : ----------------------------------------------------*/
const SBF_PVTRESIDUALS_2_1_DATA_LENGTH = SBF_PVTRESIDUALS_2_0_DATA_LENGTH

/*--RAIMStatistics_1_0_t : --------------------------------------------------*/
const RAIM_INTEGRITY_OK = 0
const RAIM_INTEGRITY_FAIL = 1
const RAIM_INTEGRITY_NOTAVL = 2

const RAIM_UI16SCALE = 50000.0
const RAIM_MDB_SCALE = 1000.0

//const SBF_RAIMSTATISTICS_1_0_RAIMCHANNEL_LENGTH = NR_OF_LOGICALCHANNELS

/*--GEOCorrections_1_0_t : --------------------------------------------------*/
const SBF_GEOCORRECTIONS_1_0_SATCORR_LENGTH = NR_OF_LOGICALCHANNELS

/*-- PosCart_1_0_t : --------------------------------------------------------*/
const POSCART_MISC_MASK_BASELINE_TO_ARP = 0x01
const POSCART_MISC_MASK_PCV_COMPENSATED = 0x02
const POSCART_MISC_MASK_SIGIL_ENABLED = 0x04
const POSCART_MISC_MASK_RTCMV_USED = 0x08
const POSCART_MISC_MASK_MARKER_IS_ARP = 0x40
const POSCART_MISC_MASK_MARKER_IS_NOT_ARP = 0x80

/*==INS/GNSS Integrated Blocks===============================================*/

/*--IntPVCart_1_0_t : -------------------------------------------------------*/
const INTPVA_INFO_HEADING_AMBIGUITY_VALIDATED = 0x0800 /* Bit 11 */
const INTPVA_INFO_MOTIONCONSTRAINTUSED = 0x1000        /* Bit 12 */
const INTPVA_INFO_GNSSPOSUSED = 0x2000                 /* Bit 13 */
const INTPVA_INFO_GNSSVELUSED = 0x4000                 /* Bit 14 */
const INTPVA_INFO_GNSSATTUSED = 0x8000                 /* Bit 15 */

/*--INSNavCart_1_0_t : ------------------------------------------------------*/
const INSNAV_INFO_OUTPUT_LOCATION_MAINANT = 0x0000     /* Bit 3-5: 0 */
const INSNAV_INFO_OUTPUT_LOCATION_POI1 = 0x0008        /* Bit 3-5: 1 */
const INSNAV_INFO_HEADING_AMBIGUITY_VALIDATED = 0x0040 /* Bit 6 */
const INSNAV_INFO_MOTIONCONSTRAINTUSED = 0x0080        /* Bit 7 */

const INSNAV_SBLIST_POSITION_STTDEV = 0x0001 /* Bit 0 */
const INSNAV_SBLIST_ATTITDE = 0x0002         /* Bit 1 */
const INSNAV_SBLIST_ATTITDE_STTDEV = 0x0004  /* Bit 2 */
const INSNAV_SBLIST_VELOCITY = 0x0008        /* Bit 3 */
const INSNAV_SBLIST_VELOCITY_STTDEV = 0x0010 /* Bit 4 */
const INSNAV_SBLIST_POSITION_COV = 0x0020    /* Bit 5 */
const INSNAV_SBLIST_ATTITDE_COV = 0x0040     /* Bit 6 */
const INSNAV_SBLIST_VELOCITY_COV = 0x0080    /* Bit 7 */

/*==GNSS Attitude Blocks=====================================================*/

/*--AttEuler_1_0_t : --------------------------------------------------------*/
const SBF_ATTERR_NONE = 0x0                    /*!< no error */
const SBF_ATTERR_NOTENOUGHMEAS = 0x1           /*!< not enough measurements */
const SBF_ATTERR_ANTENNASALIGNED = 0x2         /*!< antennas are aligned */
const SBF_ATTERR_INCONSISTENCYWITHANTPOS = 0x3 /*!< inconsistency with manual
antenna position information
*/
const SBF_ATTERR_NO_ATTITUDE_REQUESTED = 0x80 /*!< attitude not requested by
user */
/* next define is kept for backwards compatibility, better use
SBF_ATTERR_NO_ATTITUDE_REQUESTED */
const ATTEULER_ERROR_NO_ATTITUDE_REQUESTED = 0x80 /*!< attitude not requested by
user */

/*--AuxAntPositions_1_0_t : -------------------------------------------------*/
//const ATT_MAX_NBR_ANT = SBF_AUXANTPOSITIONS_1_0_AUXANTPOSITION_LENGTH

/*==Differential Correction Blocks===========================================*/

/*--DiffCorrIn_1_0_t : ------------------------------------------------------*/
const RTCM2_MAX_NBR_WORDS = 33
const CMR_MAX_NBR_WORDS = (((4 + 6 + (31 * 15) + 2) / 4) + 1) //120.25
const RTCM3_MAX_NBR_WORDS = 258

//TODO - Figure out how to do this in GO. I have no idea what they were trying to accomplish with this!

//  #if (RTCM2_MAX_NBR_WORDS  >=  CMR_MAX_NBR_WORDS) && (RTCM2_MAX_NB_WORDS >= RTCM3_MAX_NBR_WORDS)
//  #define DIFFCORR_FRAME_MAX_LEN  RTCM2_MAX_NBR_WORDS
//  #elif (CMR_MAX_NBR_WORDS >= RTCM2_MAX_NBR_WORDS) && (CMR_MAX_NBR_WORDS >= RTCM3_MAX_NBR_WORDS)
//  #define DIFFCORR_FRAME_MAX_LEN  CMR_MAX_NBR_WORDS
//  #else
//  #define DIFFCORR_FRAME_MAX_NBR_WORDS  RTCM3_MAX_NBR_WORDS
const DIFFCORR_FRAME_MAX_NBR_WORDS = RTCM3_MAX_NBR_WORDS

//  #endif

const SBF_RTCM = 0x00
const SBF_CMR = 0x01 /* used in Reserved[0] from CMR release onwards */
const SBF_RTCM3 = 0x02
const SBF_RTCMV = 0x03
const SBF_SPARTN = 0x04
const SBF_RTCM3_ENC = 0x05

//const SBF_DIFFCORRIN_1_0_FRAME_LENGTH = DIFFCORR_FRAME_MAX_NBR_WORDS * 4

/*--RawDataIn_1_0_t : -------------------------------------------------------*/

const SBF_RAWDATAIN_UNKNOWN = 0x00
const SBF_RAWDATAIN_LBMP_MAIN = 0x01
const SBF_RAWDATAIN_LBMP_AUX1 = 0x02

/*--BaseStation_1_0_t : -----------------------------------------------------*/
const BASESTATION_BASETYPE_FIXED = 0
const BASESTATION_BASETYPE_MOVING = 1

const BASESTATION_SOURCE_RTCM2X_MSG3 = 0
const BASESTATION_SOURCE_RTCM2X_MSG24 = 2
const BASESTATION_SOURCE_CMR2X_MSG1 = 4
const BASESTATION_SOURCE_RTCM3X_MSG1005 = 8
const BASESTATION_SOURCE_RTCMV = 9
const BASESTATION_SOURCE_CMRP_MSG2 = 10

/*==External Sensor Blocks===================================================*/

/*--ExtSensor macros --------------------------------------------------------*/
const EXTSENSORSBF_SENSORMODEL_MMQ50 = 0
const EXTSENSORSBF_SENSORMODEL_MTI = 1
const EXTSENSORSBF_SENSORMODEL_ELLIPSE = 2
const EXTSENSORSBF_SENSORMODEL_MTI10 = 3
const EXTSENSORSBF_SENSORMODEL_SIMSENSOR = 4
const EXTSENSORSBF_SENSORMODEL_ELLIPSE2 = 5
const EXTSENSORSBF_SENSORMODEL_EKINOX2 = 6
const EXTSENSORSBF_SENSORMODEL_VN100 = 7
const EXTSENSORSBF_SENSORMODEL_HGUIDE = 8
const EXTSENSORSBF_SENSORMODEL_RIVIANIMU = 9
const EXTSENSORSBF_SENSORMODEL_ADIS = 10
const EXTSENSORSBF_SENSORMODEL_ZUPT = 20
const EXTSENSORSBF_SENSORMODEL_VSM1 = 21
const EXTSENSORSBF_SENSORMODEL_VSM2 = 22
const EXTSENSORSBF_SENSORMODEL_VSM3 = 23
const EXTSENSORSBF_SENSORMODEL_VSM4 = 24

const EXTSENSORSBF_MEASTYPE_ACC = 0
const EXTSENSORSBF_MEASTYPE_RATE = 1
const EXTSENSORSBF_MEASTYPE_MAGNETICFIELD = 2
const EXTSENSORSBF_MEASTYPE_INFO = 3
const EXTSENSORSBF_MEASTYPE_VELOCITY = 4
const EXTSENSORSBF_MEASTYPE_ZUPT = 20

/*--ExtSensorStatus_1_0_t : -------------------------------------------------*/
const EXTSENSORSTATUS_STATUSTYPE_CONFIG = 0
const EXTSENSORSTATUS_STATUSTYPE_STATUS = 1

/*--ExtSensorSetup_1_0_t : --------------------------------------------------*/
const EXTSENSORSETUP_MEASTYPE_ACCELERATIONS = 0x0001
const EXTSENSORSETUP_MEASTYPE_ANGULAR_RATES = 0x0002
const EXTSENSORSETUP_MEASTYPE_MAGNETICFIELD = 0x0004
const EXTSENSORSETUP_MEASTYPE_PRESSURE = 0x0008
const EXTSENSORSETUP_MEASTYPE_TEMPERATURE = 0x0010

/*--ExtSensorSetup_1_1_t : --------------------------------------------------*/
const EXTSENSORSETUP_LEVERARMSOURCE_NVRAM = 0x00
const EXTSENSORSETUP_LEVERARMSOURCE_MANUAL = 0x01
const EXTSENSORSETUP_LEVERARMSOURCE_CALIBRATION = 0x02

/*==Status Blocks============================================================*/

/*--TrackingStatus_1_0_t : --------------------------------------------------*/
//const SBF_TRACKINGSTATUS_1_0_CHANNELDATA_LENGTH = NR_OF_LOGICALCHANNELS

/*--ChannelStatus_1_0_t : ---------------------------------------------------*/
// #define SBF_CHANNELSTATUS_1_0_DATA_LENGTH MAXSB_CHANNELSATINFO*sizeof(ChannelSatInfo_1_0_t) +
// 	MAXSB_CHANNELSTATEINFO*sizeof(ChannelStateInfo_1_0_t)

const SBF_CHANNELSTATUS_1_0_DATA_LENGTH = MAXSB_CHANNELSATINFO*12 + MAXSB_CHANNELSTATEINFO*8

/*--ReceiverStatus_2_0_t : --------------------------------------------------*/
const FRONTENDID_GPSGALL1 = 0
const FRONTENDID_GLOL1 = 1
const FRONTENDID_GALE6 = 2
const FRONTENDID_GPSL2 = 3
const FRONTENDID_GLOL2 = 4
const FRONTENDID_GPSGALL5E5a = 5
const FRONTENDID_GALE5b = 6
const FRONTENDID_GALE5ab = 7
const FRONTENDID_GPSGLOGALL1 = 8
const FRONTENDID_GPSGLOL2 = 9
const FRONTENDID_LBAND = 10
const FRONTENDID_CMPB1 = 11
const FRONTENDID_CMPB3 = 12
const FRONTENDID_SBAND = 13
const FRONTENDID_UNUSED = 31

/*--InputLink_1_0_t : -------------------------------------------------------*/
//const SBF_INPUTLINK_1_0_DATA_LENGTH = MAXSB_CONNDESCR * sizeof(InputStatsSub_1_0_t)
const SBF_INPUTLINK_1_0_DATA_LENGTH = MAXSB_CONNDESCR * 20

/*--OutputLink_1_0_t : ------------------------------------------------------*/
//const SBF_OUTPUTLINK_1_0_DATA_LENGTH = (MAXSB_CONNDESCR * (sizeof(OutputStatsSub_1_0_t) + (MAXSB_DATATYPES * sizeof(OutputTypeSub_1_0_t))))
const SBF_OUTPUTLINK_1_0_DATA_LENGTH = (MAXSB_CONNDESCR * (12 + (MAXSB_DATATYPES * 4)))

/*--OutputLink_1_1_t : ------------------------------------------------------*/
//const SBF_OUTPUTLINK_1_1_DATA_LENGTH = (MAXSB_CONNDESCR * (sizeof(OutputStatsSub_1_1_t) + (MAXSB_DATATYPES * sizeof(OutputTypeSub_1_0_t))))
const SBF_OUTPUTLINK_1_1_DATA_LENGTH = (MAXSB_CONNDESCR * (16 + (MAXSB_DATATYPES * 4)))

/*--QualityInd_1_0_t : ------------------------------------------------------*/
const QUALITYIND_OVERALL = 0
const QUALITYIND_MAINSIGNAL = 1
const QUALITYIND_AUX1SIGNAL = 2
const QUALITYIND_MAINANTCABLING = 11
const QUALITYIND_AUX1ANTCABLING = 12
const QUALITYIND_CPUHEADROOM = 21
const QUALITYIND_CLOCKSTABILITY = 25
const QUALITYIND_BASEMEAS = 30
const QUALITYIND_RTKPOSTPROCESS = 31

/*==TUR Specific Blocks======================================================*/

/*--TURPVTSatCorrections_1_0_t : ---------------------------------------------*/
// const SBF_TURPVTSATCORRECTIONS_1_0_DATA_LENGTH = (NR_OF_LOGICALCHANNELS * (sizeof(SatCorrInfo_1_0_t) +
// 	NR_OF_ANTENNAS*MAX_NR_OF_SIGNALS_PER_SATELLITE*sizeof(SatClkCorrInfo_1_0_t)))
const SBF_TURPVTSATCORRECTIONS_1_0_DATA_LENGTH = (NR_OF_LOGICALCHANNELS * (20 + NR_OF_ANTENNAS*MAX_NR_OF_SIGNALS_PER_SATELLITE*12))

/*--TURPVTSatCorrections_1_1_t : ---------------------------------------------*/
// const SBF_TURPVTSATCORRECTIONS_1_1_DATA_LENGTH = (NR_OF_LOGICALCHANNELS * (sizeof(SatCorrInfo_1_1_t) +
// 	NR_OF_ANTENNAS*MAX_NR_OF_SIGNALS_PER_SATELLITE*sizeof(SatClkCorrInfo_1_1_t)))
const SBF_TURPVTSATCORRECTIONS_1_1_DATA_LENGTH = (NR_OF_LOGICALCHANNELS * (28 + NR_OF_ANTENNAS*MAX_NR_OF_SIGNALS_PER_SATELLITE*12))

/*--TURHPCAInfo_1_0_t : -----------------------------------------------------*/
const SBF_TURHPCAINFO_1_0_HPCA_LENGTH = NR_OF_LOGICALCHANNELS

/*--SysTimeOffset_1_0_t : ---------------------------------------------------*/
//const SBF_SYSTIMEOFFSET_1_0_TIMEOFFSET_LENGTH = 2*MAX_NR_SATSYSTEMS - 2

/*--Meas3Ranges_1_0_t : ------------------------------------------------*/
const SBF_MEAS3RANGES_1_0_DATA_LENGTH = 32768

/* maximum number of signal types per satellite in a Meas3Ranges SBF block */
const MEAS3_SIG_MAX = 16

/* maximum number of satellites per constellation in a Meas3Ranges SBF block */
const MEAS3_SAT_MAX = 64

/* Index of constellations in a Meas3Ranges SBF block */
const (
	MEAS3_SYS_GPS  = iota
	MEAS3_SYS_GLO  = iota
	MEAS3_SYS_GAL  = iota
	MEAS3_SYS_BDS  = iota
	MEAS3_SYS_SBAS = iota
	MEAS3_SYS_QZS  = iota
	MEAS3_SYS_IRN  = iota
	MEAS3_SYS_MAX  = iota
)

/*==SBF-IDs==============================================================*/
/*! SBF ID, which uniquely identifies the SBF block type and its contents.
 It is a bit field with the following definition:
 - bits 0-12: block number;
 - bits 13-15: block revision number, starting from 0 at the initial block definition,
	 and incrementing each time backwards-compatible changes are performed to the block */
const (
	/* Measurement Blocks */
	sbfnr_GenMeasEpoch_1    = 5944
	sbfid_GenMeasEpoch_1_0  = 5944 | 0x0
	sbfnr_MeasEpoch_2       = 4027
	sbfid_MeasEpoch_2_0     = 4027 | 0x0
	sbfid_MeasEpoch_2_1     = 4027 | 0x2000
	sbfnr_MeasExtra_1       = 4000
	sbfid_MeasExtra_1_0     = 4000 | 0x0
	sbfid_MeasExtra_1_1     = 4000 | 0x2000
	sbfid_MeasExtra_1_2     = 4000 | 0x4000
	sbfid_MeasExtra_1_3     = 4000 | 0x6000
	sbfnr_MeasFullRange_1   = 4098
	sbfid_MeasFullRange_1_0 = 4098 | 0x0
	sbfid_MeasFullRange_1_1 = 4098 | 0x2000
	sbfnr_Meas3Ranges_1     = 4109
	sbfid_Meas3Ranges_1_0   = 4109 | 0x0
	sbfnr_Meas3CN0HiRes_1   = 4110
	sbfid_Meas3CN0HiRes_1_0 = 4110 | 0x0
	sbfnr_Meas3Doppler_1    = 4111
	sbfid_Meas3Doppler_1_0  = 4111 | 0x0
	sbfnr_Meas3PP_1         = 4112
	sbfid_Meas3PP_1_0       = 4112 | 0x0
	sbfnr_Meas3MP_1         = 4113
	sbfid_Meas3MP_1_0       = 4113 | 0x0
	sbfnr_IQCorr_1          = 4046
	sbfid_IQCorr_1_0        = 4046 | 0x0
	sbfid_IQCorr_1_1        = 4046 | 0x2000
	sbfnr_ISMR_1            = 4086
	sbfid_ISMR_1_0          = 4086 | 0x0
	sbfnr_SQMSamples_1      = 4087
	sbfid_SQMSamples_1_0    = 4087 | 0x0
	sbfnr_EndOfMeas_1       = 5922
	sbfid_EndOfMeas_1_0     = 5922 | 0x0
	/* Navigation Page Blocks */
	sbfnr_GPSRaw_1        = 5895
	sbfid_GPSRaw_1_0      = 5895 | 0x0
	sbfnr_CNAVRaw_1       = 5947
	sbfid_CNAVRaw_1_0     = 5947 | 0x0
	sbfnr_GEORaw_1        = 5898
	sbfid_GEORaw_1_0      = 5898 | 0x0
	sbfnr_GPSRawCA_1      = 4017
	sbfid_GPSRawCA_1_0    = 4017 | 0x0
	sbfnr_GPSRawL2C_1     = 4018
	sbfid_GPSRawL2C_1_0   = 4018 | 0x0
	sbfnr_GPSRawL5_1      = 4019
	sbfid_GPSRawL5_1_0    = 4019 | 0x0
	sbfnr_GPSRawL1C_1     = 4221
	sbfid_GPSRawL1C_1_0   = 4221 | 0x0
	sbfnr_GLORawCA_1      = 4026
	sbfid_GLORawCA_1_0    = 4026 | 0x0
	sbfnr_GALRawFNAV_1    = 4022
	sbfid_GALRawFNAV_1_0  = 4022 | 0x0
	sbfnr_GALRawINAV_1    = 4023
	sbfid_GALRawINAV_1_0  = 4023 | 0x0
	sbfnr_GALRawCNAV_1    = 4024
	sbfid_GALRawCNAV_1_0  = 4024 | 0x0
	sbfnr_GALRawGNAV_1    = 4025
	sbfid_GALRawGNAV_1_0  = 4025 | 0x0
	sbfnr_GALRawGNAVe_1   = 4029
	sbfid_GALRawGNAVe_1_0 = 4029 | 0x0
	sbfnr_GEORawL1_1      = 4020
	sbfid_GEORawL1_1_0    = 4020 | 0x0
	sbfnr_GEORawL5_1      = 4021
	sbfid_GEORawL5_1_0    = 4021 | 0x0
	sbfnr_BDSRaw_1        = 4047
	sbfid_BDSRaw_1_0      = 4047 | 0x0
	sbfnr_BDSRawB1C_1     = 4218
	sbfid_BDSRawB1C_1_0   = 4218 | 0x0
	sbfnr_BDSRawB2a_1     = 4219
	sbfid_BDSRawB2a_1_0   = 4219 | 0x0
	sbfnr_BDSRawB2b_1     = 4242
	sbfid_BDSRawB2b_1_0   = 4242 | 0x0
	sbfnr_QZSRawL1CA_1    = 4066
	sbfid_QZSRawL1CA_1_0  = 4066 | 0x0
	sbfnr_QZSRawL2C_1     = 4067
	sbfid_QZSRawL2C_1_0   = 4067 | 0x0
	sbfnr_QZSRawL5_1      = 4068
	sbfid_QZSRawL5_1_0    = 4068 | 0x0
	sbfnr_QZSRawL6_1      = 4069
	sbfid_QZSRawL6_1_0    = 4069 | 0x0
	sbfnr_QZSRawL1C_1     = 4227
	sbfid_QZSRawL1C_1_0   = 4227 | 0x0
	sbfnr_QZSRawL1S_1     = 4228
	sbfid_QZSRawL1S_1_0   = 4228 | 0x0
	sbfnr_NAVICRaw_1      = 4093
	sbfid_NAVICRaw_1_0    = 4093 | 0x0
	sbfnr_GNSSNavBits_1   = 4088
	sbfid_GNSSNavBits_1_0 = 4088 | 0x0
	sbfnr_GNSSSymbols_1   = 4099
	sbfid_GNSSSymbols_1_0 = 4099 | 0x0
	/* GPS Decoded Message Blocks */
	sbfnr_GPSNav_1    = 5891
	sbfid_GPSNav_1_0  = 5891 | 0x0
	sbfnr_GPSAlm_1    = 5892
	sbfid_GPSAlm_1_0  = 5892 | 0x0
	sbfnr_GPSIon_1    = 5893
	sbfid_GPSIon_1_0  = 5893 | 0x0
	sbfnr_GPSUtc_1    = 5894
	sbfid_GPSUtc_1_0  = 5894 | 0x0
	sbfnr_GPSCNav_1   = 4042
	sbfid_GPSCNav_1_0 = 4042 | 0x0
	/* GLONASS Decoded Message Blocks */
	sbfnr_GLONav_1    = 4004
	sbfid_GLONav_1_0  = 4004 | 0x0
	sbfnr_GLOAlm_1    = 4005
	sbfid_GLOAlm_1_0  = 4005 | 0x0
	sbfnr_GLOTime_1   = 4036
	sbfid_GLOTime_1_0 = 4036 | 0x0
	/* Galileo Decoded Message Blocks */
	sbfnr_GALNav_1      = 4002
	sbfid_GALNav_1_0    = 4002 | 0x0
	sbfnr_GALAlm_1      = 4003
	sbfid_GALAlm_1_0    = 4003 | 0x0
	sbfnr_GALIon_1      = 4030
	sbfid_GALIon_1_0    = 4030 | 0x0
	sbfnr_GALUtc_1      = 4031
	sbfid_GALUtc_1_0    = 4031 | 0x0
	sbfnr_GALGstGps_1   = 4032
	sbfid_GALGstGps_1_0 = 4032 | 0x0
	sbfnr_GALSARRLM_1   = 4034
	sbfid_GALSARRLM_1_0 = 4034 | 0x0
	/* BeiDou Decoded Message Blocks */
	sbfnr_BDSNav_1   = 4081
	sbfid_BDSNav_1_0 = 4081 | 0x0
	sbfnr_BDSAlm_1   = 4119
	sbfid_BDSAlm_1_0 = 4119 | 0x0
	sbfnr_BDSIon_1   = 4120
	sbfid_BDSIon_1_0 = 4120 | 0x0
	sbfnr_BDSUtc_1   = 4121
	sbfid_BDSUtc_1_0 = 4121 | 0x0
	/* QZSS Decoded Message Blocks */
	sbfnr_QZSNav_1   = 4095
	sbfid_QZSNav_1_0 = 4095 | 0x0
	sbfnr_QZSAlm_1   = 4116
	sbfid_QZSAlm_1_0 = 4116 | 0x0
	/* SBAS L1 Decoded Message Blocks */
	sbfnr_GEOMT00_1                = 5925
	sbfid_GEOMT00_1_0              = 5925 | 0x0
	sbfnr_GEOPRNMask_1             = 5926
	sbfid_GEOPRNMask_1_0           = 5926 | 0x0
	sbfnr_GEOFastCorr_1            = 5927
	sbfid_GEOFastCorr_1_0          = 5927 | 0x0
	sbfnr_GEOIntegrity_1           = 5928
	sbfid_GEOIntegrity_1_0         = 5928 | 0x0
	sbfnr_GEOFastCorrDegr_1        = 5929
	sbfid_GEOFastCorrDegr_1_0      = 5929 | 0x0
	sbfnr_GEONav_1                 = 5896
	sbfid_GEONav_1_0               = 5896 | 0x0
	sbfnr_GEODegrFactors_1         = 5930
	sbfid_GEODegrFactors_1_0       = 5930 | 0x0
	sbfnr_GEONetworkTime_1         = 5918
	sbfid_GEONetworkTime_1_0       = 5918 | 0x0
	sbfnr_GEOAlm_1                 = 5897
	sbfid_GEOAlm_1_0               = 5897 | 0x0
	sbfnr_GEOIGPMask_1             = 5931
	sbfid_GEOIGPMask_1_0           = 5931 | 0x0
	sbfnr_GEOLongTermCorr_1        = 5932
	sbfid_GEOLongTermCorr_1_0      = 5932 | 0x0
	sbfnr_GEOIonoDelay_1           = 5933
	sbfid_GEOIonoDelay_1_0         = 5933 | 0x0
	sbfnr_GEOServiceLevel_1        = 5917
	sbfid_GEOServiceLevel_1_0      = 5917 | 0x0
	sbfnr_GEOClockEphCovMatrix_1   = 5934
	sbfid_GEOClockEphCovMatrix_1_0 = 5934 | 0x0
	/* SBAS L5 Decoded Message Blocks */
	sbfnr_SBASL5Nav_1   = 5958
	sbfid_SBASL5Nav_1_0 = 5958 | 0x0
	sbfnr_SBASL5Alm_1   = 5959
	sbfid_SBASL5Alm_1_0 = 5959 | 0x0
	/* Position, Velocity and Time Blocks */
	sbfnr_PVTCartesian_1      = 5903
	sbfid_PVTCartesian_1_0    = 5903 | 0x0
	sbfnr_PVTGeodetic_1       = 5904
	sbfid_PVTGeodetic_1_0     = 5904 | 0x0
	sbfnr_DOP_1               = 5909
	sbfid_DOP_1_0             = 5909 | 0x0
	sbfnr_PVTResiduals_1      = 5910
	sbfid_PVTResiduals_1_0    = 5910 | 0x0
	sbfnr_RAIMStatistics_1    = 5915
	sbfid_RAIMStatistics_1_0  = 5915 | 0x0
	sbfnr_PVTCartesian_2      = 4006
	sbfid_PVTCartesian_2_0    = 4006 | 0x0
	sbfid_PVTCartesian_2_1    = 4006 | 0x2000
	sbfid_PVTCartesian_2_2    = 4006 | 0x4000
	sbfnr_PVTGeodetic_2       = 4007
	sbfid_PVTGeodetic_2_0     = 4007 | 0x0
	sbfid_PVTGeodetic_2_1     = 4007 | 0x2000
	sbfid_PVTGeodetic_2_2     = 4007 | 0x4000
	sbfnr_PVTGeodeticAuth_1   = 4232
	sbfid_PVTGeodeticAuth_1_0 = 4232 | 0x0
	sbfid_PVTGeodeticAuth_1_1 = 4232 | 0x2000
	sbfid_PVTGeodeticAuth_1_2 = 4232 | 0x4000
	sbfnr_PosCovCartesian_1   = 5905
	sbfid_PosCovCartesian_1_0 = 5905 | 0x0
	sbfnr_PosCovGeodetic_1    = 5906
	sbfid_PosCovGeodetic_1_0  = 5906 | 0x0
	sbfnr_VelCovCartesian_1   = 5907
	sbfid_VelCovCartesian_1_0 = 5907 | 0x0
	sbfnr_VelCovGeodetic_1    = 5908
	sbfid_VelCovGeodetic_1_0  = 5908 | 0x0
	sbfnr_DOP_2               = 4001
	sbfid_DOP_2_0             = 4001 | 0x0
	sbfnr_PosCart_1           = 4044
	sbfid_PosCart_1_0         = 4044 | 0x0
	sbfnr_PosLocal_1          = 4052
	sbfid_PosLocal_1_0        = 4052 | 0x0
	sbfnr_PosProjected_1      = 4094
	sbfid_PosProjected_1_0    = 4094 | 0x0
	sbfnr_PVTSatCartesian_1   = 4008
	sbfid_PVTSatCartesian_1_0 = 4008 | 0x0
	sbfid_PVTSatCartesian_1_1 = 4008 | 0x2000
	sbfnr_PVTResiduals_2      = 4009
	sbfid_PVTResiduals_2_0    = 4009 | 0x0
	sbfid_PVTResiduals_2_1    = 4009 | 0x2000
	sbfnr_RAIMStatistics_2    = 4011
	sbfid_RAIMStatistics_2_0  = 4011 | 0x0
	sbfnr_GEOCorrections_1    = 5935
	sbfid_GEOCorrections_1_0  = 5935 | 0x0
	sbfnr_BaseVectorCart_1    = 4043
	sbfid_BaseVectorCart_1_0  = 4043 | 0x0
	sbfnr_BaseVectorGeod_1    = 4028
	sbfid_BaseVectorGeod_1_0  = 4028 | 0x0
	sbfnr_Ambiguities_1       = 4240
	sbfid_Ambiguities_1_0     = 4240 | 0x0
	sbfnr_EndOfPVT_1          = 5921
	sbfid_EndOfPVT_1_0        = 5921 | 0x0
	sbfnr_BaseLine_1          = 5950
	sbfid_BaseLine_1_0        = 5950 | 0x0
	/* INS/GNSS Integrated Blocks */
	sbfnr_IntPVCart_1        = 4060
	sbfid_IntPVCart_1_0      = 4060 | 0x0
	sbfnr_IntPVGeod_1        = 4061
	sbfid_IntPVGeod_1_0      = 4061 | 0x0
	sbfnr_IntPosCovCart_1    = 4062
	sbfid_IntPosCovCart_1_0  = 4062 | 0x0
	sbfnr_IntVelCovCart_1    = 4063
	sbfid_IntVelCovCart_1_0  = 4063 | 0x0
	sbfnr_IntPosCovGeod_1    = 4064
	sbfid_IntPosCovGeod_1_0  = 4064 | 0x0
	sbfnr_IntVelCovGeod_1    = 4065
	sbfid_IntVelCovGeod_1_0  = 4065 | 0x0
	sbfnr_IntAttEuler_1      = 4070
	sbfid_IntAttEuler_1_0    = 4070 | 0x0
	sbfid_IntAttEuler_1_1    = 4070 | 0x2000
	sbfnr_IntAttCovEuler_1   = 4072
	sbfid_IntAttCovEuler_1_0 = 4072 | 0x0
	sbfnr_IntPVAAGeod_1      = 4045
	sbfid_IntPVAAGeod_1_0    = 4045 | 0x0
	sbfnr_INSNavCart_1       = 4225
	sbfid_INSNavCart_1_0     = 4225 | 0x0
	sbfnr_INSNavGeod_1       = 4226
	sbfid_INSNavGeod_1_0     = 4226 | 0x0
	sbfnr_IMUBias_1          = 4241
	sbfid_IMUBias_1_0        = 4241 | 0x0
	/* GNSS Attitude Blocks */
	sbfnr_AttEuler_1          = 5938
	sbfid_AttEuler_1_0        = 5938 | 0x0
	sbfnr_AttCovEuler_1       = 5939
	sbfid_AttCovEuler_1_0     = 5939 | 0x0
	sbfnr_AuxAntPositions_1   = 5942
	sbfid_AuxAntPositions_1_0 = 5942 | 0x0
	sbfnr_EndOfAtt_1          = 5943
	sbfid_EndOfAtt_1_0        = 5943 | 0x0
	sbfnr_AttQuat_1           = 5940
	sbfid_AttQuat_1_0         = 5940 | 0x0
	sbfnr_AttCovQuat_1        = 5941
	sbfid_AttCovQuat_1_0      = 5941 | 0x0
	/* Receiver Time Blocks */
	sbfnr_ReceiverTime_1    = 5914
	sbfid_ReceiverTime_1_0  = 5914 | 0x0
	sbfnr_xPPSOffset_1      = 5911
	sbfid_xPPSOffset_1_0    = 5911 | 0x0
	sbfnr_SysTimeOffset_1   = 4039
	sbfid_SysTimeOffset_1_0 = 4039 | 0x0
	sbfid_SysTimeOffset_1_1 = 4039 | 0x2000
	/* External Event Blocks */
	sbfnr_ExtEvent_1               = 5924
	sbfid_ExtEvent_1_0             = 5924 | 0x0
	sbfid_ExtEvent_1_1             = 5924 | 0x2000
	sbfnr_ExtEventPVTCartesian_1   = 4037
	sbfid_ExtEventPVTCartesian_1_0 = 4037 | 0x0
	sbfid_ExtEventPVTCartesian_1_1 = 4037 | 0x2000
	sbfid_ExtEventPVTCartesian_1_2 = 4037 | 0x4000
	sbfnr_ExtEventPVTGeodetic_1    = 4038
	sbfid_ExtEventPVTGeodetic_1_0  = 4038 | 0x0
	sbfid_ExtEventPVTGeodetic_1_1  = 4038 | 0x2000
	sbfid_ExtEventPVTGeodetic_1_2  = 4038 | 0x4000
	sbfnr_ExtEventBaseVectCart_1   = 4216
	sbfid_ExtEventBaseVectCart_1_0 = 4216 | 0x0
	sbfnr_ExtEventBaseVectGeod_1   = 4217
	sbfid_ExtEventBaseVectGeod_1_0 = 4217 | 0x0
	sbfnr_ExtEventINSNavCart_1     = 4229
	sbfid_ExtEventINSNavCart_1_0   = 4229 | 0x0
	sbfnr_ExtEventINSNavGeod_1     = 4230
	sbfid_ExtEventINSNavGeod_1_0   = 4230 | 0x0
	sbfnr_ExtEventAttEuler_1       = 4237
	sbfid_ExtEventAttEuler_1_0     = 4237 | 0x0
	/* Differential Correction Blocks */
	sbfnr_DiffCorrIn_1    = 5919
	sbfid_DiffCorrIn_1_0  = 5919 | 0x0
	sbfnr_BaseStation_1   = 5949
	sbfid_BaseStation_1_0 = 5949 | 0x0
	sbfnr_RTCMDatum_1     = 4049
	sbfid_RTCMDatum_1_0   = 4049 | 0x0
	sbfnr_BaseLink_1      = 5948
	sbfid_BaseLink_1_0    = 5948 | 0x0
	/* L-Band Demodulator Blocks */
	sbfnr_LBandReceiverStatus_1   = 4200
	sbfid_LBandReceiverStatus_1_0 = 4200 | 0x0
	sbfnr_LBandTrackerStatus_1    = 4201
	sbfid_LBandTrackerStatus_1_0  = 4201 | 0x0
	sbfid_LBandTrackerStatus_1_1  = 4201 | 0x2000
	sbfid_LBandTrackerStatus_1_2  = 4201 | 0x4000
	sbfid_LBandTrackerStatus_1_3  = 4201 | 0x6000
	sbfnr_LBAS1DecoderStatus_1    = 4202
	sbfid_LBAS1DecoderStatus_1_0  = 4202 | 0x0
	sbfid_LBAS1DecoderStatus_1_1  = 4202 | 0x2000
	sbfid_LBAS1DecoderStatus_1_2  = 4202 | 0x4000
	sbfnr_LBAS1Messages_1         = 4203
	sbfid_LBAS1Messages_1_0       = 4203 | 0x0
	sbfnr_LBandBeams_1            = 4204
	sbfid_LBandBeams_1_0          = 4204 | 0x0
	sbfnr_LBandRaw_1              = 4212
	sbfid_LBandRaw_1_0            = 4212 | 0x0
	sbfnr_FugroStatus_1           = 4214
	sbfid_FugroStatus_1_0         = 4214 | 0x0
	/* External Sensor Blocks */
	sbfnr_ExtSensorMeas_1     = 4050
	sbfid_ExtSensorMeas_1_0   = 4050 | 0x0
	sbfnr_ExtSensorStatus_1   = 4056
	sbfid_ExtSensorStatus_1_0 = 4056 | 0x0
	sbfnr_ExtSensorSetup_1    = 4057
	sbfid_ExtSensorSetup_1_0  = 4057 | 0x0
	sbfid_ExtSensorSetup_1_1  = 4057 | 0x2000
	sbfid_ExtSensorSetup_1_2  = 4057 | 0x4000
	sbfnr_ExtSensorStatus_2   = 4223
	sbfid_ExtSensorStatus_2_0 = 4223 | 0x0
	sbfnr_ExtSensorInfo_1     = 4222
	sbfid_ExtSensorInfo_1_0   = 4222 | 0x0
	sbfnr_IMUSetup_1          = 4224
	sbfid_IMUSetup_1_0        = 4224 | 0x0
	/* Status Blocks */
	sbfnr_ReceiverStatus_1         = 5913
	sbfid_ReceiverStatus_1_0       = 5913 | 0x0
	sbfnr_TrackingStatus_1         = 5912
	sbfid_TrackingStatus_1_0       = 5912 | 0x0
	sbfnr_ChannelStatus_1          = 4013
	sbfid_ChannelStatus_1_0        = 4013 | 0x0
	sbfnr_ReceiverStatus_2         = 4014
	sbfid_ReceiverStatus_2_0       = 4014 | 0x0
	sbfid_ReceiverStatus_2_1       = 4014 | 0x2000
	sbfnr_SatVisibility_1          = 4012
	sbfid_SatVisibility_1_0        = 4012 | 0x0
	sbfnr_InputLink_1              = 4090
	sbfid_InputLink_1_0            = 4090 | 0x0
	sbfnr_OutputLink_1             = 4091
	sbfid_OutputLink_1_0           = 4091 | 0x0
	sbfid_OutputLink_1_1           = 4091 | 0x2000
	sbfnr_NTRIPClientStatus_1      = 4053
	sbfid_NTRIPClientStatus_1_0    = 4053 | 0x0
	sbfnr_NTRIPServerStatus_1      = 4122
	sbfid_NTRIPServerStatus_1_0    = 4122 | 0x0
	sbfnr_IPStatus_1               = 4058
	sbfid_IPStatus_1_0             = 4058 | 0x0
	sbfid_IPStatus_1_1             = 4058 | 0x2000
	sbfnr_WiFiAPStatus_1           = 4054
	sbfid_WiFiAPStatus_1_0         = 4054 | 0x0
	sbfnr_WiFiClientStatus_1       = 4096
	sbfid_WiFiClientStatus_1_0     = 4096 | 0x0
	sbfnr_CellularStatus_1         = 4055
	sbfid_CellularStatus_1_0       = 4055 | 0x0
	sbfid_CellularStatus_1_1       = 4055 | 0x2000
	sbfnr_BluetoothStatus_1        = 4051
	sbfid_BluetoothStatus_1_0      = 4051 | 0x0
	sbfnr_DynDNSStatus_1           = 4105
	sbfid_DynDNSStatus_1_0         = 4105 | 0x0
	sbfid_DynDNSStatus_1_1         = 4105 | 0x2000
	sbfnr_BatteryStatus_1          = 4083
	sbfid_BatteryStatus_1_0        = 4083 | 0x0
	sbfid_BatteryStatus_1_1        = 4083 | 0x2000
	sbfid_BatteryStatus_1_2        = 4083 | 0x4000
	sbfnr_PowerStatus_1            = 4101
	sbfid_PowerStatus_1_0          = 4101 | 0x0
	sbfnr_QualityInd_1             = 4082
	sbfid_QualityInd_1_0           = 4082 | 0x0
	sbfnr_DiskStatus_1             = 4059
	sbfid_DiskStatus_1_0           = 4059 | 0x0
	sbfid_DiskStatus_1_1           = 4059 | 0x2000
	sbfnr_LogStatus_1              = 4102
	sbfid_LogStatus_1_0            = 4102 | 0x0
	sbfnr_UHFStatus_1              = 4085
	sbfid_UHFStatus_1_0            = 4085 | 0x0
	sbfnr_RFStatus_1               = 4092
	sbfid_RFStatus_1_0             = 4092 | 0x0
	sbfnr_RIMSHealth_1             = 4089
	sbfid_RIMSHealth_1_0           = 4089 | 0x0
	sbfnr_OSNMAStatus_1            = 4231
	sbfid_OSNMAStatus_1_0          = 4231 | 0x0
	sbfnr_GALNavMonitor_1          = 4108
	sbfid_GALNavMonitor_1_0        = 4108 | 0x0
	sbfnr_INAVmonitor_1            = 4233
	sbfid_INAVmonitor_1_0          = 4233 | 0x0
	sbfnr_P2PPStatus_1             = 4238
	sbfid_P2PPStatus_1_0           = 4238 | 0x0
	sbfnr_AuthenticationStatus_1   = 4239
	sbfid_AuthenticationStatus_1_0 = 4239 | 0x0
	sbfnr_CosmosStatus_1           = 4243
	sbfid_CosmosStatus_1_0         = 4243 | 0x0
	/* Miscellaneous Blocks */
	sbfnr_ReceiverSetup_1        = 5902
	sbfid_ReceiverSetup_1_0      = 5902 | 0x0
	sbfid_ReceiverSetup_1_1      = 5902 | 0x2000
	sbfid_ReceiverSetup_1_2      = 5902 | 0x4000
	sbfid_ReceiverSetup_1_3      = 5902 | 0x6000
	sbfid_ReceiverSetup_1_4      = 5902 | 0x8000
	sbfnr_RxComponents_1         = 4084
	sbfid_RxComponents_1_0       = 4084 | 0x0
	sbfnr_RxMessage_1            = 4103
	sbfid_RxMessage_1_0          = 4103 | 0x0
	sbfnr_Commands_1             = 4015
	sbfid_Commands_1_0           = 4015 | 0x0
	sbfnr_Comment_1              = 5936
	sbfid_Comment_1_0            = 5936 | 0x0
	sbfnr_BBSamples_1            = 4040
	sbfid_BBSamples_1_0          = 4040 | 0x0
	sbfnr_ASCIIIn_1              = 4075
	sbfid_ASCIIIn_1_0            = 4075 | 0x0
	sbfnr_EncapsulatedOutput_1   = 4097
	sbfid_EncapsulatedOutput_1_0 = 4097 | 0x0
	sbfnr_RawDataIn_1            = 4236
	sbfid_RawDataIn_1_0          = 4236 | 0x0
	/* TUR Specific Blocks */
	sbfnr_TURPVTSatCorrections_1   = 4035
	sbfid_TURPVTSatCorrections_1_0 = 4035 | 0x0
	sbfid_TURPVTSatCorrections_1_1 = 4035 | 0x2000
	sbfnr_TURHPCAInfo_1            = 4010
	sbfid_TURHPCAInfo_1_0          = 4010 | 0x0
	sbfnr_CorrPeakSample_1         = 4016
	sbfid_CorrPeakSample_1_0       = 4016 | 0x0
	sbfnr_CorrValues_1             = 4100
	sbfid_CorrValues_1_0           = 4100 | 0x0
	sbfnr_TURStatus_1              = 4041
	sbfid_TURStatus_1_0            = 4041 | 0x0
	sbfid_TURStatus_1_1            = 4041 | 0x2000
	sbfid_TURStatus_1_2            = 4041 | 0x4000
	sbfnr_GALIntegrity_1           = 4033
	sbfid_GALIntegrity_1_0         = 4033 | 0x0
	sbfnr_TURFormat_1              = 4080
	sbfid_TURFormat_1_0            = 4080 | 0x0
	sbfnr_CalibrationValues_1      = 4215
	sbfid_CalibrationValues_1_0    = 4215 | 0x0
	sbfnr_MultipathMonitor_1       = 4220
	sbfid_MultipathMonitor_1_0     = 4220 | 0x0
	sbfnr_FOCTURNStatus_1          = 4234
	sbfid_FOCTURNStatus_1_0        = 4234 | 0x0
	sbfnr_TGVFXStatus_1            = 4235
	sbfid_TGVFXStatus_1_0          = 4235 | 0x0
	/* PinPoint-GIS RX */
	sbfnr_GISAction_1   = 4106
	sbfid_GISAction_1_0 = 4106 | 0x0
	sbfnr_GISStatus_1   = 4107
	sbfid_GISStatus_1_0 = 4107 | 0x0
	/* some special values */
	//
	// These constants are not used anywhere. Commenting them out.
	//
	//sbfnr_MIN     = 4000
	//sbfnr_MAX     = 6015 // This is the maximum allowed SBF ID which can be output by a Septentrio receiver using the SBF protocol as is today.
	//sbfid_ALL     = 0xFFFF
	//sbfnr_INVALID = 0x0
)

/*==CONNECTION DESCRIPTOR definition=========================================*/
const (
	CD_COM1 = 1
	CD_COM2 = 2
	CD_COM3 = 3
	CD_COM4 = 4
	CD_USB1 = 33
	CD_USB2 = 34
	CD_OTG1 = 49
	CD_OTG2 = 50
	CD_IP10 = 64
	CD_IP11 = 65
	CD_IP12 = 66
	CD_IP13 = 67
	CD_IP14 = 68
	CD_IP15 = 69
	CD_IP16 = 70
	CD_IP17 = 71
	CD_DSK1 = 97
	CD_DSK2 = 98
	CD_NTR1 = 129
	CD_NTR2 = 130
	CD_NTR3 = 131
	CD_NTR4 = 132
	CD_NTR5 = 133
	CD_IPS1 = 161
	CD_IPS2 = 162
	CD_IPS3 = 163
	CD_IPS4 = 164
	CD_IPS5 = 165
	CD_BT01 = 192
	CD_BT02 = 193
	CD_UHF1 = 196
	CD_IPR1 = 201
	CD_IPR2 = 202
	CD_IPR3 = 203
	CD_IPR4 = 204
	CD_IPR5 = 205
	CD_DCL1 = 210
	CD_CAN1 = 214
	CD_SPI1 = 220
)

/*==DATA TYPE definition=====================================================*/
const (
	DATA_NONE         = 0
	DATA_DAISYCHAIN   = 1
	DATA_CMD          = 32
	DATA_SBF          = 33
	DATA_ASCIIDISPLAY = 34
	DATA_RINEX        = 35
	DATA_CGGTTS       = 36
	DATA_BINEX        = 40
	DATA_NMEA         = 64
	DATA_RTCMV2       = 96
	DATA_RTCMV3       = 97
	DATA_CMRV2        = 98
	DATA_RTCMV        = 99
	DATA_SPARTN       = 100
	DATA_LBMP         = 101
	DATA_LBAS1        = 110
	DATA_LBAS2        = 111
	DATA_LBANDBEAM1   = 118
	DATA_LBANDBEAM2   = 119
	DATA_LBANDBEAM3   = 120
	DATA_LBANDBEAM4   = 121
	DATA_MTI          = 128
	DATA_MMQ          = 129
	DATA_ELLIPSE      = 130
	DATA_SBG          = 131
	DATA_MTI10        = 132
	DATA_SPIRENT      = 133
	DATA_VECTORNAV    = 134
	DATA_HONEYWELL    = 135
	DATA_RIVIANIMU    = 136
	DATA_ADIS         = 137
	DATA_ASCIIIN      = 160
)

/*==Measurement Blocks=======================================================*/
const SBF_GENMEASEPOCH_1_0_TOW_PRECISION = 3
const SBF_GENMEASEPOCH_1_DATA_LENGTH = SBF_GENMEASEPOCH_1_0_DATA_LENGTH

/*--MeasEpoch_2_0_t : -------------------------------------------------------*/
/* Measurement set of one epoch */
const SBF_MEASEPOCH_2_0_TOW_PRECISION = 3
const SBF_MEASEPOCH_2_1_TOW_PRECISION = 3
const SBF_MEASEPOCH_2_DATA_LENGTH = SBF_MEASEPOCH_2_1_DATA_LENGTH

/*--MeasExtra_1_0_t : -------------------------------------------------------*/
/* Additional info such as observable variance */
const SBF_MEASEXTRA_1_0_MEASEXTRACHANNEL_LENGTH = 648
const SBF_MEASEXTRA_1_0_TOW_PRECISION = 3

/*--MeasExtra_1_1_t : -------------------------------------------------------*/
/* Additional info such as observable variance */
const SBF_MEASEXTRACHANNELSUB_1_1__PADDING_LENGTH = 2
const SBF_MEASEXTRA_1_1_MEASEXTRACHANNEL_LENGTH = 648
const SBF_MEASEXTRA_1_1_TOW_PRECISION = 3

/*--MeasExtra_1_2_t : -------------------------------------------------------*/
/* Additional info such as observable variance */
const SBF_MEASEXTRACHANNELSUB_1_2__PADDING_LENGTH = 1
const SBF_MEASEXTRA_1_2_MEASEXTRACHANNEL_LENGTH = 648
const SBF_MEASEXTRA_1_2_TOW_PRECISION = 3

/*--MeasExtra_1_3_t : -------------------------------------------------------*/
/* Additional info such as observable variance */
const SBF_MEASEXTRA_1_3_MEASEXTRACHANNEL_LENGTH = 648
const SBF_MEASEXTRA_1_3_TOW_PRECISION = 3
const SBF_MEASEXTRA_1_MEASEXTRACHANNEL_LENGTH = SBF_MEASEXTRA_1_3_MEASEXTRACHANNEL_LENGTH

/*--MeasFullRange_1_0_t : ---------------------------------------------------*/
/* Extended-range code and phase measurements */
const SBF_MEASFULLRANGESUB_1_0_CARRIERMINCODE_PRECISION = 3
const SBF_MEASFULLRANGE_1_0_MEASFULLRANGESUB_LENGTH = 648
const SBF_MEASFULLRANGE_1_0_TOW_PRECISION = 3

/*--MeasFullRange_1_1_t : ---------------------------------------------------*/
/* Extended-range code and phase measurements */
const SBF_MEASFULLRANGESUB_1_1__PADDING_LENGTH = 1
const SBF_MEASFULLRANGESUB_1_1_CARRIERMINCODE_PRECISION = 3
const SBF_MEASFULLRANGE_1_1_MEASFULLRANGESUB_LENGTH = 648
const SBF_MEASFULLRANGE_1_1_TOW_PRECISION = 3
const SBF_MEASFULLRANGESUB_1__PADDING_LENGTH = SBF_MEASFULLRANGESUB_1_1__PADDING_LENGTH
const SBF_MEASFULLRANGE_1_MEASFULLRANGESUB_LENGTH = SBF_MEASFULLRANGE_1_1_MEASFULLRANGESUB_LENGTH

/*--Meas3Ranges_1_0_t : -----------------------------------------------------*/
/* Code, phase and CN0 measurements */
const SBF_M3SATDATA_1_0_SATMASK_LENGTH = 8
const SBF_M3SATDATA_1_0_GLOFNLIST_LENGTH = 8
const SBF_MEAS3RANGES_1_0_TOW_PRECISION = 3
const SBF_M3SATDATA_1_SATMASK_LENGTH = SBF_M3SATDATA_1_0_SATMASK_LENGTH
const SBF_M3SATDATA_1_GLOFNLIST_LENGTH = SBF_M3SATDATA_1_0_GLOFNLIST_LENGTH
const SBF_MEAS3RANGES_1_DATA_LENGTH = SBF_MEAS3RANGES_1_0_DATA_LENGTH

/*--Meas3CN0HiRes_1_0_t : ---------------------------------------------------*/
/* Extension of Meas3Ranges containing fractional C/N0 values */
const SBF_MEAS3CN0HIRES_1_0_CN0HIRES_LENGTH = 1024
const SBF_MEAS3CN0HIRES_1_0_TOW_PRECISION = 3
const SBF_MEAS3CN0HIRES_1_CN0HIRES_LENGTH = SBF_MEAS3CN0HIRES_1_0_CN0HIRES_LENGTH

/*--Meas3Doppler_1_0_t : ----------------------------------------------------*/
/* Extension of Meas3Ranges containing Doppler values */
const SBF_MEAS3DOPPLER_1_0_DATA_LENGTH = 2048
const SBF_MEAS3DOPPLER_1_0_TOW_PRECISION = 3
const SBF_MEAS3DOPPLER_1_DATA_LENGTH = SBF_MEAS3DOPPLER_1_0_DATA_LENGTH

/*--Meas3PP_1_0_t : ---------------------------------------------------------*/
/* Extension of Meas3Ranges containing proprietary flags for data post-processing. */
const SBF_MEAS3PP_1_0_DATA_LENGTH = 2048
const SBF_MEAS3PP_1_0_TOW_PRECISION = 3
const SBF_MEAS3PP_1_DATA_LENGTH = SBF_MEAS3PP_1_0_DATA_LENGTH

/*--Meas3MP_1_0_t : ---------------------------------------------------------*/
/* Extension of Meas3Ranges containing multipath corrections applied by the receiver. */
const SBF_MEAS3MP_1_0_DATA_LENGTH = 2048
const SBF_MEAS3MP_1_0_TOW_PRECISION = 3
const SBF_MEAS3MP_1_DATA_LENGTH = SBF_MEAS3MP_1_0_DATA_LENGTH

/*--IQCorr_1_0_t : ----------------------------------------------------------*/
/* Real and imaginary post-correlation values */
const SBF_IQCORR_1_0_CORRCHANNEL_LENGTH = 648
const SBF_IQCORR_1_0_TOW_PRECISION = 3

/*--IQCorr_1_1_t : ----------------------------------------------------------*/
/* Real and imaginary post-correlation values */
const SBF_IQCORR_1_1_CORRCHANNEL_LENGTH = 648
const SBF_IQCORR_1_1_TOW_PRECISION = 3
const SBF_IQCORR_1_CORRCHANNEL_LENGTH = SBF_IQCORR_1_1_CORRCHANNEL_LENGTH

/*--ISMR_1_0_t : ------------------------------------------------------------*/
/* Ionospheric scintillation monitor (ISMR) data */
const SBF_ISMR_1_0_ISMRCHANNEL_LENGTH = 648
const SBF_ISMR_1_0_TOW_PRECISION = 3
const SBF_ISMR_1_ISMRCHANNEL_LENGTH = SBF_ISMR_1_0_ISMRCHANNEL_LENGTH

/*--SQMSamples_1_0_t : ------------------------------------------------------*/
/* Correlation samples for signal quality monitoring */
const SBF_SQMSAMPLES_1_0_TOW_PRECISION = 3
const SBF_SQMSAMPLES_1_DATA_LENGTH = SBF_SQMSAMPLES_1_0_DATA_LENGTH

/*--EndOfMeas_1_0_t : -------------------------------------------------------*/
/* Measurement epoch marker */
const SBF_ENDOFMEAS_1_0_TOW_PRECISION = 3

/*==Navigation Page Blocks===================================================*/

/*--GPSRaw_1_0_t : ----------------------------------------------------------*/
/* GPS CA navigation frame */
const SBF_GPSRAW_1_0_RAWBITS_LENGTH = 10
const SBF_GPSRAW_1_0_TOW_PRECISION = 3
const SBF_GPSRAW_1_RAWBITS_LENGTH = SBF_GPSRAW_1_0_RAWBITS_LENGTH

/*--CNAVRaw_1_0_t : ---------------------------------------------------------*/
/* GPS L2C navigation frame */
const SBF_CNAVRAW_1_0_CNAVBITS_LENGTH = 10
const SBF_CNAVRAW_1_0_TOW_PRECISION = 3
const SBF_CNAVRAW_1_CNAVBITS_LENGTH = SBF_CNAVRAW_1_0_CNAVBITS_LENGTH

/*--GEORaw_1_0_t : ----------------------------------------------------------*/
/* SBAS L1 navigation frame */
const SBF_GEORAW_1_0_RAWBITS_LENGTH = 8
const SBF_GEORAW_1_0_TOW_PRECISION = 3
const SBF_GEORAW_1_RAWBITS_LENGTH = SBF_GEORAW_1_0_RAWBITS_LENGTH

/*--GPSRawCA_1_0_t : --------------------------------------------------------*/
/* GPS CA navigation subframe */
const SBF_GPSRAWCA_1_0_NAVBITS_LENGTH = 10
const SBF_GPSRAWCA_1_0_TOW_PRECISION = 3
const SBF_GPSRAWCA_1_NAVBITS_LENGTH = SBF_GPSRAWCA_1_0_NAVBITS_LENGTH

/*--GPSRawL2C_1_0_t : -------------------------------------------------------*/
/* GPS L2C navigation frame */
const SBF_GPSRAWL2C_1_0_NAVBITS_LENGTH = 10
const SBF_GPSRAWL2C_1_0_TOW_PRECISION = 3
const SBF_GPSRAWL2C_1_NAVBITS_LENGTH = SBF_GPSRAWL2C_1_0_NAVBITS_LENGTH

/*--GPSRawL5_1_0_t : --------------------------------------------------------*/
/* GPS L5 navigation frame */
const SBF_GPSRAWL5_1_0_NAVBITS_LENGTH = 10
const SBF_GPSRAWL5_1_0_TOW_PRECISION = 3
const SBF_GPSRAWL5_1_NAVBITS_LENGTH = SBF_GPSRAWL5_1_0_NAVBITS_LENGTH

/*--GPSRawL1C_1_0_t : -------------------------------------------------------*/
/* GPS L1C navigation frame */
const SBF_GPSRAWL1C_1_0_NAVBITS_LENGTH = 57
const SBF_GPSRAWL1C_1_0_TOW_PRECISION = 3
const SBF_GPSRAWL1C_1_NAVBITS_LENGTH = SBF_GPSRAWL1C_1_0_NAVBITS_LENGTH

/*--GLORawCA_1_0_t : --------------------------------------------------------*/
/* GLONASS CA navigation string */
const SBF_GLORAWCA_1_0_NAVBITS_LENGTH = 3
const SBF_GLORAWCA_1_0_TOW_PRECISION = 3
const SBF_GLORAWCA_1_NAVBITS_LENGTH = SBF_GLORAWCA_1_0_NAVBITS_LENGTH

/*--GALRawFNAV_1_0_t : ------------------------------------------------------*/
/* Galileo F/NAV navigation page */
const SBF_GALRAWFNAV_1_0_NAVBITS_LENGTH = 8
const SBF_GALRAWFNAV_1_0_TOW_PRECISION = 3
const SBF_GALRAWFNAV_1_NAVBITS_LENGTH = SBF_GALRAWFNAV_1_0_NAVBITS_LENGTH

/*--GALRawINAV_1_0_t : ------------------------------------------------------*/
/* Galileo I/NAV navigation page */
const SBF_GALRAWINAV_1_0_NAVBITS_LENGTH = 8
const SBF_GALRAWINAV_1_0_TOW_PRECISION = 3
const SBF_GALRAWINAV_1_NAVBITS_LENGTH = SBF_GALRAWINAV_1_0_NAVBITS_LENGTH

/*--GALRawCNAV_1_0_t : ------------------------------------------------------*/
/* Galileo C/NAV navigation page */
const SBF_GALRAWCNAV_1_0_NAVBITS_LENGTH = 16
const SBF_GALRAWCNAV_1_0_TOW_PRECISION = 3
const SBF_GALRAWCNAV_1_NAVBITS_LENGTH = SBF_GALRAWCNAV_1_0_NAVBITS_LENGTH

/*--GALRawGNAV_1_0_t : ------------------------------------------------------*/
const SBF_GALRAWGNAV_1_0_NAVBITS_LENGTH = 6
const SBF_GALRAWGNAV_1_0_TOW_PRECISION = 3
const SBF_GALRAWGNAV_1_NAVBITS_LENGTH = SBF_GALRAWGNAV_1_0_NAVBITS_LENGTH

/*--GALRawGNAVe_1_0_t : -----------------------------------------------------*/
const SBF_GALRAWGNAVE_1_0_NAVBITS_LENGTH = 11
const SBF_GALRAWGNAVE_1_0_TOW_PRECISION = 3
const SBF_GALRAWGNAVE_1_NAVBITS_LENGTH = SBF_GALRAWGNAVE_1_0_NAVBITS_LENGTH

/*--GEORawL1_1_0_t : --------------------------------------------------------*/
/* SBAS L1 navigation message */
const SBF_GEORAWL1_1_0_NAVBITS_LENGTH = 8
const SBF_GEORAWL1_1_0_TOW_PRECISION = 3
const SBF_GEORAWL1_1_NAVBITS_LENGTH = SBF_GEORAWL1_1_0_NAVBITS_LENGTH

/*--GEORawL5_1_0_t : --------------------------------------------------------*/
/* SBAS L5 navigation message */
const SBF_GEORAWL5_1_0_NAVBITS_LENGTH = 8
const SBF_GEORAWL5_1_0_TOW_PRECISION = 3
const SBF_GEORAWL5_1_NAVBITS_LENGTH = SBF_GEORAWL5_1_0_NAVBITS_LENGTH

/*--BDSRaw_1_0_t : ----------------------------------------------------------*/
/* BeiDou navigation page */
const SBF_BDSRAW_1_0_NAVBITS_LENGTH = 10
const SBF_BDSRAW_1_0_TOW_PRECISION = 3
const SBF_BDSRAW_1_NAVBITS_LENGTH = SBF_BDSRAW_1_0_NAVBITS_LENGTH

/*--BDSRawB1C_1_0_t : -------------------------------------------------------*/
/* BeiDou B1C navigation frame */
const SBF_BDSRAWB1C_1_0_NAVBITS_LENGTH = 57
const SBF_BDSRAWB1C_1_0_TOW_PRECISION = 3
const SBF_BDSRAWB1C_1_NAVBITS_LENGTH = SBF_BDSRAWB1C_1_0_NAVBITS_LENGTH

/*--BDSRawB2a_1_0_t : -------------------------------------------------------*/
/* BeiDou B2a navigation frame */
const SBF_BDSRAWB2A_1_0_NAVBITS_LENGTH = 18
const SBF_BDSRAWB2A_1_0_TOW_PRECISION = 3
const SBF_BDSRAWB2A_1_NAVBITS_LENGTH = SBF_BDSRAWB2A_1_0_NAVBITS_LENGTH

/*--BDSRawB2b_1_0_t : -------------------------------------------------------*/
/* BeiDou B2b navigation frame */
const SBF_BDSRAWB2B_1_0_NAVBITS_LENGTH = 31
const SBF_BDSRAWB2B_1_0_TOW_PRECISION = 3
const SBF_BDSRAWB2B_1_NAVBITS_LENGTH = SBF_BDSRAWB2B_1_0_NAVBITS_LENGTH

/*--QZSRawL1CA_1_0_t : ------------------------------------------------------*/
/* QZSS L1 CA navigation frame */
const SBF_QZSRAWL1CA_1_0_NAVBITS_LENGTH = 10
const SBF_QZSRAWL1CA_1_0_TOW_PRECISION = 3
const SBF_QZSRAWL1CA_1_NAVBITS_LENGTH = SBF_QZSRAWL1CA_1_0_NAVBITS_LENGTH

/*--QZSRawL2C_1_0_t : -------------------------------------------------------*/
/* QZSS L2C navigation frame */
const SBF_QZSRAWL2C_1_0_NAVBITS_LENGTH = 10
const SBF_QZSRAWL2C_1_0_TOW_PRECISION = 3
const SBF_QZSRAWL2C_1_NAVBITS_LENGTH = SBF_QZSRAWL2C_1_0_NAVBITS_LENGTH

/*--QZSRawL5_1_0_t : --------------------------------------------------------*/
/* QZSS L5 navigation frame */
const SBF_QZSRAWL5_1_0_NAVBITS_LENGTH = 10
const SBF_QZSRAWL5_1_0_TOW_PRECISION = 3
const SBF_QZSRAWL5_1_NAVBITS_LENGTH = SBF_QZSRAWL5_1_0_NAVBITS_LENGTH

/*--QZSRawL6_1_0_t : --------------------------------------------------------*/
/* QZSS L6 navigation message */
const SBF_QZSRAWL6_1_0_NAVBITS_LENGTH = 63
const SBF_QZSRAWL6_1_0_TOW_PRECISION = 3
const SBF_QZSRAWL6_1_NAVBITS_LENGTH = SBF_QZSRAWL6_1_0_NAVBITS_LENGTH

/*--QZSRawL1C_1_0_t : -------------------------------------------------------*/
/* QZSS L1C navigation frame */
const SBF_QZSRAWL1C_1_0_NAVBITS_LENGTH = 57
const SBF_QZSRAWL1C_1_0_TOW_PRECISION = 3
const SBF_QZSRAWL1C_1_NAVBITS_LENGTH = SBF_QZSRAWL1C_1_0_NAVBITS_LENGTH

/*--QZSRawL1S_1_0_t : -------------------------------------------------------*/
/* QZSS L1S navigation message */
const SBF_QZSRAWL1S_1_0_NAVBITS_LENGTH = 8
const SBF_QZSRAWL1S_1_0_TOW_PRECISION = 3
const SBF_QZSRAWL1S_1_NAVBITS_LENGTH = SBF_QZSRAWL1S_1_0_NAVBITS_LENGTH

/*--NAVICRaw_1_0_t : --------------------------------------------------------*/
/* NavIC/IRNSS subframe */
const SBF_NAVICRAW_1_0_NAVBITS_LENGTH = 10
const SBF_NAVICRAW_1_0_TOW_PRECISION = 3
const SBF_NAVICRAW_1_NAVBITS_LENGTH = SBF_NAVICRAW_1_0_NAVBITS_LENGTH

/*--GNSSNavBits_1_0_t : -----------------------------------------------------*/
/* Raw navigation bits during last second */
const SBF_RAWNAVBITS_1_0_NAVBITS_LENGTH = 16
const SBF_RAWNAVBITS_1_0_TOW_PRECISION = 3
const SBF_RAWNAVBITS_1_NAVBITS_LENGTH = SBF_RAWNAVBITS_1_0_NAVBITS_LENGTH

/*--GNSSSymbols_1_0_t : -----------------------------------------------------*/
/* Raw navigation symbols */
const SBF_RAWSYMBOLS_1_0_SYMBOLS_LENGTH = 32
const SBF_RAWSYMBOLS_1_0_TOW_PRECISION = 3
const SBF_RAWSYMBOLS_1_SYMBOLS_LENGTH = SBF_RAWSYMBOLS_1_0_SYMBOLS_LENGTH

/*==GPS Decoded Message Blocks===============================================*/

/*--GPSNav_1_0_t : ----------------------------------------------------------*/
/* GPS ephemeris and clock */
const SBF_GPEPH_1_0_TOW_PRECISION = 3
const SBF_GPEPH_1_0_T_GD_PRECISION = 10

/*--GPSAlm_1_0_t : ----------------------------------------------------------*/
/* Almanac data for a GPS satellite */
const SBF_GPALM_1_0_TOW_PRECISION = 3

/*--GPSIon_1_0_t : ----------------------------------------------------------*/
/* Ionosphere data from the GPS subframe 5 */
const SBF_GPION_1_0_TOW_PRECISION = 3

/*--GPSUtc_1_0_t : ----------------------------------------------------------*/
/* GPS-UTC data from GPS subframe 5 */
const SBF_GPUTC_1_0_TOW_PRECISION = 3
const SBF_GPUTC_1_0_A_1_PRECISION = 18
const SBF_GPUTC_1_0_A_0_PRECISION = 16

/*--GPSCNav_1_0_t : ---------------------------------------------------------*/
/* CNAV Ephemeris data for one satellite. */
const SBF_GPEPHCNAV_1_0_TOW_PRECISION = 3
const SBF_GPEPHCNAV_1_0_T_GD_PRECISION = 3
const SBF_GPEPHCNAV_1_0_ISC_L1CA_PRECISION = 3
const SBF_GPEPHCNAV_1_0_ISC_L2C_PRECISION = 3
const SBF_GPEPHCNAV_1_0_ISC_L5I5_PRECISION = 3
const SBF_GPEPHCNAV_1_0_ISC_L5Q5_PRECISION = 3

/*==GLONASS Decoded Message Blocks===========================================*/

/*--GLONav_1_0_t : ----------------------------------------------------------*/
/* GLONASS ephemeris and clock */
const SBF_GLEPH_1_0_TOW_PRECISION = 3
const SBF_GLEPH_1_0_GAMMA_PRECISION = 14
const SBF_GLEPH_1_0_DTAU_PRECISION = 11

/*--GLOAlm_1_0_t : ----------------------------------------------------------*/
/* Almanac data for a GLONASS satellite */
const SBF_GLALM_1_0__PADDING_LENGTH = 2
const SBF_GLALM_1_0_TOW_PRECISION = 3
const SBF_GLALM_1__PADDING_LENGTH = SBF_GLALM_1_0__PADDING_LENGTH

/*--GLOTime_1_0_t : ---------------------------------------------------------*/
/* GLO-UTC, GLO-GPS and GLO-UT1 data */
const SBF_GLTIME_1_0_TOW_PRECISION = 3
const SBF_GLTIME_1_0_TAU_GPS_PRECISION = 9
const SBF_GLTIME_1_0_TAU_C_PRECISION = 9

/*==Galileo Decoded Message Blocks===========================================*/

/*--GALNav_1_0_t : ----------------------------------------------------------*/
/* Galileo ephemeris, clock, health and BGD */
const SBF_GAEPH_1_0_TOW_PRECISION = 3
const SBF_GAEPH_1_0_BGD_L1E5A_PRECISION = 12
const SBF_GAEPH_1_0_BGD_L1E5B_PRECISION = 12
const SBF_GAEPH_1_0_BGD_L1AE6A_PRECISION = 12

/*--GALAlm_1_0_t : ----------------------------------------------------------*/
/* Almanac data for a Galileo satellite */
const SBF_GAALM_1_0_TOW_PRECISION = 3

/*--GALIon_1_0_t : ----------------------------------------------------------*/
/* NeQuick Ionosphere model parameters */
const SBF_GAION_1_0_TOW_PRECISION = 3

/*--GALUtc_1_0_t : ----------------------------------------------------------*/
/* GST-UTC data */
const SBF_GAUTC_1_0_TOW_PRECISION = 3
const SBF_GAUTC_1_0_A_1_PRECISION = 3
const SBF_GAUTC_1_0_A_0_PRECISION = 3

/*--GALGstGps_1_0_t : -------------------------------------------------------*/
/* GST-GPS data */
const SBF_GAGSTGPS_1_0_TOW_PRECISION = 3

/*--GALSARRLM_1_0_t : -------------------------------------------------------*/
/* Search-and-rescue return link message */
const SBF_GASARRLM_1_0_RLMBITS_LENGTH = 5
const SBF_GASARRLM_1_0_TOW_PRECISION = 3
const SBF_GASARRLM_1_RLMBITS_LENGTH = SBF_GASARRLM_1_0_RLMBITS_LENGTH

/*==BeiDou Decoded Message Blocks============================================*/

/*--BDSNav_1_0_t : ----------------------------------------------------------*/
/* BeiDou ephemeris and clock */
const SBF_CMPEPH_1_0_TOW_PRECISION = 3
const SBF_CMPEPH_1_0_T_GD1_PRECISION = 10
const SBF_CMPEPH_1_0_T_GD2_PRECISION = 10

/*--BDSAlm_1_0_t : ----------------------------------------------------------*/
/* Almanac data for a BeiDou satellite */
const SBF_CMPALM_1_0_TOW_PRECISION = 3

/*--BDSIon_1_0_t : ----------------------------------------------------------*/
/* BeiDou Ionospheric delay model parameters */
const SBF_CMPION_1_0_TOW_PRECISION = 3

/*--BDSUtc_1_0_t : ----------------------------------------------------------*/
/* BDT-UTC data */
const SBF_CMPUTC_1_0_TOW_PRECISION = 3
const SBF_CMPUTC_1_0_A_1_PRECISION = 18
const SBF_CMPUTC_1_0_A_0_PRECISION = 16

/*==QZSS Decoded Message Blocks==============================================*/

/*--QZSNav_1_0_t : ----------------------------------------------------------*/
/* QZSS ephemeris and clock */
const SBF_QZEPH_1_0_TOW_PRECISION = 3
const SBF_QZEPH_1_0_T_GD_PRECISION = 10

/*--QZSAlm_1_0_t : ----------------------------------------------------------*/
/* Almanac data for a QZSS satellite */
const SBF_QZALM_1_0_TOW_PRECISION = 3

/*==SBAS L1 Decoded Message Blocks===========================================*/

/*--GEOMT00_1_0_t : ---------------------------------------------------------*/
/* MT00 : SBAS Don't use for safety applications */
const SBF_RAMT00_1_0_TOW_PRECISION = 3

/*--GEOPRNMask_1_0_t : ------------------------------------------------------*/
/* MT01 : PRN Mask assignments */
const SBF_RAPRNMASK_1_0_PRNMASK_LENGTH = 51
const SBF_RAPRNMASK_1_0_TOW_PRECISION = 3
const SBF_RAPRNMASK_1_PRNMASK_LENGTH = SBF_RAPRNMASK_1_0_PRNMASK_LENGTH

/*--GEOFastCorr_1_0_t : -----------------------------------------------------*/
/* MT02-05/24: Fast Corrections */
const SBF_RAFASTCORR_1_0_FASTCORR_LENGTH = 13
const SBF_RAFASTCORR_1_0_TOW_PRECISION = 3
const SBF_RAFASTCORR_1_FASTCORR_LENGTH = SBF_RAFASTCORR_1_0_FASTCORR_LENGTH

/*--GEOIntegrity_1_0_t : ----------------------------------------------------*/
/* MT06 : Integrity information */
const SBF_RAINTEGRITY_1_0_IODF_LENGTH = 4
const SBF_RAINTEGRITY_1_0_UDREI_LENGTH = 51
const SBF_RAINTEGRITY_1_0_TOW_PRECISION = 3
const SBF_RAINTEGRITY_1_IODF_LENGTH = SBF_RAINTEGRITY_1_0_IODF_LENGTH
const SBF_RAINTEGRITY_1_UDREI_LENGTH = SBF_RAINTEGRITY_1_0_UDREI_LENGTH

/*--GEOFastCorrDegr_1_0_t : -------------------------------------------------*/
/* MT07 : Fast correction degradation factors */
const SBF_RAFASTCORRDEGR_1_0_AI_LENGTH = 51
const SBF_RAFASTCORRDEGR_1_0_TOW_PRECISION = 3
const SBF_RAFASTCORRDEGR_1_AI_LENGTH = SBF_RAFASTCORRDEGR_1_0_AI_LENGTH

/*--GEONav_1_0_t : ----------------------------------------------------------*/
/* MT09 : SBAS navigation message */
const SBF_RAEPH_1_0_TOW_PRECISION = 3

/*--GEODegrFactors_1_0_t : --------------------------------------------------*/
/* MT10 : Degradation factors */
const SBF_RADEGRFACTORS_1_0_TOW_PRECISION = 3

/*--GEONetworkTime_1_0_t : --------------------------------------------------*/
/* MT12 : SBAS Network Time/UTC offset parameters */
const SBF_RANETWORKTIME_1_0_TOW_PRECISION = 3

/*--GEOAlm_1_0_t : ----------------------------------------------------------*/
/* MT17 : SBAS satellite almanac */
const SBF_RAALM_1_0_TOW_PRECISION = 3

/*--GEOIGPMask_1_0_t : ------------------------------------------------------*/
/* MT18 : Ionospheric grid point mask */
const SBF_RAIGPMASK_1_0_IGPMASK_LENGTH = 201
const SBF_RAIGPMASK_1_0_TOW_PRECISION = 3
const SBF_RAIGPMASK_1_IGPMASK_LENGTH = SBF_RAIGPMASK_1_0_IGPMASK_LENGTH

/*--GEOLongTermCorr_1_0_t : -------------------------------------------------*/
/* MT24/25 : Long term satellite error corrections */
const SBF_RALONGTERMCORR_1_0_LTCORR_LENGTH = 4
const SBF_RALONGTERMCORR_1_0_TOW_PRECISION = 3
const SBF_RALONGTERMCORR_1_LTCORR_LENGTH = SBF_RALONGTERMCORR_1_0_LTCORR_LENGTH

/*--GEOIonoDelay_1_0_t : ----------------------------------------------------*/
/* MT26 : Ionospheric delay corrections */
const SBF_RAIONODELAY_1_0_IDC_LENGTH = 15
const SBF_RAIONODELAY_1_0_TOW_PRECISION = 3
const SBF_RAIONODELAY_1_IDC_LENGTH = SBF_RAIONODELAY_1_0_IDC_LENGTH

/*--GEOServiceLevel_1_0_t : -------------------------------------------------*/
/* MT27 : SBAS Service Message */
const SBF_SERVICEREGION_1_0__PADDING_LENGTH = 1
const SBF_RASERVICEMSG_1_0_REGIONS_LENGTH = 8
const SBF_RASERVICELEVEL_1_0_TOW_PRECISION = 3
const SBF_SERVICEREGION_1__PADDING_LENGTH = SBF_SERVICEREGION_1_0__PADDING_LENGTH
const SBF_RASERVICEMSG_1_REGIONS_LENGTH = SBF_RASERVICEMSG_1_0_REGIONS_LENGTH

/*--GEOClockEphCovMatrix_1_0_t : --------------------------------------------*/
/* MT28 : Clock-Ephemeris Covariance Matrix */
const SBF_RACLOCKEPHCOVMATRIX_1_0_COVMATRIX_LENGTH = 2
const SBF_RACLOCKEPHCOVMATRIX_1_0_TOW_PRECISION = 3
const SBF_RACLOCKEPHCOVMATRIX_1_COVMATRIX_LENGTH = SBF_RACLOCKEPHCOVMATRIX_1_0_COVMATRIX_LENGTH

/*==SBAS L5 Decoded Message Blocks===========================================*/

/*--SBASL5Nav_1_0_t : -------------------------------------------------------*/
/* DFMC SBAS ephemeris and clock data */
const SBF_SBSEPH_1_0_TOW_PRECISION = 3

/*--SBASL5Alm_1_0_t : -------------------------------------------------------*/
/* DFMC SBAS almanac data */
const SBF_SBSALM_1_0__PADDING_LENGTH = 1
const SBF_SBSALM_1_0_TOW_PRECISION = 3
const SBF_SBSALM_1__PADDING_LENGTH = SBF_SBSALM_1_0__PADDING_LENGTH

/*==Position, Velocity and Time Blocks=======================================*/

/*--PVTCartesian_1_0_t : ----------------------------------------------------*/
/* PVT in Cartesian coordinates */
const SBF_PVTCARTESIAN_1_0_TOW_PRECISION = 3

/*--PVTGeodetic_1_0_t : -----------------------------------------------------*/
/* PVT in geodetic coordinates */
const SBF_PVTGEODETIC_1_0_TOW_PRECISION = 3

/*--DOP_1_0_t : -------------------------------------------------------------*/
/* Dilution of precision */
const SBF_DOP_1_0_TOW_PRECISION = 3

/*--PVTResiduals_1_0_t : ----------------------------------------------------*/
/* Measurement residuals */
const SBF_PVTRESIDUAL_1_0__PADDING_LENGTH = 1
const SBF_PVTRESIDUALS_1_0_SATRESIDUAL_LENGTH = 72
const SBF_PVTRESIDUALS_1_0_TOW_PRECISION = 3
const SBF_PVTRESIDUAL_1__PADDING_LENGTH = SBF_PVTRESIDUAL_1_0__PADDING_LENGTH
const SBF_PVTRESIDUALS_1_SATRESIDUAL_LENGTH = SBF_PVTRESIDUALS_1_0_SATRESIDUAL_LENGTH

/*--RAIMStatistics_1_0_t : --------------------------------------------------*/
/* Integrity statistics */
const SBF_RAIMSTATISTICS_1_0_RAIMCHANNEL_LENGTH = 72
const SBF_RAIMSTATISTICS_1_0_TOW_PRECISION = 3
const SBF_RAIMSTATISTICS_1_RAIMCHANNEL_LENGTH = SBF_RAIMSTATISTICS_1_0_RAIMCHANNEL_LENGTH

/*--PVTCartesian_2_0_t : ----------------------------------------------------*/
/* Position, velocity, and time in Cartesian coordinates */
const SBF_PVTCARTESIAN_2_0__PADDING_LENGTH = 3
const SBF_PVTCARTESIAN_2_0_TOW_PRECISION = 3
const SBF_PVTCARTESIAN_2_0_X_PRECISION = 3
const SBF_PVTCARTESIAN_2_0_Y_PRECISION = 3
const SBF_PVTCARTESIAN_2_0_Z_PRECISION = 3
const SBF_PVTCARTESIAN_2_0_UNDULATION_PRECISION = 3
const SBF_PVTCARTESIAN_2_0_VX_PRECISION = 3
const SBF_PVTCARTESIAN_2_0_VY_PRECISION = 3
const SBF_PVTCARTESIAN_2_0_VZ_PRECISION = 3
const SBF_PVTCARTESIAN_2_0_COG_PRECISION = 2
const SBF_PVTCARTESIAN_2_0_RXCLKBIAS_PRECISION = 6
const SBF_PVTCARTESIAN_2_0_RXCLKDRIFT_PRECISION = 3

/*--PVTCartesian_2_1_t : ----------------------------------------------------*/
/* Position, velocity, and time in Cartesian coordinates */
const SBF_PVTCARTESIAN_2_1_TOW_PRECISION = 3
const SBF_PVTCARTESIAN_2_1_X_PRECISION = 3
const SBF_PVTCARTESIAN_2_1_Y_PRECISION = 3
const SBF_PVTCARTESIAN_2_1_Z_PRECISION = 3
const SBF_PVTCARTESIAN_2_1_UNDULATION_PRECISION = 3
const SBF_PVTCARTESIAN_2_1_VX_PRECISION = 3
const SBF_PVTCARTESIAN_2_1_VY_PRECISION = 3
const SBF_PVTCARTESIAN_2_1_VZ_PRECISION = 3
const SBF_PVTCARTESIAN_2_1_COG_PRECISION = 2
const SBF_PVTCARTESIAN_2_1_RXCLKBIAS_PRECISION = 6
const SBF_PVTCARTESIAN_2_1_RXCLKDRIFT_PRECISION = 3

/*--PVTCartesian_2_2_t : ----------------------------------------------------*/
/* Position, velocity, and time in Cartesian coordinates */
const SBF_PVTCARTESIAN_2_2__PADDING_LENGTH = 1
const SBF_PVTCARTESIAN_2_2_TOW_PRECISION = 3
const SBF_PVTCARTESIAN_2_2_X_PRECISION = 3
const SBF_PVTCARTESIAN_2_2_Y_PRECISION = 3
const SBF_PVTCARTESIAN_2_2_Z_PRECISION = 3
const SBF_PVTCARTESIAN_2_2_UNDULATION_PRECISION = 3
const SBF_PVTCARTESIAN_2_2_VX_PRECISION = 3
const SBF_PVTCARTESIAN_2_2_VY_PRECISION = 3
const SBF_PVTCARTESIAN_2_2_VZ_PRECISION = 3
const SBF_PVTCARTESIAN_2_2_COG_PRECISION = 2
const SBF_PVTCARTESIAN_2_2_RXCLKBIAS_PRECISION = 6
const SBF_PVTCARTESIAN_2_2_RXCLKDRIFT_PRECISION = 3
const SBF_PVTCARTESIAN_2_2_HACCURACY_PRECISION = 2
const SBF_PVTCARTESIAN_2_2_VACCURACY_PRECISION = 2
const SBF_PVTCARTESIAN_2__PADDING_LENGTH = SBF_PVTCARTESIAN_2_2__PADDING_LENGTH

/*--PVTGeodetic_2_0_t : -----------------------------------------------------*/
/* Position, velocity, and time in geodetic coordinates */
const SBF_PVTGEODETIC_2_0__PADDING_LENGTH = 3
const SBF_PVTGEODETIC_2_0_TOW_PRECISION = 3
const SBF_PVTGEODETIC_2_0_LAT_PRECISION = 12
const SBF_PVTGEODETIC_2_0_LON_PRECISION = 12
const SBF_PVTGEODETIC_2_0_ALT_PRECISION = 3
const SBF_PVTGEODETIC_2_0_UNDULATION_PRECISION = 3
const SBF_PVTGEODETIC_2_0_VN_PRECISION = 3
const SBF_PVTGEODETIC_2_0_VE_PRECISION = 3
const SBF_PVTGEODETIC_2_0_VU_PRECISION = 3
const SBF_PVTGEODETIC_2_0_COG_PRECISION = 2
const SBF_PVTGEODETIC_2_0_RXCLKBIAS_PRECISION = 6
const SBF_PVTGEODETIC_2_0_RXCLKDRIFT_PRECISION = 3

/*--PVTGeodetic_2_1_t : -----------------------------------------------------*/
/* Position, velocity, and time in geodetic coordinates */
const SBF_PVTGEODETIC_2_1_TOW_PRECISION = 3
const SBF_PVTGEODETIC_2_1_LAT_PRECISION = 12
const SBF_PVTGEODETIC_2_1_LON_PRECISION = 12
const SBF_PVTGEODETIC_2_1_ALT_PRECISION = 3
const SBF_PVTGEODETIC_2_1_UNDULATION_PRECISION = 3
const SBF_PVTGEODETIC_2_1_VN_PRECISION = 3
const SBF_PVTGEODETIC_2_1_VE_PRECISION = 3
const SBF_PVTGEODETIC_2_1_VU_PRECISION = 3
const SBF_PVTGEODETIC_2_1_COG_PRECISION = 2
const SBF_PVTGEODETIC_2_1_RXCLKBIAS_PRECISION = 6
const SBF_PVTGEODETIC_2_1_RXCLKDRIFT_PRECISION = 3

/*--PVTGeodetic_2_2_t : -----------------------------------------------------*/
/* Position, velocity, and time in geodetic coordinates */
const SBF_PVTGEODETIC_2_2__PADDING_LENGTH = 1
const SBF_PVTGEODETIC_2_2_TOW_PRECISION = 3
const SBF_PVTGEODETIC_2_2_LAT_PRECISION = 12
const SBF_PVTGEODETIC_2_2_LON_PRECISION = 12
const SBF_PVTGEODETIC_2_2_ALT_PRECISION = 3
const SBF_PVTGEODETIC_2_2_UNDULATION_PRECISION = 3
const SBF_PVTGEODETIC_2_2_VN_PRECISION = 3
const SBF_PVTGEODETIC_2_2_VE_PRECISION = 3
const SBF_PVTGEODETIC_2_2_VU_PRECISION = 3
const SBF_PVTGEODETIC_2_2_COG_PRECISION = 2
const SBF_PVTGEODETIC_2_2_RXCLKBIAS_PRECISION = 6
const SBF_PVTGEODETIC_2_2_RXCLKDRIFT_PRECISION = 3
const SBF_PVTGEODETIC_2_2_HACCURACY_PRECISION = 2
const SBF_PVTGEODETIC_2_2_VACCURACY_PRECISION = 2
const SBF_PVTGEODETIC_2__PADDING_LENGTH = SBF_PVTGEODETIC_2_2__PADDING_LENGTH

/*--PVTGeodeticAuth_1_0_t : -------------------------------------------------*/
/* OSNMA-Authenticated Position, velocity, and time in geodetic coordinates */
const SBF_PVTGEODETICAUTH_1_0__PADDING_LENGTH = 3
const SBF_PVTGEODETICAUTH_1_0_TOW_PRECISION = 3
const SBF_PVTGEODETICAUTH_1_0_LATITUDE_PRECISION = 12
const SBF_PVTGEODETICAUTH_1_0_LONGITUDE_PRECISION = 12
const SBF_PVTGEODETICAUTH_1_0_HEIGHT_PRECISION = 3
const SBF_PVTGEODETICAUTH_1_0_UNDULATION_PRECISION = 3
const SBF_PVTGEODETICAUTH_1_0_VN_PRECISION = 3
const SBF_PVTGEODETICAUTH_1_0_VE_PRECISION = 3
const SBF_PVTGEODETICAUTH_1_0_VU_PRECISION = 3
const SBF_PVTGEODETICAUTH_1_0_COG_PRECISION = 2
const SBF_PVTGEODETICAUTH_1_0_RXCLKBIAS_PRECISION = 6
const SBF_PVTGEODETICAUTH_1_0_RXCLKDRIFT_PRECISION = 3

/*--PVTGeodeticAuth_1_1_t : -------------------------------------------------*/
/* OSNMA-Authenticated Position, velocity, and time in geodetic coordinates */
const SBF_PVTGEODETICAUTH_1_1_TOW_PRECISION = 3
const SBF_PVTGEODETICAUTH_1_1_LATITUDE_PRECISION = 12
const SBF_PVTGEODETICAUTH_1_1_LONGITUDE_PRECISION = 12
const SBF_PVTGEODETICAUTH_1_1_HEIGHT_PRECISION = 3
const SBF_PVTGEODETICAUTH_1_1_UNDULATION_PRECISION = 3
const SBF_PVTGEODETICAUTH_1_1_VN_PRECISION = 3
const SBF_PVTGEODETICAUTH_1_1_VE_PRECISION = 3
const SBF_PVTGEODETICAUTH_1_1_VU_PRECISION = 3
const SBF_PVTGEODETICAUTH_1_1_COG_PRECISION = 2
const SBF_PVTGEODETICAUTH_1_1_RXCLKBIAS_PRECISION = 6
const SBF_PVTGEODETICAUTH_1_1_RXCLKDRIFT_PRECISION = 3

/*--PVTGeodeticAuth_1_2_t : -------------------------------------------------*/
/* OSNMA-Authenticated Position, velocity, and time in geodetic coordinates */
const SBF_PVTGEODETICAUTH_1_2__PADDING_LENGTH = 1
const SBF_PVTGEODETICAUTH_1_2_TOW_PRECISION = 3
const SBF_PVTGEODETICAUTH_1_2_LATITUDE_PRECISION = 12
const SBF_PVTGEODETICAUTH_1_2_LONGITUDE_PRECISION = 12
const SBF_PVTGEODETICAUTH_1_2_HEIGHT_PRECISION = 3
const SBF_PVTGEODETICAUTH_1_2_UNDULATION_PRECISION = 3
const SBF_PVTGEODETICAUTH_1_2_VN_PRECISION = 3
const SBF_PVTGEODETICAUTH_1_2_VE_PRECISION = 3
const SBF_PVTGEODETICAUTH_1_2_VU_PRECISION = 3
const SBF_PVTGEODETICAUTH_1_2_COG_PRECISION = 2
const SBF_PVTGEODETICAUTH_1_2_RXCLKBIAS_PRECISION = 6
const SBF_PVTGEODETICAUTH_1_2_RXCLKDRIFT_PRECISION = 3
const SBF_PVTGEODETICAUTH_1_2_HACCURACY_PRECISION = 2
const SBF_PVTGEODETICAUTH_1_2_VACCURACY_PRECISION = 2
const SBF_PVTGEODETICAUTH_1__PADDING_LENGTH = SBF_PVTGEODETICAUTH_1_2__PADDING_LENGTH

/*--PosCovCartesian_1_0_t : -------------------------------------------------*/
/* Position covariance matrix (X,Y, Z) */
const SBF_POSCOVCARTESIAN_1_0_TOW_PRECISION = 3
const SBF_POSCOVCARTESIAN_1_0_COV_XX_PRECISION = 8
const SBF_POSCOVCARTESIAN_1_0_COV_YY_PRECISION = 8
const SBF_POSCOVCARTESIAN_1_0_COV_ZZ_PRECISION = 8
const SBF_POSCOVCARTESIAN_1_0_COV_TT_PRECISION = 8
const SBF_POSCOVCARTESIAN_1_0_COV_XY_PRECISION = 8
const SBF_POSCOVCARTESIAN_1_0_COV_XZ_PRECISION = 8
const SBF_POSCOVCARTESIAN_1_0_COV_XT_PRECISION = 8
const SBF_POSCOVCARTESIAN_1_0_COV_YZ_PRECISION = 8
const SBF_POSCOVCARTESIAN_1_0_COV_YT_PRECISION = 8
const SBF_POSCOVCARTESIAN_1_0_COV_ZT_PRECISION = 8

/*--PosCovGeodetic_1_0_t : --------------------------------------------------*/
/* Position covariance matrix (Lat, Lon, Alt) */
const SBF_POSCOVGEODETIC_1_0_TOW_PRECISION = 3
const SBF_POSCOVGEODETIC_1_0_COV_LATLAT_PRECISION = 8
const SBF_POSCOVGEODETIC_1_0_COV_LONLON_PRECISION = 8
const SBF_POSCOVGEODETIC_1_0_COV_ALTALT_PRECISION = 8
const SBF_POSCOVGEODETIC_1_0_COV_TT_PRECISION = 8
const SBF_POSCOVGEODETIC_1_0_COV_LATLON_PRECISION = 8
const SBF_POSCOVGEODETIC_1_0_COV_LATALT_PRECISION = 8
const SBF_POSCOVGEODETIC_1_0_COV_LATT_PRECISION = 8
const SBF_POSCOVGEODETIC_1_0_COV_LONALT_PRECISION = 8
const SBF_POSCOVGEODETIC_1_0_COV_LONT_PRECISION = 8
const SBF_POSCOVGEODETIC_1_0_COV_ALTT_PRECISION = 8

/*--VelCovCartesian_1_0_t : -------------------------------------------------*/
/* Velocity covariance matrix (X, Y, Z) */
const SBF_VELCOVCARTESIAN_1_0_TOW_PRECISION = 3
const SBF_VELCOVCARTESIAN_1_0_COV_VXVX_PRECISION = 3
const SBF_VELCOVCARTESIAN_1_0_COV_VYVY_PRECISION = 3
const SBF_VELCOVCARTESIAN_1_0_COV_VZVZ_PRECISION = 3
const SBF_VELCOVCARTESIAN_1_0_COV_DTDT_PRECISION = 3
const SBF_VELCOVCARTESIAN_1_0_COV_VXVY_PRECISION = 3
const SBF_VELCOVCARTESIAN_1_0_COV_VXVZ_PRECISION = 3
const SBF_VELCOVCARTESIAN_1_0_COV_VXDT_PRECISION = 3
const SBF_VELCOVCARTESIAN_1_0_COV_VYVZ_PRECISION = 3
const SBF_VELCOVCARTESIAN_1_0_COV_VYDT_PRECISION = 3
const SBF_VELCOVCARTESIAN_1_0_COV_VZDT_PRECISION = 3

/*--VelCovGeodetic_1_0_t : --------------------------------------------------*/
/* Velocity covariance matrix (North, East, Up) */
const SBF_VELCOVGEODETIC_1_0_TOW_PRECISION = 3
const SBF_VELCOVGEODETIC_1_0_COV_VNVN_PRECISION = 3
const SBF_VELCOVGEODETIC_1_0_COV_VEVE_PRECISION = 3
const SBF_VELCOVGEODETIC_1_0_COV_VUVU_PRECISION = 3
const SBF_VELCOVGEODETIC_1_0_COV_DTDT_PRECISION = 3
const SBF_VELCOVGEODETIC_1_0_COV_VNVE_PRECISION = 3
const SBF_VELCOVGEODETIC_1_0_COV_VNVU_PRECISION = 3
const SBF_VELCOVGEODETIC_1_0_COV_VNDT_PRECISION = 3
const SBF_VELCOVGEODETIC_1_0_COV_VEVU_PRECISION = 3
const SBF_VELCOVGEODETIC_1_0_COV_VEDT_PRECISION = 3
const SBF_VELCOVGEODETIC_1_0_COV_VUDT_PRECISION = 3

/*--DOP_2_0_t : -------------------------------------------------------------*/
/* Dilution of precision */
const SBF_DOP_2_0_TOW_PRECISION = 3
const SBF_DOP_2_0_HPL_PRECISION = 3
const SBF_DOP_2_0_VPL_PRECISION = 3

/*--PosCart_1_0_t : ---------------------------------------------------------*/
/* Position, variance and baseline in Cartesian coordinates */
const SBF_POSCART_1_0_TOW_PRECISION = 3
const SBF_POSCART_1_0_X_PRECISION = 3
const SBF_POSCART_1_0_Y_PRECISION = 3
const SBF_POSCART_1_0_Z_PRECISION = 3
const SBF_POSCART_1_0_BASE2ROVERX_PRECISION = 3
const SBF_POSCART_1_0_BASE2ROVERY_PRECISION = 3
const SBF_POSCART_1_0_BASE2ROVERZ_PRECISION = 3
const SBF_POSCART_1_0_COV_XX_PRECISION = 8
const SBF_POSCART_1_0_COV_YY_PRECISION = 8
const SBF_POSCART_1_0_COV_ZZ_PRECISION = 8
const SBF_POSCART_1_0_COV_XY_PRECISION = 8
const SBF_POSCART_1_0_COV_XZ_PRECISION = 8
const SBF_POSCART_1_0_COV_YZ_PRECISION = 8

/*--PosLocal_1_0_t : --------------------------------------------------------*/
/* Position in a local datum */
const SBF_POSLOCAL_1_0__PADDING_LENGTH = 3
const SBF_POSLOCAL_1_0_TOW_PRECISION = 3
const SBF_POSLOCAL_1_0_LAT_PRECISION = 12
const SBF_POSLOCAL_1_0_LON_PRECISION = 12
const SBF_POSLOCAL_1_0_ALT_PRECISION = 3
const SBF_POSLOCAL_1__PADDING_LENGTH = SBF_POSLOCAL_1_0__PADDING_LENGTH

/*--PosProjected_1_0_t : ----------------------------------------------------*/
/* Plane grid coordinates */
const SBF_POSPROJECTED_1_0__PADDING_LENGTH = 3
const SBF_POSPROJECTED_1_0_TOW_PRECISION = 3
const SBF_POSPROJECTED_1_0_NORTHING_PRECISION = 3
const SBF_POSPROJECTED_1_0_EASTING_PRECISION = 3
const SBF_POSPROJECTED_1_0_ALT_PRECISION = 3
const SBF_POSPROJECTED_1__PADDING_LENGTH = SBF_POSPROJECTED_1_0__PADDING_LENGTH

/*--PVTSatCartesian_1_0_t : -------------------------------------------------*/
/* Satellite positions */
const SBF_SATPOS_1_0_X_PRECISION = 3
const SBF_SATPOS_1_0_Y_PRECISION = 3
const SBF_SATPOS_1_0_Z_PRECISION = 3
const SBF_SATPOS_1_0_VX_PRECISION = 3
const SBF_SATPOS_1_0_VY_PRECISION = 3
const SBF_SATPOS_1_0_VZ_PRECISION = 3
const SBF_PVTSATCARTESIAN_1_0_SATPOS_LENGTH = 72
const SBF_PVTSATCARTESIAN_1_0_TOW_PRECISION = 3

/*--PVTSatCartesian_1_1_t : -------------------------------------------------*/
/* Satellite positions */
const SBF_SATPOS_1_1__PADDING_LENGTH = 1
const SBF_SATPOS_1_1_X_PRECISION = 3
const SBF_SATPOS_1_1_Y_PRECISION = 3
const SBF_SATPOS_1_1_Z_PRECISION = 3
const SBF_SATPOS_1_1_VX_PRECISION = 3
const SBF_SATPOS_1_1_VY_PRECISION = 3
const SBF_SATPOS_1_1_VZ_PRECISION = 3
const SBF_PVTSATCARTESIAN_1_1_SATPOS_LENGTH = 72
const SBF_PVTSATCARTESIAN_1_1_TOW_PRECISION = 3
const SBF_SATPOS_1__PADDING_LENGTH = SBF_SATPOS_1_1__PADDING_LENGTH
const SBF_PVTSATCARTESIAN_1_SATPOS_LENGTH = SBF_PVTSATCARTESIAN_1_1_SATPOS_LENGTH

/*--PVTResiduals_2_0_t : ----------------------------------------------------*/
/* Measurement residuals */
const SBF_RESIDUALINFOCODE_2_0_RESIDUAL_PRECISION = 3
const SBF_RESIDUALINFOPHASE_2_0_RESIDUAL_PRECISION = 3
const SBF_RESIDUALINFODOPPLER_2_0_RESIDUAL_PRECISION = 3
const SBF_PVTRESIDUALS_2_0_TOW_PRECISION = 3

/*--PVTResiduals_2_1_t : ----------------------------------------------------*/
/* Measurement residuals */
const SBF_RESIDUALINFOCODE_2_1_RESIDUAL_PRECISION = 3
const SBF_RESIDUALINFOPHASE_2_1_RESIDUAL_PRECISION = 3
const SBF_RESIDUALINFODOPPLER_2_1_RESIDUAL_PRECISION = 3
const SBF_PVTRESIDUALS_2_1_TOW_PRECISION = 3
const SBF_PVTRESIDUALS_2_DATA_LENGTH = SBF_PVTRESIDUALS_2_1_DATA_LENGTH

/*--RAIMStatistics_2_0_t : --------------------------------------------------*/
/* Integrity statistics */
const SBF_RAIMSTATISTICS_2_0__PADDING_LENGTH = 2
const SBF_RAIMSTATISTICS_2_0_TOW_PRECISION = 3
const SBF_RAIMSTATISTICS_2_0_POSITIONHERL_PRECISION = 3
const SBF_RAIMSTATISTICS_2_0_POSITIONVERL_PRECISION = 3
const SBF_RAIMSTATISTICS_2_0_VELOCITYHERL_PRECISION = 3
const SBF_RAIMSTATISTICS_2_0_VELOCITYVERL_PRECISION = 3
const SBF_RAIMSTATISTICS_2__PADDING_LENGTH = SBF_RAIMSTATISTICS_2_0__PADDING_LENGTH

/*--GEOCorrections_1_0_t : --------------------------------------------------*/
/* Orbit, Clock and pseudoranges SBAS corrections */
const SBF_SATCORR_1_0_CORRAGEFC_PRECISION = 3
const SBF_SATCORR_1_0_CORRAGELT_PRECISION = 3
const SBF_SATCORR_1_0_IONOPPLAT_PRECISION = 3
const SBF_SATCORR_1_0_IONOPPLON_PRECISION = 3
const SBF_SATCORR_1_0_SLANTIONO_PRECISION = 3
const SBF_SATCORR_1_0_CORRAGEIONO_PRECISION = 3
const SBF_SATCORR_1_0_VARFLT_PRECISION = 3
const SBF_SATCORR_1_0_VARUIRE_PRECISION = 3
const SBF_SATCORR_1_0_VARAIR_PRECISION = 3
const SBF_SATCORR_1_0_VARTROPO_PRECISION = 3
const SBF_GEOCORRECTIONS_1_0_GEOCORRCHANNEL_LENGTH = 72
const SBF_GEOCORRECTIONS_1_0_TOW_PRECISION = 3
const SBF_GEOCORRECTIONS_1_GEOCORRCHANNEL_LENGTH = SBF_GEOCORRECTIONS_1_0_GEOCORRCHANNEL_LENGTH

/*--BaseVectorCart_1_0_t : --------------------------------------------------*/
/* XYZ relative position and velocity with respect to base(s) */
const SBF_VECTORINFOCART_1_0_DELTAX_PRECISION = 3
const SBF_VECTORINFOCART_1_0_DELTAY_PRECISION = 3
const SBF_VECTORINFOCART_1_0_DELTAZ_PRECISION = 3
const SBF_VECTORINFOCART_1_0_DELTAXVELOCITY_PRECISION = 3
const SBF_VECTORINFOCART_1_0_DELTAYVELOCITY_PRECISION = 3
const SBF_VECTORINFOCART_1_0_DELTAZVELOCITY_PRECISION = 3
const SBF_VECTORINFOCART_1_0_AZIMUTH_PRECISION = 2
const SBF_VECTORINFOCART_1_0_ELEVATION_PRECISION = 2
const SBF_BASEVECTORCART_1_0_VECTORINFO_LENGTH = 30
const SBF_BASEVECTORCART_1_0_TOW_PRECISION = 3
const SBF_BASEVECTORCART_1_VECTORINFO_LENGTH = SBF_BASEVECTORCART_1_0_VECTORINFO_LENGTH

/*--BaseVectorGeod_1_0_t : --------------------------------------------------*/
/* ENU relative position and velocity with respect to base(s) */
const SBF_VECTORINFOGEOD_1_0_DELTAEAST_PRECISION = 3
const SBF_VECTORINFOGEOD_1_0_DELTANORTH_PRECISION = 3
const SBF_VECTORINFOGEOD_1_0_DELTAUP_PRECISION = 3
const SBF_VECTORINFOGEOD_1_0_DELTAEASTVELOCITY_PRECISION = 3
const SBF_VECTORINFOGEOD_1_0_DELTANORTHVELOCITY_PRECISION = 3
const SBF_VECTORINFOGEOD_1_0_DELTAUPVELOCITY_PRECISION = 3
const SBF_VECTORINFOGEOD_1_0_AZIMUTH_PRECISION = 2
const SBF_VECTORINFOGEOD_1_0_ELEVATION_PRECISION = 2
const SBF_BASEVECTORGEOD_1_0_VECTORINFO_LENGTH = 30
const SBF_BASEVECTORGEOD_1_0_TOW_PRECISION = 3
const SBF_BASEVECTORGEOD_1_VECTORINFO_LENGTH = SBF_BASEVECTORGEOD_1_0_VECTORINFO_LENGTH

/*--Ambiguities_1_0_t : -----------------------------------------------------*/
/* Carrier phase ambiguity states */
const SBF_AMBIGUITIES_1_0_AMBIGUITIES_LENGTH = 72
const SBF_AMBIGUITIES_1_0_TOW_PRECISION = 3
const SBF_AMBIGUITIES_1_AMBIGUITIES_LENGTH = SBF_AMBIGUITIES_1_0_AMBIGUITIES_LENGTH

/*--EndOfPVT_1_0_t : --------------------------------------------------------*/
/* PVT epoch marker */
const SBF_ENDOFPVT_1_0_TOW_PRECISION = 3

/*--BaseLine_1_0_t : --------------------------------------------------------*/
/* Base-rover vector (deprecated block - not to be used) */
const SBF_BASELINE_1_0_TOW_PRECISION = 3
const SBF_BASELINE_1_0_EAST_PRECISION = 3
const SBF_BASELINE_1_0_NORTH_PRECISION = 3
const SBF_BASELINE_1_0_UP_PRECISION = 3

/*==INS/GNSS Integrated Blocks===============================================*/

/*--IntPVCart_1_0_t : -------------------------------------------------------*/
/* Integrated PV in Cartesian coordinates */
const SBF_INTPVCART_1_0_TOW_PRECISION = 3
const SBF_INTPVCART_1_0_X_PRECISION = 3
const SBF_INTPVCART_1_0_Y_PRECISION = 3
const SBF_INTPVCART_1_0_Z_PRECISION = 3
const SBF_INTPVCART_1_0_VX_PRECISION = 3
const SBF_INTPVCART_1_0_VY_PRECISION = 3
const SBF_INTPVCART_1_0_VZ_PRECISION = 3
const SBF_INTPVCART_1_0_COG_PRECISION = 2

/*--IntPVGeod_1_0_t : -------------------------------------------------------*/
/* Integrated PV in Geodetic coordinates */
const SBF_INTPVGEOD_1_0_TOW_PRECISION = 3
const SBF_INTPVGEOD_1_0_LAT_PRECISION = 12
const SBF_INTPVGEOD_1_0_LON_PRECISION = 12
const SBF_INTPVGEOD_1_0_ALT_PRECISION = 3
const SBF_INTPVGEOD_1_0_VN_PRECISION = 3
const SBF_INTPVGEOD_1_0_VE_PRECISION = 3
const SBF_INTPVGEOD_1_0_VU_PRECISION = 3
const SBF_INTPVGEOD_1_0_COG_PRECISION = 2

/*--IntPosCovCart_1_0_t : ---------------------------------------------------*/
/* Integrated position covariance matrix (X, Y, Z) */
const SBF_INTPOSCOVCART_1_0_TOW_PRECISION = 3
const SBF_INTPOSCOVCART_1_0_COV_XX_PRECISION = 8
const SBF_INTPOSCOVCART_1_0_COV_YY_PRECISION = 8
const SBF_INTPOSCOVCART_1_0_COV_ZZ_PRECISION = 8
const SBF_INTPOSCOVCART_1_0_COV_XY_PRECISION = 8
const SBF_INTPOSCOVCART_1_0_COV_XZ_PRECISION = 8
const SBF_INTPOSCOVCART_1_0_COV_YZ_PRECISION = 8

/*--IntVelCovCart_1_0_t : ---------------------------------------------------*/
/* Integrated velocity covariance matrix (X, Y, Z) */
const SBF_INTVELCOVCART_1_0_TOW_PRECISION = 3
const SBF_INTVELCOVCART_1_0_COV_VXVX_PRECISION = 3
const SBF_INTVELCOVCART_1_0_COV_VYVY_PRECISION = 3
const SBF_INTVELCOVCART_1_0_COV_VZVZ_PRECISION = 3
const SBF_INTVELCOVCART_1_0_COV_VXVY_PRECISION = 3
const SBF_INTVELCOVCART_1_0_COV_VXVZ_PRECISION = 3
const SBF_INTVELCOVCART_1_0_COV_VYVZ_PRECISION = 3

/*--IntPosCovGeod_1_0_t : ---------------------------------------------------*/
/* Integrated position covariance matrix (Lat, Lon, Alt) */
const SBF_INTPOSCOVGEOD_1_0_TOW_PRECISION = 3
const SBF_INTPOSCOVGEOD_1_0_COV_LATLAT_PRECISION = 8
const SBF_INTPOSCOVGEOD_1_0_COV_LONLON_PRECISION = 8
const SBF_INTPOSCOVGEOD_1_0_COV_ALTALT_PRECISION = 8
const SBF_INTPOSCOVGEOD_1_0_COV_LATLON_PRECISION = 8
const SBF_INTPOSCOVGEOD_1_0_COV_LATALT_PRECISION = 8
const SBF_INTPOSCOVGEOD_1_0_COV_LONALT_PRECISION = 8

/*--IntVelCovGeod_1_0_t : ---------------------------------------------------*/
/* Integrated velocity covariance matrix (North, East, Up) */
const SBF_INTVELCOVGEOD_1_0_TOW_PRECISION = 3
const SBF_INTVELCOVGEOD_1_0_COV_VNVN_PRECISION = 3
const SBF_INTVELCOVGEOD_1_0_COV_VEVE_PRECISION = 3
const SBF_INTVELCOVGEOD_1_0_COV_VUVU_PRECISION = 3
const SBF_INTVELCOVGEOD_1_0_COV_VNVE_PRECISION = 3
const SBF_INTVELCOVGEOD_1_0_COV_VNVU_PRECISION = 3
const SBF_INTVELCOVGEOD_1_0_COV_VEVU_PRECISION = 3

/*--IntAttEuler_1_0_t : -----------------------------------------------------*/
/* Integrated attitude in Euler angles */
const SBF_INTATTEULER_1_0_TOW_PRECISION = 3
const SBF_INTATTEULER_1_0_HEADING_PRECISION = 6
const SBF_INTATTEULER_1_0_PITCH_PRECISION = 6
const SBF_INTATTEULER_1_0_ROLL_PRECISION = 6
const SBF_INTATTEULER_1_0_PITCHDOT_PRECISION = 3
const SBF_INTATTEULER_1_0_ROLLDOT_PRECISION = 3
const SBF_INTATTEULER_1_0_HEADINGDOT_PRECISION = 3

/*--IntAttEuler_1_1_t : -----------------------------------------------------*/
/* Integrated attitude in Euler angles */
const SBF_INTATTEULER_1_1_TOW_PRECISION = 3
const SBF_INTATTEULER_1_1_HEADING_PRECISION = 6
const SBF_INTATTEULER_1_1_PITCH_PRECISION = 6
const SBF_INTATTEULER_1_1_ROLL_PRECISION = 6
const SBF_INTATTEULER_1_1_PITCHDOT_PRECISION = 3
const SBF_INTATTEULER_1_1_ROLLDOT_PRECISION = 3
const SBF_INTATTEULER_1_1_HEADINGDOT_PRECISION = 3

/*--IntAttCovEuler_1_0_t : --------------------------------------------------*/
/* Integrated attitude covariance matrix of Euler angles */
const SBF_INTATTCOVEULER_1_0_TOW_PRECISION = 3
const SBF_INTATTCOVEULER_1_0_COV_HEADHEAD_PRECISION = 3
const SBF_INTATTCOVEULER_1_0_COV_PITCHPITCH_PRECISION = 3
const SBF_INTATTCOVEULER_1_0_COV_ROLLROLL_PRECISION = 3
const SBF_INTATTCOVEULER_1_0_COV_HEADPITCH_PRECISION = 3
const SBF_INTATTCOVEULER_1_0_COV_HEADROLL_PRECISION = 3
const SBF_INTATTCOVEULER_1_0_COV_PITCHROLL_PRECISION = 3

/*--IntPVAAGeod_1_0_t : -----------------------------------------------------*/
/* Integrated position, velocity, acceleration and attitude */
const SBF_INTPVAAGEOD_1_0_TOW_PRECISION = 3

/*--INSNavCart_1_0_t : ------------------------------------------------------*/
/* INS solution in Cartesian coordinates */

const SBF_INSNAVCARTPOSSTDDEV_1_0_XSTDDEV_PRECISION = 3
const SBF_INSNAVCARTPOSSTDDEV_1_0_YSTDDEV_PRECISION = 3
const SBF_INSNAVCARTPOSSTDDEV_1_0_ZSTDDEV_PRECISION = 3
const SBF_INSNAVCARTPOSCOV_1_0_XYCOV_PRECISION = 3
const SBF_INSNAVCARTPOSCOV_1_0_XZCOV_PRECISION = 3
const SBF_INSNAVCARTPOSCOV_1_0_YZCOV_PRECISION = 3
const SBF_INSNAVCARTATT_1_0_HEADING_PRECISION = 6
const SBF_INSNAVCARTATT_1_0_PITCH_PRECISION = 6
const SBF_INSNAVCARTATT_1_0_ROLL_PRECISION = 6
const SBF_INSNAVCARTATTSTDDEV_1_0_HEADINGSTDDEV_PRECISION = 6
const SBF_INSNAVCARTATTSTDDEV_1_0_PITCHSTDDEV_PRECISION = 6
const SBF_INSNAVCARTATTSTDDEV_1_0_ROLLSTDDEV_PRECISION = 6
const SBF_INSNAVCARTATTCOV_1_0_HEADINGPITCHCOV_PRECISION = 6
const SBF_INSNAVCARTATTCOV_1_0_HEADINGROLLCOV_PRECISION = 6
const SBF_INSNAVCARTATTCOV_1_0_PITCHROLLCOV_PRECISION = 6
const SBF_INSNAVCARTVEL_1_0_VX_PRECISION = 3
const SBF_INSNAVCARTVEL_1_0_VY_PRECISION = 3
const SBF_INSNAVCARTVEL_1_0_VZ_PRECISION = 3
const SBF_INSNAVCARTVELSTDDEV_1_0_VXSTDDEV_PRECISION = 3
const SBF_INSNAVCARTVELSTDDEV_1_0_VYSTDDEV_PRECISION = 3
const SBF_INSNAVCARTVELSTDDEV_1_0_VZSTDDEV_PRECISION = 3
const SBF_INSNAVCARTVELCOV_1_0_VXVYCOV_PRECISION = 3
const SBF_INSNAVCARTVELCOV_1_0_VXVZCOV_PRECISION = 3
const SBF_INSNAVCARTVELCOV_1_0_VYVZCOV_PRECISION = 3
const SBF_INSNAVCART_1_0_INSNAVCARTDATA_LENGTH = 16
const SBF_INSNAVCART_1_0_TOW_PRECISION = 3
const SBF_INSNAVCART_1_0_X_PRECISION = 3
const SBF_INSNAVCART_1_0_Y_PRECISION = 3
const SBF_INSNAVCART_1_0_Z_PRECISION = 3
const SBF_INSNAVCART_1_0_ACCURACY_PRECISION = 2
const SBF_INSNAVCART_1_INSNAVCARTDATA_LENGTH = SBF_INSNAVCART_1_0_INSNAVCARTDATA_LENGTH

/*--INSNavGeod_1_0_t : ------------------------------------------------------*/
/* INS solution in Geodetic coordinates */
const SBF_INSNAVGEODPOSSTDDEV_1_0_LATITUDESTDDEV_PRECISION = 3
const SBF_INSNAVGEODPOSSTDDEV_1_0_LONGITUDESTDDEV_PRECISION = 3
const SBF_INSNAVGEODPOSSTDDEV_1_0_HEIGHTSTDDEV_PRECISION = 3
const SBF_INSNAVGEODPOSCOV_1_0_LATITUDELONGITUDECOV_PRECISION = 3
const SBF_INSNAVGEODPOSCOV_1_0_LATITUDEHEIGHTCOV_PRECISION = 3
const SBF_INSNAVGEODPOSCOV_1_0_LONGITUDEHEIGHTCOV_PRECISION = 3
const SBF_INSNAVGEODATT_1_0_HEADING_PRECISION = 6
const SBF_INSNAVGEODATT_1_0_PITCH_PRECISION = 6
const SBF_INSNAVGEODATT_1_0_ROLL_PRECISION = 6
const SBF_INSNAVGEODATTSTDDEV_1_0_HEADINGSTDDEV_PRECISION = 6
const SBF_INSNAVGEODATTSTDDEV_1_0_PITCHSTDDEV_PRECISION = 6
const SBF_INSNAVGEODATTSTDDEV_1_0_ROLLSTDDEV_PRECISION = 6
const SBF_INSNAVGEODATTCOV_1_0_HEADINGPITCHCOV_PRECISION = 6
const SBF_INSNAVGEODATTCOV_1_0_HEADINGROLLCOV_PRECISION = 6
const SBF_INSNAVGEODATTCOV_1_0_PITCHROLLCOV_PRECISION = 6
const SBF_INSNAVGEODVEL_1_0_VE_PRECISION = 3
const SBF_INSNAVGEODVEL_1_0_VN_PRECISION = 3
const SBF_INSNAVGEODVEL_1_0_VU_PRECISION = 3
const SBF_INSNAVGEODVELSTDDEV_1_0_VESTDDEV_PRECISION = 3
const SBF_INSNAVGEODVELSTDDEV_1_0_VNSTDDEV_PRECISION = 3
const SBF_INSNAVGEODVELSTDDEV_1_0_VUSTDDEV_PRECISION = 3
const SBF_INSNAVGEODVELCOV_1_0_VEVNCOV_PRECISION = 3
const SBF_INSNAVGEODVELCOV_1_0_VEVUCOV_PRECISION = 3
const SBF_INSNAVGEODVELCOV_1_0_VNVUCOV_PRECISION = 3
const SBF_INSNAVGEOD_1_0_INSNAVGEODDATA_LENGTH = 16
const SBF_INSNAVGEOD_1_0_TOW_PRECISION = 3
const SBF_INSNAVGEOD_1_0_LATITUDE_PRECISION = 12
const SBF_INSNAVGEOD_1_0_LONGITUDE_PRECISION = 12
const SBF_INSNAVGEOD_1_0_HEIGHT_PRECISION = 3
const SBF_INSNAVGEOD_1_0_UNDULATION_PRECISION = 3
const SBF_INSNAVGEOD_1_0_ACCURACY_PRECISION = 2
const SBF_INSNAVGEOD_1_INSNAVGEODDATA_LENGTH = SBF_INSNAVGEOD_1_0_INSNAVGEODDATA_LENGTH

/*--IMUBias_1_0_t : ---------------------------------------------------------*/
/* Estimated parameters of the IMU, such as the IMU biases and their standard deviation */
const SBF_IMUBIAS_1_0_TOW_PRECISION = 3
const SBF_IMUBIAS_1_0_XACCBIAS_PRECISION = 3
const SBF_IMUBIAS_1_0_YACCBIAS_PRECISION = 3
const SBF_IMUBIAS_1_0_ZACCBIAS_PRECISION = 3
const SBF_IMUBIAS_1_0_XGYROBIAS_PRECISION = 3
const SBF_IMUBIAS_1_0_YGYROBIAS_PRECISION = 3
const SBF_IMUBIAS_1_0_ZGYROBIAS_PRECISION = 3
const SBF_IMUBIAS_1_0_STDXACCBIAS_PRECISION = 3
const SBF_IMUBIAS_1_0_STDYACCBIAS_PRECISION = 3
const SBF_IMUBIAS_1_0_STDZACCBIAS_PRECISION = 3
const SBF_IMUBIAS_1_0_STDXGYROBIAS_PRECISION = 3
const SBF_IMUBIAS_1_0_STDYGYROBIAS_PRECISION = 3
const SBF_IMUBIAS_1_0_STDZGYROBIAS_PRECISION = 3

/*==GNSS Attitude Blocks=====================================================*/

/*--AttEuler_1_0_t : --------------------------------------------------------*/
/* GNSS attitude expressed as Euler angles */
const SBF_ATTEULER_1_0_TOW_PRECISION = 3
const SBF_ATTEULER_1_0_HEADING_PRECISION = 6
const SBF_ATTEULER_1_0_PITCH_PRECISION = 6
const SBF_ATTEULER_1_0_ROLL_PRECISION = 6
const SBF_ATTEULER_1_0_PITCHDOT_PRECISION = 3
const SBF_ATTEULER_1_0_ROLLDOT_PRECISION = 3
const SBF_ATTEULER_1_0_HEADINGDOT_PRECISION = 3

/*--AttCovEuler_1_0_t : -----------------------------------------------------*/
/* Covariance matrix of attitude */
const SBF_ATTCOVEULER_1_0_TOW_PRECISION = 3
const SBF_ATTCOVEULER_1_0_COV_HEADHEAD_PRECISION = 3
const SBF_ATTCOVEULER_1_0_COV_PITCHPITCH_PRECISION = 3
const SBF_ATTCOVEULER_1_0_COV_ROLLROLL_PRECISION = 3
const SBF_ATTCOVEULER_1_0_COV_HEADPITCH_PRECISION = 3
const SBF_ATTCOVEULER_1_0_COV_HEADROLL_PRECISION = 3
const SBF_ATTCOVEULER_1_0_COV_PITCHROLL_PRECISION = 3

/*--AuxAntPositions_1_0_t : -------------------------------------------------*/
/* Relative position and velocity estimates of auxiliary antennas */
const SBF_AUXANTPOSITIONSUB_1_0_DELTAEAST_PRECISION = 3
const SBF_AUXANTPOSITIONSUB_1_0_DELTANORTH_PRECISION = 3
const SBF_AUXANTPOSITIONSUB_1_0_DELTAUP_PRECISION = 3
const SBF_AUXANTPOSITIONSUB_1_0_EASTVELOCITY_PRECISION = 3
const SBF_AUXANTPOSITIONSUB_1_0_NORTHVELOCITY_PRECISION = 3
const SBF_AUXANTPOSITIONSUB_1_0_UPVELOCITY_PRECISION = 3
const SBF_AUXANTPOSITIONS_1_0_AUXANTPOSITIONS_LENGTH = 4
const SBF_AUXANTPOSITIONS_1_0_TOW_PRECISION = 3
const SBF_AUXANTPOSITIONS_1_AUXANTPOSITIONS_LENGTH = SBF_AUXANTPOSITIONS_1_0_AUXANTPOSITIONS_LENGTH

/*--EndOfAtt_1_0_t : --------------------------------------------------------*/
/* GNSS attitude epoch marker */
const SBF_ENDOFATT_1_0_TOW_PRECISION = 3

/*--AttQuat_1_0_t : ---------------------------------------------------------*/
const SBF_ATTQUAT_1_0_TOW_PRECISION = 3
const SBF_ATTQUAT_1_0_Q1_PRECISION = 3
const SBF_ATTQUAT_1_0_Q2_PRECISION = 3
const SBF_ATTQUAT_1_0_Q3_PRECISION = 3
const SBF_ATTQUAT_1_0_Q4_PRECISION = 3
const SBF_ATTQUAT_1_0_PITCHDOT_PRECISION = 3
const SBF_ATTQUAT_1_0_ROLLDOT_PRECISION = 3
const SBF_ATTQUAT_1_0_HEADINGDOT_PRECISION = 3

/*--AttCovQuat_1_0_t : ------------------------------------------------------*/
const SBF_ATTCOVQUAT_1_0_TOW_PRECISION = 3
const SBF_ATTCOVQUAT_1_0_COV_Q1Q1_PRECISION = 3
const SBF_ATTCOVQUAT_1_0_COV_Q2Q2_PRECISION = 3
const SBF_ATTCOVQUAT_1_0_COV_Q3Q3_PRECISION = 3
const SBF_ATTCOVQUAT_1_0_COV_Q4Q4_PRECISION = 3
const SBF_ATTCOVQUAT_1_0_COV_Q1Q2_PRECISION = 3
const SBF_ATTCOVQUAT_1_0_COV_Q1Q3_PRECISION = 3
const SBF_ATTCOVQUAT_1_0_COV_Q1Q4_PRECISION = 3
const SBF_ATTCOVQUAT_1_0_COV_Q2Q3_PRECISION = 3
const SBF_ATTCOVQUAT_1_0_COV_Q2Q4_PRECISION = 3
const SBF_ATTCOVQUAT_1_0_COV_Q3Q4_PRECISION = 3

/*==Receiver Time Blocks=====================================================*/

/*--ReceiverTime_1_0_t : ----------------------------------------------------*/
/* Current receiver and UTC time */
const SBF_RECEIVERTIME_1_0__PADDING_LENGTH = 2
const SBF_RECEIVERTIME_1_0_TOW_PRECISION = 3
const SBF_RECEIVERTIME_1__PADDING_LENGTH = SBF_RECEIVERTIME_1_0__PADDING_LENGTH

/*--xPPSOffset_1_0_t : ------------------------------------------------------*/
/* Offset of the xPPS pulse with respect to GNSS time */
const SBF_PPSDATA_1_0_TOW_PRECISION = 3

/*--SysTimeOffset_1_0_t : ---------------------------------------------------*/
/* Time offset between different constellations */
const SBF_TIMEOFFSETSUB_1_0__PADDING_LENGTH = 3
const SBF_TIMEOFFSETSUB_1_0_OFFSET_PRECISION = 3
const SBF_SYSTIMEOFFSET_1_0_TIMEOFFSET_LENGTH = 11
const SBF_SYSTIMEOFFSET_1_0_TOW_PRECISION = 3

/*--SysTimeOffset_1_1_t : ---------------------------------------------------*/
/* Time offset between different constellations */
const SBF_TIMEOFFSETSUB_1_1_OFFSET_PRECISION = 3
const SBF_SYSTIMEOFFSET_1_1_TIMEOFFSET_LENGTH = 11
const SBF_SYSTIMEOFFSET_1_1_TOW_PRECISION = 3
const SBF_SYSTIMEOFFSET_1_TIMEOFFSET_LENGTH = SBF_SYSTIMEOFFSET_1_1_TIMEOFFSET_LENGTH

/*==External Event Blocks====================================================*/

/*--ExtEvent_1_0_t : --------------------------------------------------------*/
/* Time at the instant of an external event */
const SBF_TIMERDATA_1_0_TOW_PRECISION = 3
const SBF_TIMERDATA_1_0_RXCLKBIAS_PRECISION = 9

/*--ExtEvent_1_1_t : --------------------------------------------------------*/
/* Time at the instant of an external event */
const SBF_TIMERDATA_1_1_TOW_PRECISION = 3
const SBF_TIMERDATA_1_1_RXCLKBIAS_PRECISION = 9

/*--ExtEventPVTCartesian_1_0_t : --------------------------------------------*/
/* Cartesian position at the instant of an event */
const SBF_EXTEVENTPVTCARTESIAN_1_0__PADDING_LENGTH = 3
const SBF_EXTEVENTPVTCARTESIAN_1_0_TOW_PRECISION = 3
const SBF_EXTEVENTPVTCARTESIAN_1_0_X_PRECISION = 3
const SBF_EXTEVENTPVTCARTESIAN_1_0_Y_PRECISION = 3
const SBF_EXTEVENTPVTCARTESIAN_1_0_Z_PRECISION = 3
const SBF_EXTEVENTPVTCARTESIAN_1_0_UNDULATION_PRECISION = 3
const SBF_EXTEVENTPVTCARTESIAN_1_0_VX_PRECISION = 3
const SBF_EXTEVENTPVTCARTESIAN_1_0_VY_PRECISION = 3
const SBF_EXTEVENTPVTCARTESIAN_1_0_VZ_PRECISION = 3
const SBF_EXTEVENTPVTCARTESIAN_1_0_COG_PRECISION = 2
const SBF_EXTEVENTPVTCARTESIAN_1_0_RXCLKBIAS_PRECISION = 6
const SBF_EXTEVENTPVTCARTESIAN_1_0_RXCLKDRIFT_PRECISION = 3

/*--ExtEventPVTCartesian_1_1_t : --------------------------------------------*/
/* Cartesian position at the instant of an event */
const SBF_EXTEVENTPVTCARTESIAN_1_1_TOW_PRECISION = 3
const SBF_EXTEVENTPVTCARTESIAN_1_1_X_PRECISION = 3
const SBF_EXTEVENTPVTCARTESIAN_1_1_Y_PRECISION = 3
const SBF_EXTEVENTPVTCARTESIAN_1_1_Z_PRECISION = 3
const SBF_EXTEVENTPVTCARTESIAN_1_1_UNDULATION_PRECISION = 3
const SBF_EXTEVENTPVTCARTESIAN_1_1_VX_PRECISION = 3
const SBF_EXTEVENTPVTCARTESIAN_1_1_VY_PRECISION = 3
const SBF_EXTEVENTPVTCARTESIAN_1_1_VZ_PRECISION = 3
const SBF_EXTEVENTPVTCARTESIAN_1_1_COG_PRECISION = 2
const SBF_EXTEVENTPVTCARTESIAN_1_1_RXCLKBIAS_PRECISION = 6
const SBF_EXTEVENTPVTCARTESIAN_1_1_RXCLKDRIFT_PRECISION = 3

/*--ExtEventPVTCartesian_1_2_t : --------------------------------------------*/
/* Cartesian position at the instant of an event */
const SBF_EXTEVENTPVTCARTESIAN_1_2__PADDING_LENGTH = 1
const SBF_EXTEVENTPVTCARTESIAN_1_2_TOW_PRECISION = 3
const SBF_EXTEVENTPVTCARTESIAN_1_2_X_PRECISION = 3
const SBF_EXTEVENTPVTCARTESIAN_1_2_Y_PRECISION = 3
const SBF_EXTEVENTPVTCARTESIAN_1_2_Z_PRECISION = 3
const SBF_EXTEVENTPVTCARTESIAN_1_2_UNDULATION_PRECISION = 3
const SBF_EXTEVENTPVTCARTESIAN_1_2_VX_PRECISION = 3
const SBF_EXTEVENTPVTCARTESIAN_1_2_VY_PRECISION = 3
const SBF_EXTEVENTPVTCARTESIAN_1_2_VZ_PRECISION = 3
const SBF_EXTEVENTPVTCARTESIAN_1_2_COG_PRECISION = 2
const SBF_EXTEVENTPVTCARTESIAN_1_2_RXCLKBIAS_PRECISION = 6
const SBF_EXTEVENTPVTCARTESIAN_1_2_RXCLKDRIFT_PRECISION = 3
const SBF_EXTEVENTPVTCARTESIAN_1_2_HACCURACY_PRECISION = 2
const SBF_EXTEVENTPVTCARTESIAN_1_2_VACCURACY_PRECISION = 2
const SBF_EXTEVENTPVTCARTESIAN_1__PADDING_LENGTH = SBF_EXTEVENTPVTCARTESIAN_1_2__PADDING_LENGTH

/*--ExtEventPVTGeodetic_1_0_t : ---------------------------------------------*/
/* Geodetic position at the instant of an event */
const SBF_EXTEVENTPVTGEODETIC_1_0__PADDING_LENGTH = 3
const SBF_EXTEVENTPVTGEODETIC_1_0_TOW_PRECISION = 3
const SBF_EXTEVENTPVTGEODETIC_1_0_LAT_PRECISION = 12
const SBF_EXTEVENTPVTGEODETIC_1_0_LON_PRECISION = 12
const SBF_EXTEVENTPVTGEODETIC_1_0_ALT_PRECISION = 3
const SBF_EXTEVENTPVTGEODETIC_1_0_UNDULATION_PRECISION = 3
const SBF_EXTEVENTPVTGEODETIC_1_0_VN_PRECISION = 3
const SBF_EXTEVENTPVTGEODETIC_1_0_VE_PRECISION = 3
const SBF_EXTEVENTPVTGEODETIC_1_0_VU_PRECISION = 3
const SBF_EXTEVENTPVTGEODETIC_1_0_COG_PRECISION = 2
const SBF_EXTEVENTPVTGEODETIC_1_0_RXCLKBIAS_PRECISION = 6
const SBF_EXTEVENTPVTGEODETIC_1_0_RXCLKDRIFT_PRECISION = 3

/*--ExtEventPVTGeodetic_1_1_t : ---------------------------------------------*/
/* Geodetic position at the instant of an event */
const SBF_EXTEVENTPVTGEODETIC_1_1_TOW_PRECISION = 3
const SBF_EXTEVENTPVTGEODETIC_1_1_LAT_PRECISION = 12
const SBF_EXTEVENTPVTGEODETIC_1_1_LON_PRECISION = 12
const SBF_EXTEVENTPVTGEODETIC_1_1_ALT_PRECISION = 3
const SBF_EXTEVENTPVTGEODETIC_1_1_UNDULATION_PRECISION = 3
const SBF_EXTEVENTPVTGEODETIC_1_1_VN_PRECISION = 3
const SBF_EXTEVENTPVTGEODETIC_1_1_VE_PRECISION = 3
const SBF_EXTEVENTPVTGEODETIC_1_1_VU_PRECISION = 3
const SBF_EXTEVENTPVTGEODETIC_1_1_COG_PRECISION = 2
const SBF_EXTEVENTPVTGEODETIC_1_1_RXCLKBIAS_PRECISION = 6
const SBF_EXTEVENTPVTGEODETIC_1_1_RXCLKDRIFT_PRECISION = 3

/*--ExtEventPVTGeodetic_1_2_t : ---------------------------------------------*/
/* Geodetic position at the instant of an event */
const SBF_EXTEVENTPVTGEODETIC_1_2__PADDING_LENGTH = 1
const SBF_EXTEVENTPVTGEODETIC_1_2_TOW_PRECISION = 3
const SBF_EXTEVENTPVTGEODETIC_1_2_LAT_PRECISION = 12
const SBF_EXTEVENTPVTGEODETIC_1_2_LON_PRECISION = 12
const SBF_EXTEVENTPVTGEODETIC_1_2_ALT_PRECISION = 3
const SBF_EXTEVENTPVTGEODETIC_1_2_UNDULATION_PRECISION = 3
const SBF_EXTEVENTPVTGEODETIC_1_2_VN_PRECISION = 3
const SBF_EXTEVENTPVTGEODETIC_1_2_VE_PRECISION = 3
const SBF_EXTEVENTPVTGEODETIC_1_2_VU_PRECISION = 3
const SBF_EXTEVENTPVTGEODETIC_1_2_COG_PRECISION = 2
const SBF_EXTEVENTPVTGEODETIC_1_2_RXCLKBIAS_PRECISION = 6
const SBF_EXTEVENTPVTGEODETIC_1_2_RXCLKDRIFT_PRECISION = 3
const SBF_EXTEVENTPVTGEODETIC_1_2_HACCURACY_PRECISION = 2
const SBF_EXTEVENTPVTGEODETIC_1_2_VACCURACY_PRECISION = 2
const SBF_EXTEVENTPVTGEODETIC_1__PADDING_LENGTH = SBF_EXTEVENTPVTGEODETIC_1_2__PADDING_LENGTH

/*--ExtEventBaseVectCart_1_0_t : --------------------------------------------*/
/* XYZ relative position with respect to base(s) at the instant of an event */
const SBF_EXTEVENTVECTORINFOCART_1_0_DELTAX_PRECISION = 3
const SBF_EXTEVENTVECTORINFOCART_1_0_DELTAY_PRECISION = 3
const SBF_EXTEVENTVECTORINFOCART_1_0_DELTAZ_PRECISION = 3
const SBF_EXTEVENTVECTORINFOCART_1_0_DELTAVX_PRECISION = 3
const SBF_EXTEVENTVECTORINFOCART_1_0_DELTAVY_PRECISION = 3
const SBF_EXTEVENTVECTORINFOCART_1_0_DELTAVZ_PRECISION = 3
const SBF_EXTEVENTVECTORINFOCART_1_0_AZIMUTH_PRECISION = 2
const SBF_EXTEVENTVECTORINFOCART_1_0_ELEVATION_PRECISION = 2
const SBF_EXTEVENTBASEVECTCART_1_0_EXTEVENTVECTORINFOCART_LENGTH = 30
const SBF_EXTEVENTBASEVECTCART_1_0_TOW_PRECISION = 3
const SBF_EXTEVENTBASEVECTCART_1_EXTEVENTVECTORINFOCART_LENGTH = SBF_EXTEVENTBASEVECTCART_1_0_EXTEVENTVECTORINFOCART_LENGTH

/*--ExtEventBaseVectGeod_1_0_t : --------------------------------------------*/
/* ENU relative position with respect to base(s) at the instant of an event */
const SBF_EXTEVENTVECTORINFOGEOD_1_0_DELTAEAST_PRECISION = 3
const SBF_EXTEVENTVECTORINFOGEOD_1_0_DELTANORTH_PRECISION = 3
const SBF_EXTEVENTVECTORINFOGEOD_1_0_DELTAUP_PRECISION = 3
const SBF_EXTEVENTVECTORINFOGEOD_1_0_DELTAVE_PRECISION = 3
const SBF_EXTEVENTVECTORINFOGEOD_1_0_DELTAVN_PRECISION = 3
const SBF_EXTEVENTVECTORINFOGEOD_1_0_DELTAVU_PRECISION = 3
const SBF_EXTEVENTVECTORINFOGEOD_1_0_AZIMUTH_PRECISION = 2
const SBF_EXTEVENTVECTORINFOGEOD_1_0_ELEVATION_PRECISION = 2
const SBF_EXTEVENTBASEVECTGEOD_1_0_EXTEVENTVECTORINFOGEOD_LENGTH = 30
const SBF_EXTEVENTBASEVECTGEOD_1_0_TOW_PRECISION = 3
const SBF_EXTEVENTBASEVECTGEOD_1_EXTEVENTVECTORINFOGEOD_LENGTH = SBF_EXTEVENTBASEVECTGEOD_1_0_EXTEVENTVECTORINFOGEOD_LENGTH

/*--ExtEventINSNavCart_1_0_t : ----------------------------------------------*/
/* INS solution in Cartesian coordinates at the instant of an event */
const SBF_EXTEVENTINSNAVCARTPOSSTDDEV_1_0_XSTDDEV_PRECISION = 3
const SBF_EXTEVENTINSNAVCARTPOSSTDDEV_1_0_YSTDDEV_PRECISION = 3
const SBF_EXTEVENTINSNAVCARTPOSSTDDEV_1_0_ZSTDDEV_PRECISION = 3
const SBF_EXTEVENTINSNAVCARTATT_1_0_HEADING_PRECISION = 6
const SBF_EXTEVENTINSNAVCARTATT_1_0_PITCH_PRECISION = 6
const SBF_EXTEVENTINSNAVCARTATT_1_0_ROLL_PRECISION = 6
const SBF_EXTEVENTINSNAVCARTATTSTDDEV_1_0_HEADINGSTDDEV_PRECISION = 6
const SBF_EXTEVENTINSNAVCARTATTSTDDEV_1_0_PITCHSTDDEV_PRECISION = 6
const SBF_EXTEVENTINSNAVCARTATTSTDDEV_1_0_ROLLSTDDEV_PRECISION = 6
const SBF_EXTEVENTINSNAVCARTVEL_1_0_VX_PRECISION = 3
const SBF_EXTEVENTINSNAVCARTVEL_1_0_VY_PRECISION = 3
const SBF_EXTEVENTINSNAVCARTVEL_1_0_VZ_PRECISION = 3
const SBF_EXTEVENTINSNAVCARTVELSTDDEV_1_0_VXSTDDEV_PRECISION = 3
const SBF_EXTEVENTINSNAVCARTVELSTDDEV_1_0_VYSTDDEV_PRECISION = 3
const SBF_EXTEVENTINSNAVCARTVELSTDDEV_1_0_VZSTDDEV_PRECISION = 3
const SBF_EXTEVENTINSNAVCART_1_0_EXTEVENTINSNAVCARTDATA_LENGTH = 16
const SBF_EXTEVENTINSNAVCART_1_0_TOW_PRECISION = 3
const SBF_EXTEVENTINSNAVCART_1_0_X_PRECISION = 3
const SBF_EXTEVENTINSNAVCART_1_0_Y_PRECISION = 3
const SBF_EXTEVENTINSNAVCART_1_0_Z_PRECISION = 3
const SBF_EXTEVENTINSNAVCART_1_0_ACCURACY_PRECISION = 2
const SBF_EXTEVENTINSNAVCART_1_EXTEVENTINSNAVCARTDATA_LENGTH = SBF_EXTEVENTINSNAVCART_1_0_EXTEVENTINSNAVCARTDATA_LENGTH

/*--ExtEventINSNavGeod_1_0_t : ----------------------------------------------*/
/* INS solution in Geodetic coordinates at the instant of an event */
const SBF_EXTEVENTINSNAVGEODPOSSTDDEV_1_0_LATITUDESTDDEV_PRECISION = 3
const SBF_EXTEVENTINSNAVGEODPOSSTDDEV_1_0_LONGITUDESTDDEV_PRECISION = 3
const SBF_EXTEVENTINSNAVGEODPOSSTDDEV_1_0_HEIGHTSTDDEV_PRECISION = 3
const SBF_EXTEVENTINSNAVGEODATT_1_0_HEADING_PRECISION = 6
const SBF_EXTEVENTINSNAVGEODATT_1_0_PITCH_PRECISION = 6
const SBF_EXTEVENTINSNAVGEODATT_1_0_ROLL_PRECISION = 6
const SBF_EXTEVENTINSNAVGEODATTSTDDEV_1_0_HEADINGSTDDEV_PRECISION = 6
const SBF_EXTEVENTINSNAVGEODATTSTDDEV_1_0_PITCHSTDDEV_PRECISION = 6
const SBF_EXTEVENTINSNAVGEODATTSTDDEV_1_0_ROLLSTDDEV_PRECISION = 6
const SBF_EXTEVENTINSNAVGEODVEL_1_0_VE_PRECISION = 3
const SBF_EXTEVENTINSNAVGEODVEL_1_0_VN_PRECISION = 3
const SBF_EXTEVENTINSNAVGEODVEL_1_0_VU_PRECISION = 3
const SBF_EXTEVENTINSNAVGEODVELSTDDEV_1_0_VESTDDEV_PRECISION = 3
const SBF_EXTEVENTINSNAVGEODVELSTDDEV_1_0_VNSTDDEV_PRECISION = 3
const SBF_EXTEVENTINSNAVGEODVELSTDDEV_1_0_VUSTDDEV_PRECISION = 3
const SBF_EXTEVENTINSNAVGEOD_1_0_EXTEVENTINSNAVGEODDATA_LENGTH = 16
const SBF_EXTEVENTINSNAVGEOD_1_0_TOW_PRECISION = 3
const SBF_EXTEVENTINSNAVGEOD_1_0_LATITUDE_PRECISION = 12
const SBF_EXTEVENTINSNAVGEOD_1_0_LONGITUDE_PRECISION = 12
const SBF_EXTEVENTINSNAVGEOD_1_0_HEIGHT_PRECISION = 3
const SBF_EXTEVENTINSNAVGEOD_1_0_UNDULATION_PRECISION = 3
const SBF_EXTEVENTINSNAVGEOD_1_0_ACCURACY_PRECISION = 2
const SBF_EXTEVENTINSNAVGEOD_1_EXTEVENTINSNAVGEODDATA_LENGTH = SBF_EXTEVENTINSNAVGEOD_1_0_EXTEVENTINSNAVGEODDATA_LENGTH

/*--ExtEventAttEuler_1_0_t : ------------------------------------------------*/
/* GNSS attitude expressed as Euler angles at the instant of an event */
const SBF_EXTEVENTATTEULER_1_0_TOW_PRECISION = 3
const SBF_EXTEVENTATTEULER_1_0_HEADING_PRECISION = 6
const SBF_EXTEVENTATTEULER_1_0_PITCH_PRECISION = 6
const SBF_EXTEVENTATTEULER_1_0_ROLL_PRECISION = 6
const SBF_EXTEVENTATTEULER_1_0_PITCHDOT_PRECISION = 3
const SBF_EXTEVENTATTEULER_1_0_ROLLDOT_PRECISION = 3
const SBF_EXTEVENTATTEULER_1_0_HEADINGDOT_PRECISION = 3

/*==Differential Correction Blocks===========================================*/

/*--DiffCorrIn_1_0_t : ------------------------------------------------------*/
/* Incoming RTCM or CMR message */
const SBF_DIFFCORRIN_1_0_FRAME_LENGTH = 4096
const SBF_DIFFCORRIN_1_0_TOW_PRECISION = 3
const SBF_DIFFCORRIN_1_FRAME_LENGTH = SBF_DIFFCORRIN_1_0_FRAME_LENGTH

/*--BaseStation_1_0_t : -----------------------------------------------------*/
/* Base station coordinates */
const SBF_BASESTATION_1_0_TOW_PRECISION = 3
const SBF_BASESTATION_1_0_X_L1PHASECENTER_PRECISION = 3
const SBF_BASESTATION_1_0_Y_L1PHASECENTER_PRECISION = 3
const SBF_BASESTATION_1_0_Z_L1PHASECENTER_PRECISION = 3

/*--RTCMDatum_1_0_t : -------------------------------------------------------*/
/* Datum information from the RTK service provider */
const SBF_RTCMDATUM_1_0_SOURCECRS_LENGTH = 32
const SBF_RTCMDATUM_1_0_TARGETCRS_LENGTH = 32
const SBF_RTCMDATUM_1_0_TOW_PRECISION = 3
const SBF_RTCMDATUM_1_SOURCECRS_LENGTH = SBF_RTCMDATUM_1_0_SOURCECRS_LENGTH
const SBF_RTCMDATUM_1_TARGETCRS_LENGTH = SBF_RTCMDATUM_1_0_TARGETCRS_LENGTH

/*--BaseLink_1_0_t : --------------------------------------------------------*/
const SBF_BASELINK_1_0_TOW_PRECISION = 3
const SBF_BASELINK_1_0_AGEOFLASTMSG_PRECISION = 3

/*==L-Band Demodulator Blocks================================================*/

/*--LBandReceiverStatus_1_0_t : ---------------------------------------------*/
const SBF_LBANDRECEIVERSTATUS_1_0_TOW_PRECISION = 3

/*--LBandTrackerStatus_1_0_t : ----------------------------------------------*/
/* Status of the L-band signal tracking */
const SBF_TRACKDATA_1_0_FREQOFFSET_PRECISION = 3
const SBF_TRACKDATA_1_0_CN0_PRECISION = 2
const SBF_LBANDTRACKERSTATUS_1_0_TRACKDATA_LENGTH = 5
const SBF_LBANDTRACKERSTATUS_1_0_TOW_PRECISION = 3

/*--LBandTrackerStatus_1_1_t : ----------------------------------------------*/
/* Status of the L-band signal tracking */
const SBF_TRACKDATA_1_1__PADDING_LENGTH = 2
const SBF_TRACKDATA_1_1_FREQOFFSET_PRECISION = 3
const SBF_TRACKDATA_1_1_CN0_PRECISION = 2
const SBF_LBANDTRACKERSTATUS_1_1_TRACKDATA_LENGTH = 5
const SBF_LBANDTRACKERSTATUS_1_1_TOW_PRECISION = 3

/*--LBandTrackerStatus_1_2_t : ----------------------------------------------*/
/* Status of the L-band signal tracking */
const SBF_TRACKDATA_1_2__PADDING_LENGTH = 2
const SBF_TRACKDATA_1_2_FREQOFFSET_PRECISION = 3
const SBF_TRACKDATA_1_2_CN0_PRECISION = 2
const SBF_LBANDTRACKERSTATUS_1_2_TRACKDATA_LENGTH = 5
const SBF_LBANDTRACKERSTATUS_1_2_TOW_PRECISION = 3

/*--LBandTrackerStatus_1_3_t : ----------------------------------------------*/
/* Status of the L-band signal tracking */
const SBF_TRACKDATA_1_3__PADDING_LENGTH = 1
const SBF_TRACKDATA_1_3_FREQOFFSET_PRECISION = 3
const SBF_TRACKDATA_1_3_CN0_PRECISION = 2
const SBF_LBANDTRACKERSTATUS_1_3_TRACKDATA_LENGTH = 5
const SBF_LBANDTRACKERSTATUS_1_3_TOW_PRECISION = 3
const SBF_TRACKDATA_1__PADDING_LENGTH = SBF_TRACKDATA_1_3__PADDING_LENGTH
const SBF_LBANDTRACKERSTATUS_1_TRACKDATA_LENGTH = SBF_LBANDTRACKERSTATUS_1_3_TRACKDATA_LENGTH

/*--LBAS1DecoderStatus_1_0_t : ----------------------------------------------*/
/* Status of the LBAS1 L-band service */
const SBF_LBAS1DECODERSTATUS_1_0_TOW_PRECISION = 3

/*--LBAS1DecoderStatus_1_1_t : ----------------------------------------------*/
/* Status of the LBAS1 L-band service */
const SBF_LBAS1DECODERSTATUS_1_1_PAC_LENGTH = 20
const SBF_LBAS1DECODERSTATUS_1_1_TOW_PRECISION = 3

/*--LBAS1DecoderStatus_1_2_t : ----------------------------------------------*/
/* Status of the LBAS1 L-band service */
const SBF_LBAS1DECODERSTATUS_1_2_PAC_LENGTH = 20
const SBF_LBAS1DECODERSTATUS_1_2__PADDING_LENGTH = 2
const SBF_LBAS1DECODERSTATUS_1_2_TOW_PRECISION = 3
const SBF_LBAS1DECODERSTATUS_1_PAC_LENGTH = SBF_LBAS1DECODERSTATUS_1_2_PAC_LENGTH
const SBF_LBAS1DECODERSTATUS_1__PADDING_LENGTH = SBF_LBAS1DECODERSTATUS_1_2__PADDING_LENGTH

/*--LBAS1Messages_1_0_t : ---------------------------------------------------*/
/* LBAS1over-the-air message */
const SBF_LBAS1MESSAGES_1_0_MESSAGE_LENGTH = 512
const SBF_LBAS1MESSAGES_1_0_TOW_PRECISION = 3

/*--LBandBeams_1_0_t : ------------------------------------------------------*/
/* L-band satellite/beam information */
const SBF_BEAMINFO_1_0_SATNAME_LENGTH = 9
const SBF_LBANDBEAMS_1_0_BEAMDATA_LENGTH = 13
const SBF_LBANDBEAMS_1_0_TOW_PRECISION = 3
const SBF_BEAMINFO_1_SATNAME_LENGTH = SBF_BEAMINFO_1_0_SATNAME_LENGTH
const SBF_LBANDBEAMS_1_BEAMDATA_LENGTH = SBF_LBANDBEAMS_1_0_BEAMDATA_LENGTH

/*--LBandRaw_1_0_t : --------------------------------------------------------*/
/* L-Band raw user data */
const SBF_LBANDRAW_1_0_USERDATA_LENGTH = 2048
const SBF_LBANDRAW_1_0_TOW_PRECISION = 3
const SBF_LBANDRAW_1_USERDATA_LENGTH = SBF_LBANDRAW_1_0_USERDATA_LENGTH

/*--FugroStatus_1_0_t : -----------------------------------------------------*/
/* Fugro Status Information */
const SBF_FUGROSTATUS_1_0_TOW_PRECISION = 3

/*==External Sensor Blocks===================================================*/

/*--ExtSensorMeas_1_0_t : ---------------------------------------------------*/
/* Measurement set of external sensors of one epoch */
const SBF_EXTSENSORMEASACCELERATION_1_0_ACCELERATIONX_PRECISION = 3
const SBF_EXTSENSORMEASACCELERATION_1_0_ACCELERATIONY_PRECISION = 3
const SBF_EXTSENSORMEASACCELERATION_1_0_ACCELERATIONZ_PRECISION = 3
const SBF_EXTSENSORMEASANGULARRATE_1_0_ANGULARRATEX_PRECISION = 3
const SBF_EXTSENSORMEASANGULARRATE_1_0_ANGULARRATEY_PRECISION = 3
const SBF_EXTSENSORMEASANGULARRATE_1_0_ANGULARRATEZ_PRECISION = 3
const SBF_EXTSENSORMEASVELOCITY_1_0_VELOCITYX_PRECISION = 3
const SBF_EXTSENSORMEASVELOCITY_1_0_VELOCITYY_PRECISION = 3
const SBF_EXTSENSORMEASVELOCITY_1_0_VELOCITYZ_PRECISION = 3
const SBF_EXTSENSORMEASVELOCITY_1_0_STDDEVX_PRECISION = 3
const SBF_EXTSENSORMEASVELOCITY_1_0_STDDEVY_PRECISION = 3
const SBF_EXTSENSORMEASVELOCITY_1_0_STDDEVZ_PRECISION = 3
const SBF_EXTSENSORMEAS_1_0_EXTSENSORMEAS_LENGTH = 4
const SBF_EXTSENSORMEAS_1_0_TOW_PRECISION = 3
const SBF_EXTSENSORMEAS_1_EXTSENSORMEAS_LENGTH = SBF_EXTSENSORMEAS_1_0_EXTSENSORMEAS_LENGTH

/*--ExtSensorStatus_1_0_t : -------------------------------------------------*/
/* Overall status of external sensors */
const SBF_EXTSENSORSTATUS_1_0_STATUSBITS_LENGTH = 60
const SBF_EXTSENSORSTATUS_1_0_TOW_PRECISION = 3
const SBF_EXTSENSORSTATUS_1_STATUSBITS_LENGTH = SBF_EXTSENSORSTATUS_1_0_STATUSBITS_LENGTH

/*--ExtSensorSetup_1_0_t : --------------------------------------------------*/
/* General information about the setup of external sensors */
const SBF_EXTSENSORSETUP_1_0_EXTSENSORSETUP_LENGTH = 4
const SBF_EXTSENSORSETUP_1_0_TOW_PRECISION = 3

/*--ExtSensorSetup_1_1_t : --------------------------------------------------*/
/* General information about the setup of external sensors */
const SBF_ONESENSOR_1_1__PADDING_LENGTH = 2
const SBF_EXTSENSORSETUP_1_1_EXTSENSORSETUP_LENGTH = 4
const SBF_EXTSENSORSETUP_1_1_TOW_PRECISION = 3

/*--ExtSensorSetup_1_2_t : --------------------------------------------------*/
/* General information about the setup of external sensors */
const SBF_EXTSENSORSETUP_1_2_EXTSENSORSETUP_LENGTH = 4
const SBF_EXTSENSORSETUP_1_2_TOW_PRECISION = 3
const SBF_EXTSENSORSETUP_1_EXTSENSORSETUP_LENGTH = SBF_EXTSENSORSETUP_1_2_EXTSENSORSETUP_LENGTH

/*--ExtSensorStatus_2_0_t : -------------------------------------------------*/
/* Overall status of external sensors */
const SBF_EXTSENSORSTATUS_2_0_DATA_LENGTH = 4
const SBF_EXTSENSORSTATUS_2_0_TOW_PRECISION = 3
const SBF_EXTSENSORSTATUS_2_DATA_LENGTH = SBF_EXTSENSORSTATUS_2_0_DATA_LENGTH

/*--ExtSensorInfo_1_0_t : ---------------------------------------------------*/
/* Configuration information of external sensors */
const SBF_EXTSENSORINFO_1_0_DATA_LENGTH = 60
const SBF_EXTSENSORINFO_1_0_TOW_PRECISION = 3
const SBF_EXTSENSORINFO_1_DATA_LENGTH = SBF_EXTSENSORINFO_1_0_DATA_LENGTH

/*--IMUSetup_1_0_t : --------------------------------------------------------*/
/* General information about the setup of the IMU */
const SBF_IMUSETUP_1_0_TOW_PRECISION = 3
const SBF_IMUSETUP_1_0_ANTLEVERARMX_PRECISION = 3
const SBF_IMUSETUP_1_0_ANTLEVERARMY_PRECISION = 3
const SBF_IMUSETUP_1_0_ANTLEVERARMZ_PRECISION = 3
const SBF_IMUSETUP_1_0_THETAX_PRECISION = 2
const SBF_IMUSETUP_1_0_THETAY_PRECISION = 2
const SBF_IMUSETUP_1_0_THETAZ_PRECISION = 2

/*==Status Blocks============================================================*/

/*--ReceiverStatus_1_0_t : --------------------------------------------------*/
/* Overall status information of the receiver */
const SBF_RECEIVERSTATUS_1_0_TOW_PRECISION = 3

/*--TrackingStatus_1_0_t : --------------------------------------------------*/
/* Status of the tracking for all receiver channels */
const SBF_TRACKINGSTATUSCHANNEL_1_0__PADDING_LENGTH = 3
const SBF_TRACKINGSTATUS_1_0_CHANNELDATA_LENGTH = 72
const SBF_TRACKINGSTATUS_1_0_TOW_PRECISION = 3
const SBF_TRACKINGSTATUSCHANNEL_1__PADDING_LENGTH = SBF_TRACKINGSTATUSCHANNEL_1_0__PADDING_LENGTH
const SBF_TRACKINGSTATUS_1_CHANNELDATA_LENGTH = SBF_TRACKINGSTATUS_1_0_CHANNELDATA_LENGTH

/*--ChannelStatus_1_0_t : ---------------------------------------------------*/
/* Status of the tracking for all receiver channels */
const SBF_CHANNELSTATUS_1_0_TOW_PRECISION = 3
const SBF_CHANNELSTATUS_1_DATA_LENGTH = SBF_CHANNELSTATUS_1_0_DATA_LENGTH

/*--ReceiverStatus_2_0_t : --------------------------------------------------*/
/* Overall status information of the receiver */
const SBF_RECEIVERSTATUS_2_0_AGCSTATE_LENGTH = 24
const SBF_RECEIVERSTATUS_2_0_TOW_PRECISION = 3

/*--ReceiverStatus_2_1_t : --------------------------------------------------*/
/* Overall status information of the receiver */
const SBF_RECEIVERSTATUS_2_1_AGCSTATE_LENGTH = 24
const SBF_RECEIVERSTATUS_2_1_TOW_PRECISION = 3
const SBF_RECEIVERSTATUS_2_AGCSTATE_LENGTH = SBF_RECEIVERSTATUS_2_1_AGCSTATE_LENGTH

/*--SatVisibility_1_0_t : ---------------------------------------------------*/
/* Azimuth/elevation of visible satellites */
const SBF_SATINFO_1_0_AZIMUTH_PRECISION = 2
const SBF_SATINFO_1_0_ELEVATION_PRECISION = 2
const SBF_SATVISIBILITY_1_0_SATINFO_LENGTH = 100
const SBF_SATVISIBILITY_1_0_TOW_PRECISION = 3
const SBF_SATVISIBILITY_1_SATINFO_LENGTH = SBF_SATVISIBILITY_1_0_SATINFO_LENGTH

/*--InputLink_1_0_t : -------------------------------------------------------*/
/* Statistics on input streams */
const SBF_INPUTLINK_1_0_TOW_PRECISION = 3
const SBF_INPUTLINK_1_DATA_LENGTH = SBF_INPUTLINK_1_0_DATA_LENGTH

/*--OutputLink_1_0_t : ------------------------------------------------------*/
/* Statistics on output streams */
const SBF_OUTPUTTYPESUB_1_0__PADDING_LENGTH = 2
const SBF_OUTPUTLINK_1_0_TOW_PRECISION = 3

/*--OutputLink_1_1_t : ------------------------------------------------------*/
/* Statistics on output streams */
const SBF_OUTPUTLINK_1_1_TOW_PRECISION = 3
const SBF_OUTPUTLINK_1_DATA_LENGTH = SBF_OUTPUTLINK_1_1_DATA_LENGTH

/*--NTRIPClientStatus_1_0_t : -----------------------------------------------*/
/* NTRIP client connection status */
const SBF_NTRIPCLIENTSTATUS_1_0_NTRIPCLIENTCONNECTION_LENGTH = 5
const SBF_NTRIPCLIENTSTATUS_1_0_TOW_PRECISION = 3
const SBF_NTRIPCLIENTSTATUS_1_NTRIPCLIENTCONNECTION_LENGTH = SBF_NTRIPCLIENTSTATUS_1_0_NTRIPCLIENTCONNECTION_LENGTH

/*--NTRIPServerStatus_1_0_t : -----------------------------------------------*/
/* NTRIP server connection status */
const SBF_NTRIPSERVERSTATUS_1_0_NTRIPSERVERCONNECTION_LENGTH = 5
const SBF_NTRIPSERVERSTATUS_1_0_TOW_PRECISION = 3
const SBF_NTRIPSERVERSTATUS_1_NTRIPSERVERCONNECTION_LENGTH = SBF_NTRIPSERVERSTATUS_1_0_NTRIPSERVERCONNECTION_LENGTH

/*--IPStatus_1_0_t : --------------------------------------------------------*/
/* IP address, gateway and MAC address of Ethernet interface */
const SBF_IPSTATUS_1_0_MACADDRESS_LENGTH = 6
const SBF_IPSTATUS_1_0_IPADDRESS_LENGTH = 16
const SBF_IPSTATUS_1_0_GATEWAY_LENGTH = 16
const SBF_IPSTATUS_1_0_TOW_PRECISION = 3

/*--IPStatus_1_1_t : --------------------------------------------------------*/
/* IP address, gateway and MAC address of Ethernet interface */
const SBF_IPSTATUS_1_1_MACADDRESS_LENGTH = 6
const SBF_IPSTATUS_1_1_IPADDRESS_LENGTH = 16
const SBF_IPSTATUS_1_1_GATEWAY_LENGTH = 16
const SBF_IPSTATUS_1_1_HOSTNAME_LENGTH = 32
const SBF_IPSTATUS_1_1_TOW_PRECISION = 3
const SBF_IPSTATUS_1_MACADDRESS_LENGTH = SBF_IPSTATUS_1_1_MACADDRESS_LENGTH
const SBF_IPSTATUS_1_IPADDRESS_LENGTH = SBF_IPSTATUS_1_1_IPADDRESS_LENGTH
const SBF_IPSTATUS_1_GATEWAY_LENGTH = SBF_IPSTATUS_1_1_GATEWAY_LENGTH
const SBF_IPSTATUS_1_HOSTNAME_LENGTH = SBF_IPSTATUS_1_1_HOSTNAME_LENGTH

/*--WiFiAPStatus_1_0_t : ----------------------------------------------------*/
/* WiFi status in access point mode */
const SBF_WIFICLIENT_1_0_CLIENTHOSTNAME_LENGTH = 32
const SBF_WIFICLIENT_1_0_CLIENTMACADDRESS_LENGTH = 6
const SBF_WIFICLIENT_1_0_CLIENTIPADDRESS_LENGTH = 16
const SBF_WIFICLIENT_1_0__PADDING_LENGTH = 2
const SBF_WIFIAPSTATUS_1_0_APIPADDRESS_LENGTH = 16
const SBF_WIFIAPSTATUS_1_0_WIFICLIENT_LENGTH = 20
const SBF_WIFIAPSTATUS_1_0_TOW_PRECISION = 3
const SBF_WIFICLIENT_1_CLIENTHOSTNAME_LENGTH = SBF_WIFICLIENT_1_0_CLIENTHOSTNAME_LENGTH
const SBF_WIFICLIENT_1_CLIENTMACADDRESS_LENGTH = SBF_WIFICLIENT_1_0_CLIENTMACADDRESS_LENGTH
const SBF_WIFICLIENT_1_CLIENTIPADDRESS_LENGTH = SBF_WIFICLIENT_1_0_CLIENTIPADDRESS_LENGTH
const SBF_WIFICLIENT_1__PADDING_LENGTH = SBF_WIFICLIENT_1_0__PADDING_LENGTH
const SBF_WIFIAPSTATUS_1_APIPADDRESS_LENGTH = SBF_WIFIAPSTATUS_1_0_APIPADDRESS_LENGTH
const SBF_WIFIAPSTATUS_1_WIFICLIENT_LENGTH = SBF_WIFIAPSTATUS_1_0_WIFICLIENT_LENGTH

/*--WiFiClientStatus_1_0_t : ------------------------------------------------*/
/* WiFi status in client mode */
const SBF_WIFICLIENTSTATUS_1_0_SSID_AP_LENGTH = 32
const SBF_WIFICLIENTSTATUS_1_0_IPADDRESS_LENGTH = 16
const SBF_WIFICLIENTSTATUS_1_0__PADDING_LENGTH = 2
const SBF_WIFICLIENTSTATUS_1_0_TOW_PRECISION = 3
const SBF_WIFICLIENTSTATUS_1_SSID_AP_LENGTH = SBF_WIFICLIENTSTATUS_1_0_SSID_AP_LENGTH
const SBF_WIFICLIENTSTATUS_1_IPADDRESS_LENGTH = SBF_WIFICLIENTSTATUS_1_0_IPADDRESS_LENGTH
const SBF_WIFICLIENTSTATUS_1__PADDING_LENGTH = SBF_WIFICLIENTSTATUS_1_0__PADDING_LENGTH

/*--CellularStatus_1_0_t : --------------------------------------------------*/
/* Cellular status */
const SBF_CELLULARSTATUS_1_0_OPERATORNAME_LENGTH = 20
const SBF_CELLULARSTATUS_1_0__PADDING_LENGTH = 2
const SBF_CELLULARSTATUS_1_0_TOW_PRECISION = 3

/*--CellularStatus_1_1_t : --------------------------------------------------*/
/* Cellular status */
const SBF_CELLULARSTATUS_1_1_OPERATORNAME_LENGTH = 20
const SBF_CELLULARSTATUS_1_1_TOW_PRECISION = 3
const SBF_CELLULARSTATUS_1_OPERATORNAME_LENGTH = SBF_CELLULARSTATUS_1_1_OPERATORNAME_LENGTH

/*--BluetoothStatus_1_0_t : -------------------------------------------------*/
/* Bluetooth status */
const SBF_BTDEVICE_1_0_DEVICENAME_LENGTH = 30
const SBF_BTDEVICE_1_0__PADDING_LENGTH = 1
const SBF_BLUETOOTHSTATUS_1_0_BTDEVICE_LENGTH = 20
const SBF_BLUETOOTHSTATUS_1_0_TOW_PRECISION = 3
const SBF_BTDEVICE_1_DEVICENAME_LENGTH = SBF_BTDEVICE_1_0_DEVICENAME_LENGTH
const SBF_BTDEVICE_1__PADDING_LENGTH = SBF_BTDEVICE_1_0__PADDING_LENGTH
const SBF_BLUETOOTHSTATUS_1_BTDEVICE_LENGTH = SBF_BLUETOOTHSTATUS_1_0_BTDEVICE_LENGTH

/*--DynDNSStatus_1_0_t : ----------------------------------------------------*/
/* DynDNS status */
const SBF_DYNDNSSTATUS_1_0_TOW_PRECISION = 3

/*--DynDNSStatus_1_1_t : ----------------------------------------------------*/
/* DynDNS status */
const SBF_DYNDNSSTATUS_1_1_IPADDRESS_LENGTH = 16
const SBF_DYNDNSSTATUS_1_1_TOW_PRECISION = 3
const SBF_DYNDNSSTATUS_1_IPADDRESS_LENGTH = SBF_DYNDNSSTATUS_1_1_IPADDRESS_LENGTH

/*--BatteryStatus_1_0_t : ---------------------------------------------------*/
/* Battery status */
const SBF_BATTERYSTATUS_1_0_BATTERY_LENGTH = 4
const SBF_BATTERYSTATUS_1_0_TOW_PRECISION = 3

/*--BatteryStatus_1_1_t : ---------------------------------------------------*/
/* Battery status */
const SBF_BATTERYSTATUS_1_1_BATTERY_LENGTH = 4
const SBF_BATTERYSTATUS_1_1_TOW_PRECISION = 3

/*--BatteryStatus_1_2_t : ---------------------------------------------------*/
/* Battery status */
const SBF_BATTERYSTATUS_1_2_BATTERY_LENGTH = 4
const SBF_BATTERYSTATUS_1_2_TOW_PRECISION = 3
const SBF_BATTERYSTATUS_1_BATTERY_LENGTH = SBF_BATTERYSTATUS_1_2_BATTERY_LENGTH

/*--PowerStatus_1_0_t : -----------------------------------------------------*/
/* Power supply source and voltage */
const SBF_POWERSTATUS_1_0_TOW_PRECISION = 3

/*--QualityInd_1_0_t : ------------------------------------------------------*/
/* Quality indicators */
const SBF_QUALITYIND_1_0_INDICATORS_LENGTH = 40
const SBF_QUALITYIND_1_0_TOW_PRECISION = 3
const SBF_QUALITYIND_1_INDICATORS_LENGTH = SBF_QUALITYIND_1_0_INDICATORS_LENGTH

/*--DiskStatus_1_0_t : ------------------------------------------------------*/
/* Internal logging status */
const SBF_DISKDATA_1_0__PADDING_LENGTH = 3
const SBF_DISKSTATUS_1_0_DISKDATA_LENGTH = 2
const SBF_DISKSTATUS_1_0_TOW_PRECISION = 3

/*--DiskStatus_1_1_t : ------------------------------------------------------*/
/* Internal logging status */
const SBF_DISKDATA_1_1__PADDING_LENGTH = 2
const SBF_DISKSTATUS_1_1_DISKDATA_LENGTH = 2
const SBF_DISKSTATUS_1_1_TOW_PRECISION = 3
const SBF_DISKDATA_1__PADDING_LENGTH = SBF_DISKDATA_1_1__PADDING_LENGTH
const SBF_DISKSTATUS_1_DISKDATA_LENGTH = SBF_DISKSTATUS_1_1_DISKDATA_LENGTH

/*--LogStatus_1_0_t : -------------------------------------------------------*/
/* Log sessions status */
const SBF_LOGSESSION_1_0_FILEUPLOADSTATUS_LENGTH = 11
const SBF_LOGSTATUS_1_0_LOGSESSIONS_LENGTH = 8
const SBF_LOGSTATUS_1_0_TOW_PRECISION = 3
const SBF_LOGSESSION_1_FILEUPLOADSTATUS_LENGTH = SBF_LOGSESSION_1_0_FILEUPLOADSTATUS_LENGTH
const SBF_LOGSTATUS_1_LOGSESSIONS_LENGTH = SBF_LOGSTATUS_1_0_LOGSESSIONS_LENGTH

/*--UHFStatus_1_0_t : -------------------------------------------------------*/
/* UHF status */
const SBF_UHFDATA_1_0__PADDING_LENGTH = 1
const SBF_UHFSTATUS_1_0_UHFDATA_LENGTH = 20
const SBF_UHFSTATUS_1_0_TOW_PRECISION = 3
const SBF_UHFDATA_1__PADDING_LENGTH = SBF_UHFDATA_1_0__PADDING_LENGTH
const SBF_UHFSTATUS_1_UHFDATA_LENGTH = SBF_UHFSTATUS_1_0_UHFDATA_LENGTH

/*--RFStatus_1_0_t : --------------------------------------------------------*/
/* Radio-frequency interference mitigation status */
const SBF_RFBAND_1_0__PADDING_LENGTH = 1
const SBF_RFSTATUS_1_0_RFBAND_LENGTH = 10
const SBF_RFSTATUS_1_0_TOW_PRECISION = 3
const SBF_RFBAND_1__PADDING_LENGTH = SBF_RFBAND_1_0__PADDING_LENGTH
const SBF_RFSTATUS_1_RFBAND_LENGTH = SBF_RFSTATUS_1_0_RFBAND_LENGTH

/*--RIMSHealth_1_0_t : ------------------------------------------------------*/
/* Health status of the receiver */
const SBF_RIMSHEALTH_1_0__PADDING_LENGTH = 1
const SBF_RIMSHEALTH_1_0_TOW_PRECISION = 3
const SBF_RIMSHEALTH_1__PADDING_LENGTH = SBF_RIMSHEALTH_1_0__PADDING_LENGTH

/*--OSNMAStatus_1_0_t : -----------------------------------------------------*/
/* OSNMA status information */
const SBF_OSNMASTATUS_1_0_RESERVED1_LENGTH = 7
const SBF_OSNMASTATUS_1_0_MACKSTATUS_LENGTH = 52
const SBF_OSNMASTATUS_1_0_TOW_PRECISION = 3
const SBF_OSNMASTATUS_1_RESERVED1_LENGTH = SBF_OSNMASTATUS_1_0_RESERVED1_LENGTH
const SBF_OSNMASTATUS_1_MACKSTATUS_LENGTH = SBF_OSNMASTATUS_1_0_MACKSTATUS_LENGTH

/*--GALNavMonitor_1_0_t : ---------------------------------------------------*/
/* Monitoring navigation data per Galileo satellite. */
const SBF_GALNAVMONITOR_1_0__PADDING_LENGTH = 3
const SBF_GALNAVMONITOR_1_0_TOW_PRECISION = 3
const SBF_GALNAVMONITOR_1_0_POSITIONDIFF_PRECISION = 3
const SBF_GALNAVMONITOR_1_0_TIMECORRDIFF_PRECISION = 3
const SBF_GALNAVMONITOR_1__PADDING_LENGTH = SBF_GALNAVMONITOR_1_0__PADDING_LENGTH

/*--INAVmonitor_1_0_t : -----------------------------------------------------*/
/* Reed-Solomon and SSP status information */
const SBF_INAVMONITOR_1_0__PADDING_LENGTH = 3
const SBF_INAVMONITOR_1_0_TOW_PRECISION = 3
const SBF_INAVMONITOR_1__PADDING_LENGTH = SBF_INAVMONITOR_1_0__PADDING_LENGTH

/*--P2PPStatus_1_0_t : ------------------------------------------------------*/
/* P2PP client/server status */
const SBF_P2PPSTATUS_1_0_P2PPSESSION_LENGTH = 4
const SBF_P2PPSTATUS_1_0_TOW_PRECISION = 3
const SBF_P2PPSTATUS_1_P2PPSESSION_LENGTH = SBF_P2PPSTATUS_1_0_P2PPSESSION_LENGTH

/*--AuthenticationStatus_1_0_t : --------------------------------------------*/
const SBF_AUTHENTICATIONSTATUS_1_0_TOW_PRECISION = 3

/*--CosmosStatus_1_0_t : ----------------------------------------------------*/
/* Cosmos receiver service status */
const SBF_COSMOSSTATUS_1_0__PADDING_LENGTH = 1
const SBF_COSMOSSTATUS_1_0_TOW_PRECISION = 3
const SBF_COSMOSSTATUS_1__PADDING_LENGTH = SBF_COSMOSSTATUS_1_0__PADDING_LENGTH

/*==Miscellaneous Blocks=====================================================*/

/*--ReceiverSetup_1_0_t : ---------------------------------------------------*/
/* General information about the receiver installation */
const SBF_RECEIVERSETUP_1_0_MARKERNAME_LENGTH = 60
const SBF_RECEIVERSETUP_1_0_MARKERNUMBER_LENGTH = 20
const SBF_RECEIVERSETUP_1_0_OBSERVER_LENGTH = 20
const SBF_RECEIVERSETUP_1_0_AGENCY_LENGTH = 40
const SBF_RECEIVERSETUP_1_0_RXSERIALNBR_LENGTH = 20
const SBF_RECEIVERSETUP_1_0_RXNAME_LENGTH = 20
const SBF_RECEIVERSETUP_1_0_RXVERSION_LENGTH = 20
const SBF_RECEIVERSETUP_1_0_ANTSERIALNBR_LENGTH = 20
const SBF_RECEIVERSETUP_1_0_ANTTYPE_LENGTH = 20
const SBF_RECEIVERSETUP_1_0_TOW_PRECISION = 3
const SBF_RECEIVERSETUP_1_0_DELTAH_PRECISION = 7
const SBF_RECEIVERSETUP_1_0_DELTAE_PRECISION = 7
const SBF_RECEIVERSETUP_1_0_DELTAN_PRECISION = 7

/*--ReceiverSetup_1_1_t : ---------------------------------------------------*/
/* General information about the receiver installation */
const SBF_RECEIVERSETUP_1_1_MARKERNAME_LENGTH = 60
const SBF_RECEIVERSETUP_1_1_MARKERNUMBER_LENGTH = 20
const SBF_RECEIVERSETUP_1_1_OBSERVER_LENGTH = 20
const SBF_RECEIVERSETUP_1_1_AGENCY_LENGTH = 40
const SBF_RECEIVERSETUP_1_1_RXSERIALNBR_LENGTH = 20
const SBF_RECEIVERSETUP_1_1_RXNAME_LENGTH = 20
const SBF_RECEIVERSETUP_1_1_RXVERSION_LENGTH = 20
const SBF_RECEIVERSETUP_1_1_ANTSERIALNBR_LENGTH = 20
const SBF_RECEIVERSETUP_1_1_ANTTYPE_LENGTH = 20
const SBF_RECEIVERSETUP_1_1_MARKERTYPE_LENGTH = 20
const SBF_RECEIVERSETUP_1_1_TOW_PRECISION = 3
const SBF_RECEIVERSETUP_1_1_DELTAH_PRECISION = 7
const SBF_RECEIVERSETUP_1_1_DELTAE_PRECISION = 7
const SBF_RECEIVERSETUP_1_1_DELTAN_PRECISION = 7

/*--ReceiverSetup_1_2_t : ---------------------------------------------------*/
/* General information about the receiver installation */
const SBF_RECEIVERSETUP_1_2_MARKERNAME_LENGTH = 60
const SBF_RECEIVERSETUP_1_2_MARKERNUMBER_LENGTH = 20
const SBF_RECEIVERSETUP_1_2_OBSERVER_LENGTH = 20
const SBF_RECEIVERSETUP_1_2_AGENCY_LENGTH = 40
const SBF_RECEIVERSETUP_1_2_RXSERIALNBR_LENGTH = 20
const SBF_RECEIVERSETUP_1_2_RXNAME_LENGTH = 20
const SBF_RECEIVERSETUP_1_2_RXVERSION_LENGTH = 20
const SBF_RECEIVERSETUP_1_2_ANTSERIALNBR_LENGTH = 20
const SBF_RECEIVERSETUP_1_2_ANTTYPE_LENGTH = 20
const SBF_RECEIVERSETUP_1_2_MARKERTYPE_LENGTH = 20
const SBF_RECEIVERSETUP_1_2_GNSSFWVERSION_LENGTH = 40
const SBF_RECEIVERSETUP_1_2_TOW_PRECISION = 3
const SBF_RECEIVERSETUP_1_2_DELTAH_PRECISION = 7
const SBF_RECEIVERSETUP_1_2_DELTAE_PRECISION = 7
const SBF_RECEIVERSETUP_1_2_DELTAN_PRECISION = 7

/*--ReceiverSetup_1_3_t : ---------------------------------------------------*/
/* General information about the receiver installation */
const SBF_RECEIVERSETUP_1_3_MARKERNAME_LENGTH = 60
const SBF_RECEIVERSETUP_1_3_MARKERNUMBER_LENGTH = 20
const SBF_RECEIVERSETUP_1_3_OBSERVER_LENGTH = 20
const SBF_RECEIVERSETUP_1_3_AGENCY_LENGTH = 40
const SBF_RECEIVERSETUP_1_3_RXSERIALNBR_LENGTH = 20
const SBF_RECEIVERSETUP_1_3_RXNAME_LENGTH = 20
const SBF_RECEIVERSETUP_1_3_RXVERSION_LENGTH = 20
const SBF_RECEIVERSETUP_1_3_ANTSERIALNBR_LENGTH = 20
const SBF_RECEIVERSETUP_1_3_ANTTYPE_LENGTH = 20
const SBF_RECEIVERSETUP_1_3_MARKERTYPE_LENGTH = 20
const SBF_RECEIVERSETUP_1_3_GNSSFWVERSION_LENGTH = 40
const SBF_RECEIVERSETUP_1_3_PRODUCTNAME_LENGTH = 40
const SBF_RECEIVERSETUP_1_3_TOW_PRECISION = 3
const SBF_RECEIVERSETUP_1_3_DELTAH_PRECISION = 7
const SBF_RECEIVERSETUP_1_3_DELTAE_PRECISION = 7
const SBF_RECEIVERSETUP_1_3_DELTAN_PRECISION = 7

/*--ReceiverSetup_1_4_t : ---------------------------------------------------*/
/* General information about the receiver installation */
const SBF_RECEIVERSETUP_1_4_MARKERNAME_LENGTH = 60
const SBF_RECEIVERSETUP_1_4_MARKERNUMBER_LENGTH = 20
const SBF_RECEIVERSETUP_1_4_OBSERVER_LENGTH = 20
const SBF_RECEIVERSETUP_1_4_AGENCY_LENGTH = 40
const SBF_RECEIVERSETUP_1_4_RXSERIALNBR_LENGTH = 20
const SBF_RECEIVERSETUP_1_4_RXNAME_LENGTH = 20
const SBF_RECEIVERSETUP_1_4_RXVERSION_LENGTH = 20
const SBF_RECEIVERSETUP_1_4_ANTSERIALNBR_LENGTH = 20
const SBF_RECEIVERSETUP_1_4_ANTTYPE_LENGTH = 20
const SBF_RECEIVERSETUP_1_4_MARKERTYPE_LENGTH = 20
const SBF_RECEIVERSETUP_1_4_GNSSFWVERSION_LENGTH = 40
const SBF_RECEIVERSETUP_1_4_PRODUCTNAME_LENGTH = 40
const SBF_RECEIVERSETUP_1_4_STATIONCODE_LENGTH = 10
const SBF_RECEIVERSETUP_1_4_COUNTRYCODE_LENGTH = 3
const SBF_RECEIVERSETUP_1_4_RESERVED1_LENGTH = 21
const SBF_RECEIVERSETUP_1_4_TOW_PRECISION = 3
const SBF_RECEIVERSETUP_1_4_DELTAH_PRECISION = 7
const SBF_RECEIVERSETUP_1_4_DELTAE_PRECISION = 7
const SBF_RECEIVERSETUP_1_4_DELTAN_PRECISION = 7
const SBF_RECEIVERSETUP_1_4_LATITUDE_PRECISION = 12
const SBF_RECEIVERSETUP_1_4_LONGITUDE_PRECISION = 12
const SBF_RECEIVERSETUP_1_4_HEIGHT_PRECISION = 3
const SBF_RECEIVERSETUP_1_MARKERNAME_LENGTH = SBF_RECEIVERSETUP_1_4_MARKERNAME_LENGTH
const SBF_RECEIVERSETUP_1_MARKERNUMBER_LENGTH = SBF_RECEIVERSETUP_1_4_MARKERNUMBER_LENGTH
const SBF_RECEIVERSETUP_1_OBSERVER_LENGTH = SBF_RECEIVERSETUP_1_4_OBSERVER_LENGTH
const SBF_RECEIVERSETUP_1_AGENCY_LENGTH = SBF_RECEIVERSETUP_1_4_AGENCY_LENGTH
const SBF_RECEIVERSETUP_1_RXSERIALNBR_LENGTH = SBF_RECEIVERSETUP_1_4_RXSERIALNBR_LENGTH
const SBF_RECEIVERSETUP_1_RXNAME_LENGTH = SBF_RECEIVERSETUP_1_4_RXNAME_LENGTH
const SBF_RECEIVERSETUP_1_RXVERSION_LENGTH = SBF_RECEIVERSETUP_1_4_RXVERSION_LENGTH
const SBF_RECEIVERSETUP_1_ANTSERIALNBR_LENGTH = SBF_RECEIVERSETUP_1_4_ANTSERIALNBR_LENGTH
const SBF_RECEIVERSETUP_1_ANTTYPE_LENGTH = SBF_RECEIVERSETUP_1_4_ANTTYPE_LENGTH
const SBF_RECEIVERSETUP_1_MARKERTYPE_LENGTH = SBF_RECEIVERSETUP_1_4_MARKERTYPE_LENGTH
const SBF_RECEIVERSETUP_1_GNSSFWVERSION_LENGTH = SBF_RECEIVERSETUP_1_4_GNSSFWVERSION_LENGTH
const SBF_RECEIVERSETUP_1_PRODUCTNAME_LENGTH = SBF_RECEIVERSETUP_1_4_PRODUCTNAME_LENGTH
const SBF_RECEIVERSETUP_1_STATIONCODE_LENGTH = SBF_RECEIVERSETUP_1_4_STATIONCODE_LENGTH
const SBF_RECEIVERSETUP_1_COUNTRYCODE_LENGTH = SBF_RECEIVERSETUP_1_4_COUNTRYCODE_LENGTH
const SBF_RECEIVERSETUP_1_RESERVED1_LENGTH = SBF_RECEIVERSETUP_1_4_RESERVED1_LENGTH

/*--RxComponents_1_0_t : ----------------------------------------------------*/
/* Information on various receiver components */
const SBF_COMPONENT_1_0_NAME_LENGTH = 40
const SBF_COMPONENT_1_0_SERIALNUMBER_LENGTH = 20
const SBF_COMPONENT_1_0_FWVERSION_LENGTH = 40
const SBF_COMPONENT_1_0_MACADDRESS_LENGTH = 6
const SBF_COMPONENT_1_0__PADDING_LENGTH = 2
const SBF_RXCOMPONENTS_1_0_COMPONENT_LENGTH = 20
const SBF_RXCOMPONENTS_1_0_TOW_PRECISION = 3
const SBF_COMPONENT_1_NAME_LENGTH = SBF_COMPONENT_1_0_NAME_LENGTH
const SBF_COMPONENT_1_SERIALNUMBER_LENGTH = SBF_COMPONENT_1_0_SERIALNUMBER_LENGTH
const SBF_COMPONENT_1_FWVERSION_LENGTH = SBF_COMPONENT_1_0_FWVERSION_LENGTH
const SBF_COMPONENT_1_MACADDRESS_LENGTH = SBF_COMPONENT_1_0_MACADDRESS_LENGTH
const SBF_COMPONENT_1__PADDING_LENGTH = SBF_COMPONENT_1_0__PADDING_LENGTH
const SBF_RXCOMPONENTS_1_COMPONENT_LENGTH = SBF_RXCOMPONENTS_1_0_COMPONENT_LENGTH

/*--RxMessage_1_0_t : -------------------------------------------------------*/
/* Receiver message */
const SBF_RXMESSAGE_1_0_RESERVED2_LENGTH = 2
const SBF_RXMESSAGE_1_0_MESSAGE_LENGTH = 2000
const SBF_RXMESSAGE_1_0_TOW_PRECISION = 3
const SBF_RXMESSAGE_1_RESERVED2_LENGTH = SBF_RXMESSAGE_1_0_RESERVED2_LENGTH
const SBF_RXMESSAGE_1_MESSAGE_LENGTH = SBF_RXMESSAGE_1_0_MESSAGE_LENGTH

/*--Commands_1_0_t : --------------------------------------------------------*/
/* Commands entered by the user */
const SBF_COMMANDS_1_0_CMDDATA_LENGTH = 2048
const SBF_COMMANDS_1_0_TOW_PRECISION = 3
const SBF_COMMANDS_1_CMDDATA_LENGTH = SBF_COMMANDS_1_0_CMDDATA_LENGTH

/*--Comment_1_0_t : ---------------------------------------------------------*/
/* Comment entered by the user */
const SBF_COMMENT_1_0_COMMENT_LENGTH = 120
const SBF_COMMENT_1_0_TOW_PRECISION = 3
const SBF_COMMENT_1_COMMENT_LENGTH = SBF_COMMENT_1_0_COMMENT_LENGTH

/*--BBSamples_1_0_t : -------------------------------------------------------*/
/* Baseband samples */
const SBF_BBSAMPLESDATA_1_0_SAMPLES_LENGTH = 2048
const SBF_BBSAMPLESDATA_1_0_TOW_PRECISION = 3
const SBF_BBSAMPLESDATA_1_SAMPLES_LENGTH = SBF_BBSAMPLESDATA_1_0_SAMPLES_LENGTH

/*--ASCIIIn_1_0_t : ---------------------------------------------------------*/
/* ASCII input from external sensor */
const SBF_ASCIIIN_1_0_SENSORMODEL_LENGTH = 20
const SBF_ASCIIIN_1_0_SENSORTYPE_LENGTH = 20
const SBF_ASCIIIN_1_0_ASCIISTRING_LENGTH = 2000
const SBF_ASCIIIN_1_0_TOW_PRECISION = 3
const SBF_ASCIIIN_1_SENSORMODEL_LENGTH = SBF_ASCIIIN_1_0_SENSORMODEL_LENGTH
const SBF_ASCIIIN_1_SENSORTYPE_LENGTH = SBF_ASCIIIN_1_0_SENSORTYPE_LENGTH
const SBF_ASCIIIN_1_ASCIISTRING_LENGTH = SBF_ASCIIIN_1_0_ASCIISTRING_LENGTH

/*--EncapsulatedOutput_1_0_t : ----------------------------------------------*/
/* SBF encapsulation of non-SBF messages */
const SBF_ENCAPSULATEDOUTPUT_1_0_PAYLOAD_LENGTH = 4096
const SBF_ENCAPSULATEDOUTPUT_1_0_TOW_PRECISION = 3
const SBF_ENCAPSULATEDOUTPUT_1_PAYLOAD_LENGTH = SBF_ENCAPSULATEDOUTPUT_1_0_PAYLOAD_LENGTH

/*--RawDataIn_1_0_t : -------------------------------------------------------*/
/* Incoming raw data message */
const SBF_RAWDATAIN_1_0_BYTES_LENGTH = 4096
const SBF_RAWDATAIN_1_0_TOW_PRECISION = 3
const SBF_RAWDATAIN_1_BYTES_LENGTH = SBF_RAWDATAIN_1_0_BYTES_LENGTH

/*==TUR Specific Blocks======================================================*/

/*--TURPVTSatCorrections_1_0_t : --------------------------------------------*/
/* Satellite range corrections */
const SBF_TURPVTSATCORRECTIONS_1_0_TOW_PRECISION = 3

/*--TURPVTSatCorrections_1_1_t : --------------------------------------------*/
/* Satellite range corrections */
const SBF_TURPVTSATCORRECTIONS_1_1_TOW_PRECISION = 3
const SBF_TURPVTSATCORRECTIONS_1_DATA_LENGTH = SBF_TURPVTSATCORRECTIONS_1_1_DATA_LENGTH

/*--TURHPCAInfo_1_0_t : -----------------------------------------------------*/
/* HMI Probability information */
const SBF_SATHPCA_1_0__PADDING_LENGTH = 1
const SBF_SATHPCA_1_0_HORDEV_PRECISION = 3
const SBF_SATHPCA_1_0_VERDEV_PRECISION = 3
const SBF_SATHPCA_1_0_HMIPROB_PRECISION = 3
const SBF_TURHPCAINFO_1_0_HPCADATA_LENGTH = 72
const SBF_TURHPCAINFO_1_0_TOW_PRECISION = 3
const SBF_TURHPCAINFO_1_0_HMIPROB_PRECISION = 3
const SBF_TURHPCAINFO_1_0_HORDEV_PRECISION = 3
const SBF_TURHPCAINFO_1_0_VERDEV_PRECISION = 3
const SBF_SATHPCA_1__PADDING_LENGTH = SBF_SATHPCA_1_0__PADDING_LENGTH
const SBF_TURHPCAINFO_1_HPCADATA_LENGTH = SBF_TURHPCAINFO_1_0_HPCADATA_LENGTH

/*--CorrPeakSample_1_0_t : --------------------------------------------------*/
/* Real-time samples of the correlation peak function */
const SBF_CORRPEAKSAMPLE_1_0_CORRSAMPLEDATA_LENGTH = 11
const SBF_CORRPEAKSAMPLE_1_0_TOW_PRECISION = 3
const SBF_CORRPEAKSAMPLE_1_CORRSAMPLEDATA_LENGTH = SBF_CORRPEAKSAMPLE_1_0_CORRSAMPLEDATA_LENGTH

/*--CorrValues_1_0_t : ------------------------------------------------------*/
/* Raw correlation values */
const SBF_CORRVALUES_1_0_CORRELATIONVALUES_LENGTH = 2000
const SBF_CORRVALUES_1_0__PADDING_LENGTH = 2
const SBF_CORRVALUES_1_0_TOW_PRECISION = 3
const SBF_CORRVALUES_1_CORRELATIONVALUES_LENGTH = SBF_CORRVALUES_1_0_CORRELATIONVALUES_LENGTH
const SBF_CORRVALUES_1__PADDING_LENGTH = SBF_CORRVALUES_1_0__PADDING_LENGTH

/*--TURStatus_1_0_t : -------------------------------------------------------*/
/* TUR-specific status information */
const SBF_TURSTATUS_1_0_TOW_PRECISION = 3

/*--TURStatus_1_1_t : -------------------------------------------------------*/
/* TUR-specific status information */
const SBF_TURSTATUS_1_1_TOW_PRECISION = 3

/*--TURStatus_1_2_t : -------------------------------------------------------*/
/* TUR-specific status information */
const SBF_TURSTATUS_1_2_TOW_PRECISION = 3

/*--GALIntegrity_1_0_t : ----------------------------------------------------*/
/* Galileo integrity data */
const SBF_GAINTEGRITY_1_0_INTEGRITYSVI_LENGTH = 36
const SBF_GAINTEGRITY_1_0_TOW_PRECISION = 3
const SBF_GAINTEGRITY_1_INTEGRITYSVI_LENGTH = SBF_GAINTEGRITY_1_0_INTEGRITYSVI_LENGTH

/*--TURFormat_1_0_t : -------------------------------------------------------*/
const SBF_TURFORMATMODE1RGCREQUEST_1_0__PADDING_LENGTH = 2
const SBF_TURFORMATRANGINGCODEREADY_1_0__PADDING_LENGTH = 2
const SBF_TURFORMATGETMODE2TIME_1_0__PADDING_LENGTH = 2
const SBF_TURFORMATMODE2TIMEREPLY_1_0__PADDING_LENGTH = 2
const SBF_TURFORMATENCRYPTEDDATA_1_0_UDATA_LENGTH = 41
const SBF_TURFORMATENCRYPTEDDATA_1_0__PADDING_LENGTH = 2
const SBF_TURFORMATDECRYPTEDDATA_1_0_UDATA_LENGTH = 21
const SBF_TURFORMATDECRYPTEDDATA_1_0__PADDING_LENGTH = 1
const SBF_TURFORMATDENIAL_1_0__PADDING_LENGTH = 2
const SBF_TURFORMATSTATUSRESPONSE_1_0__PADDING_LENGTH = 2
const SBF_TURFORMATOPTIONA_1_0__PADDING_LENGTH = 2
const SBF_TURFORMATOPTIONB_1_0_UMASK1_LENGTH = 5
const SBF_TURFORMATOPTIONB_1_0_UMASK2_LENGTH = 5
const SBF_TURFORMAT_1_0_TOW_PRECISION = 3
const SBF_TURFORMATMODE1RGCREQUEST_1__PADDING_LENGTH = SBF_TURFORMATMODE1RGCREQUEST_1_0__PADDING_LENGTH
const SBF_TURFORMATRANGINGCODEREADY_1__PADDING_LENGTH = SBF_TURFORMATRANGINGCODEREADY_1_0__PADDING_LENGTH
const SBF_TURFORMATGETMODE2TIME_1__PADDING_LENGTH = SBF_TURFORMATGETMODE2TIME_1_0__PADDING_LENGTH
const SBF_TURFORMATMODE2TIMEREPLY_1__PADDING_LENGTH = SBF_TURFORMATMODE2TIMEREPLY_1_0__PADDING_LENGTH
const SBF_TURFORMATRANGINGCODEREQUEST_1__PADDING_LENGTH = SBF_TURFORMATRANGINGCODEREQUEST_1_0__PADDING_LENGTH
const SBF_TURFORMATENCRYPTEDDATA_1_UDATA_LENGTH = SBF_TURFORMATENCRYPTEDDATA_1_0_UDATA_LENGTH
const SBF_TURFORMATENCRYPTEDDATA_1__PADDING_LENGTH = SBF_TURFORMATENCRYPTEDDATA_1_0__PADDING_LENGTH
const SBF_TURFORMATDECRYPTEDDATA_1_UDATA_LENGTH = SBF_TURFORMATDECRYPTEDDATA_1_0_UDATA_LENGTH
const SBF_TURFORMATDECRYPTEDDATA_1__PADDING_LENGTH = SBF_TURFORMATDECRYPTEDDATA_1_0__PADDING_LENGTH
const SBF_TURFORMATDENIAL_1__PADDING_LENGTH = SBF_TURFORMATDENIAL_1_0__PADDING_LENGTH
const SBF_TURFORMATSTATUSRESPONSE_1__PADDING_LENGTH = SBF_TURFORMATSTATUSRESPONSE_1_0__PADDING_LENGTH
const SBF_TURFORMATOPTIONA_1__PADDING_LENGTH = SBF_TURFORMATOPTIONA_1_0__PADDING_LENGTH
const SBF_TURFORMATOPTIONB_1_UMASK1_LENGTH = SBF_TURFORMATOPTIONB_1_0_UMASK1_LENGTH
const SBF_TURFORMATOPTIONB_1_UMASK2_LENGTH = SBF_TURFORMATOPTIONB_1_0_UMASK2_LENGTH
const SBF_TURFORMATRANGINGCODEREQUEST_1_0__PADDING_LENGTH = 3

/*--CalibrationValues_1_0_t : -----------------------------------------------*/
/* Calibration values */
const SBF_CALIBRATIONVALUES_1_0_CALIBRATIONINFO_LENGTH = 32
const SBF_CALIBRATIONVALUES_1_0_TOW_PRECISION = 3
const SBF_CALIBRATIONVALUES_1_0_TOW_LAST_PRECISION = 3
const SBF_CALIBRATIONVALUES_1_CALIBRATIONINFO_LENGTH = SBF_CALIBRATIONVALUES_1_0_CALIBRATIONINFO_LENGTH

/*--MultipathMonitor_1_0_t : ------------------------------------------------*/
/* MultipathMonitorSub blocks */
const SBF_MULTIPATHMONITOR_1_0_MULTIPATHMONITORSUB_LENGTH = 72
const SBF_MULTIPATHMONITOR_1_0_TOW_PRECISION = 3
const SBF_MULTIPATHMONITOR_1_MULTIPATHMONITORSUB_LENGTH = SBF_MULTIPATHMONITOR_1_0_MULTIPATHMONITORSUB_LENGTH

/*--FOCTURNStatus_1_0_t : ---------------------------------------------------*/
const SBF_FOCTURNSTATUS_1_0_TOW_PRECISION = 3

/*--TGVFXStatus_1_0_t : -----------------------------------------------------*/
const SBF_TGVFXSTATUS_1_0_TOW_PRECISION = 3

/*==PinPoint-GIS RX==========================================================*/

/*--GISAction_1_0_t : -------------------------------------------------------*/
/* PinPoint-GIS RX Action */
const SBF_GISACTION_1_0_COMMENT_LENGTH = 250
const SBF_GISACTION_1_0__PADDING_LENGTH = 2
const SBF_GISACTION_1_0_TOW_PRECISION = 3
const SBF_GISACTION_1_COMMENT_LENGTH = SBF_GISACTION_1_0_COMMENT_LENGTH
const SBF_GISACTION_1__PADDING_LENGTH = SBF_GISACTION_1_0__PADDING_LENGTH

/*--GISStatus_1_0_t : -------------------------------------------------------*/
/* Status of the different PinPoint-GIS collection databases */
const SBF_GISSTATUS_1_0_DATABASESTATUS_LENGTH = 3
const SBF_GISSTATUS_1_0_TOW_PRECISION = 3
const SBF_GISSTATUS_1_DATABASESTATUS_LENGTH = SBF_GISSTATUS_1_0_DATABASESTATUS_LENGTH

// compatibility definitions follow
//
// These constants are not currently being used. Commenting out for now.
//
// const sbfnr_GALSar_1 = sbfnr_GALSARRLM_1
// const sbfid_GALSar_1_0 = sbfid_GALSARRLM_1_0

// const sbfnr_PvtSatCartesian_1 = sbfnr_PVTSatCartesian_1
// const sbfid_PvtSatCartesian_1_0 = sbfid_PVTSatCartesian_1_0
// const sbfid_PvtSatCartesian_1_1 = sbfid_PVTSatCartesian_1_1

// const sbfnr_GeoCorrections_1 = sbfnr_GEOCorrections_1
// const sbfid_GeoCorrections_1_0 = sbfid_GEOCorrections_1_0

// const sbfnr_LBAS1Message_1 = sbfnr_LBAS1Messages_1
// const sbfid_LBAS1Message_1_0 = sbfid_LBAS1Messages_1_0

// const sbfnr_LbandBeams_1 = sbfnr_LBandBeams_1
// const sbfid_LbandBeams_1_0 = sbfid_LBandBeams_1_0

/*==SBF-BLOCKS definition====================================================*/

/*==Measurement Blocks=======================================================*/

/*--GenMeasEpoch_1_0_t : ----------------------------------------------------*/
const MAX_MEASEPOCH_DATA = SBF_GENMEASEPOCH_1_0_DATA_LENGTH

/*--MeasExtra_1_0_t : -------------------------------------------------------*/
const MAXSB_MEASEXTRA = SBF_MEASEXTRA_1_0_MEASEXTRACHANNEL_LENGTH

/*--IQCorr_1_0_t : ----------------------------------------------------------*/
const MAXSB_IQCORR = SBF_IQCORR_1_0_CORRCHANNEL_LENGTH

/*==Navigation Page Blocks===================================================*/

/*--GPSRaw_1_0_t : ----------------------------------------------------------*/
const GPS_NR_OF_RAW_BITS = SBF_GPSRAW_1_0_RAWBITS_LENGTH

/*--CNAVRaw_1_0_t : ---------------------------------------------------------*/
const CNAV_NR_OF_RAW_BITS = SBF_CNAVRAW_1_0_CNAVBITS_LENGTH

/*--GEORaw_1_0_t : ----------------------------------------------------------*/
const GEO_NR_OF_RAW_BITS = SBF_GEORAW_1_0_RAWBITS_LENGTH

/*--GPSRawCA_1_0_t : --------------------------------------------------------*/
const NB_NAVBITS_WORDS_GPSCA = SBF_GPSRAWCA_1_0_NAVBITS_LENGTH

/*--GPSRawL2C_1_0_t : -------------------------------------------------------*/
const NB_NAVBITS_WORDS_GPSL2C = SBF_GPSRAWL2C_1_0_NAVBITS_LENGTH

/*--GLORawCA_1_0_t : --------------------------------------------------------*/
const NB_NAVBITS_WORDS_GLOCA = SBF_GLORAWCA_1_0_NAVBITS_LENGTH

/*--GALRawFNAV_1_0_t : ------------------------------------------------------*/
const NB_NAVBITS_WORDS_GALFNAV = SBF_GALRAWFNAV_1_0_NAVBITS_LENGTH

/*--GALRawINAV_1_0_t : ------------------------------------------------------*/
const NB_NAVBITS_WORDS_GALINAV = SBF_GALRAWINAV_1_0_NAVBITS_LENGTH

/*--GALRawCNAV_1_0_t : ------------------------------------------------------*/
const NB_NAVBITS_WORDS_GALCNAV = SBF_GALRAWCNAV_1_0_NAVBITS_LENGTH

/*--GALRawGNAV_1_0_t : ------------------------------------------------------*/
const NB_NAVBITS_WORDS_GALGNAV = SBF_GALRAWGNAV_1_0_NAVBITS_LENGTH

/*--GALRawGNAVe_1_0_t : ------------------------------------------------------*/
const NB_NAVBITS_WORDS_GALGNAVE = SBF_GALRAWGNAVE_1_0_NAVBITS_LENGTH

/*--GEORawL1_1_0_t : --------------------------------------------------------*/
const NB_NAVBITS_WORDS_GEOL1 = SBF_GEORAWL1_1_0_NAVBITS_LENGTH

/*--CMPRaw_1_0_t : ----------------------------------------------------------*/
const NB_NAVBITS_WORDS_CMP = SBF_BDSRAW_1_0_NAVBITS_LENGTH

/* For backwards compatibility after renaming CMP to BDS */
const SBF_CMPRAW_1_NAVBITS_LENGTH = SBF_BDSRAW_1_NAVBITS_LENGTH
const SBF_CMPRAW_1_0_NAVBITS_LENGTH = SBF_BDSRAW_1_0_NAVBITS_LENGTH

//
// These constants are not currently being used. Commenting out for now.
//
// const sbfnr_CMPRaw_1 = sbfnr_BDSRaw_1
// const sbfid_CMPRaw_1_0 = sbfid_BDSRaw_1_0
// const sbfnr_CMPNav_1 = sbfnr_BDSNav_1
// const sbfid_CMPNav_1_0 = sbfid_BDSNav_1_0

// const sbfnr_IRNSSRaw_1 = sbfnr_NAVICRaw_1
// const sbfid_IRNSSRaw_1_0 = sbfid_NAVICRaw_1_0

/*==Galileo Decoded Message Blocks===========================================*/

/*--GALSARRLM_1_0_t : -------------------------------------------------------*/
const MAX_GAL_SAR_RLM_DATA = SBF_GASARRLM_1_0_RLMBITS_LENGTH

/*==SBAS Decoded Message Blocks==============================================*/

/*--GEOPRNMask_1_0_t : ------------------------------------------------------*/
const MAXSB_SBAS_NRCORRSAT = SBF_RAPRNMASK_1_0_PRNMASK_LENGTH

// same as SBF_GEOINTEGRITY_1_0_UDREI_LENGTH
// same as SBF_GEOFASTCORRDEGR_1_0_AI_LENGTH

/*--GEOIntegrity_1_0_t : ----------------------------------------------------*/
const MAXSB_SBAS_NRIODF = SBF_RAINTEGRITY_1_0_IODF_LENGTH

/*==Position, Velocity and Time Blocks=======================================*/

/*--PVTSatCartesian_1_0_t : -------------------------------------------------*/
const MAXSB_SATPOS = SBF_PVTSATCARTESIAN_1_0_SATPOS_LENGTH

/*--PVTResiduals_1_0_t : ----------------------------------------------------*/
const MAXSB_SATRESIDUAL = SBF_PVTRESIDUALS_1_0_SATRESIDUAL_LENGTH

/*--PVTResiduals_2_0_t : ----------------------------------------------------*/
const MAX_PVTRES_DATA = SBF_PVTRESIDUALS_2_0_DATA_LENGTH

/*--RAIMStatistics_1_0_t : --------------------------------------------------*/
const MAXSB_RAIM = SBF_RAIMSTATISTICS_1_0_RAIMCHANNEL_LENGTH

/*--GEOCorrections_1_0_t : --------------------------------------------------*/
const MAXSB_GEOCORRCHANNEL = SBF_GEOCORRECTIONS_1_0_GEOCORRCHANNEL_LENGTH

/*--BaseVectorCart_1_0_t : --------------------------------------------------*/
const MAXSB_VECTORINFO = SBF_BASEVECTORCART_1_0_VECTORINFO_LENGTH

// same as SBF_BASEVECTORGEOD_1_0_VECTORINFOGEOD_LENGTH

/*==Differential Correction Blocks===========================================*/

/*--RTCMDatum_1_0_t : -------------------------------------------------------*/
const RTCMDATUM_CRS_NAME_LENGTH = SBF_RTCMDATUM_1_0_SOURCECRS_LENGTH

// same as SBF_RTCMDATUM_1_0_TARGETCRS_LENGTH

/*==L-Band Demodulator Blocks================================================*/

/*--LBandTrackerStatus_1_0_t : ----------------------------------------------*/
const MAX_NROF_LBANDTRACKERDATA = SBF_LBANDTRACKERSTATUS_1_0_TRACKDATA_LENGTH

/*--LBAS1DecoderStatus_1_1_t : ----------------------------------------------*/
const LBAS1_PAC_LENGTH = SBF_LBAS1DECODERSTATUS_1_1_PAC_LENGTH

/*--LBAS1Message_1_0_t : ----------------------------------------------------*/
const MAX_LBAS1MSG_LEN = SBF_LBAS1MESSAGES_1_0_MESSAGE_LENGTH

/*--LBandBeams_1_0_t : ------------------------------------------------------*/
const MAX_LBANDSAT = SBF_BEAMINFO_1_0_SATNAME_LENGTH

/*==External Sensor Blocks===================================================*/

/*--ExtSensorMeas_1_0_t : ---------------------------------------------------*/
const MAXSB_EXTSENSORMEAS = SBF_EXTSENSORMEAS_1_0_EXTSENSORMEAS_LENGTH

/*--ExtSensorStatus_1_0_t : -------------------------------------------------*/
const EXTSENSORSTATUS_STATUSBITS_SIZE = SBF_EXTSENSORSTATUS_1_0_STATUSBITS_LENGTH

/*--ExtSensorSetup_1_0_t : --------------------------------------------------*/
const MAXSB_EXTSENSORSETUP = SBF_EXTSENSORSETUP_1_0_EXTSENSORSETUP_LENGTH

/*--ExtSensorStatus_2_0_t : -------------------------------------------------*/
//const EXTSENSORSTATUS2_STATUSBITS_SIZE = SBF_EXTSENSORSTATUS_2_0_STATUSBITS_LENGTH

/*--ExtSensorInfo_1_0_t : ---------------------------------------------------*/
//const EXTSENSORINFO_STATUSBITS_SIZE = SBF_EXTSENSORINFO_1_0_STATUSBITS_LENGTH

/*==Status Blocks============================================================*/

/*--TrackingStatus_1_0_t : --------------------------------------------------*/
const MAXSB_TRACKINGSTATUS = SBF_TRACKINGSTATUS_1_0_CHANNELDATA_LENGTH

/*--ChannelStatus_1_0_t : ---------------------------------------------------*/
const MAXSB_CHANNELSTS = SBF_CHANNELSTATUS_1_0_DATA_LENGTH

/*--ReceiverStatus_2_0_t : --------------------------------------------------*/
const MAXSB_AGCSTATE = SBF_RECEIVERSTATUS_2_0_AGCSTATE_LENGTH

/*--SatVisibility_1_0_t : ---------------------------------------------------*/
const MAXSB_SATVISIBILITY = SBF_SATVISIBILITY_1_0_SATINFO_LENGTH

/*--InputLink_1_0_t : -------------------------------------------------------*/
const MAX_INPUTSTATS_DATA = SBF_INPUTLINK_1_0_DATA_LENGTH

/*--OutputLink_1_0_t : ------------------------------------------------------*/
const MAX_OUTPUTSTATS_DATA_1_0 = SBF_OUTPUTLINK_1_0_DATA_LENGTH

/*--OutputLink_1_1_t : ------------------------------------------------------*/
const MAX_OUTPUTSTATS_DATA_1_1 = SBF_OUTPUTLINK_1_1_DATA_LENGTH

/*--IPStatus_1_0_t : --------------------------------------------------------*/
const MAC_ADDR_LENGTH = SBF_IPSTATUS_1_0_MACADDRESS_LENGTH
const IP_ADDR_LENGTH = SBF_IPSTATUS_1_0_IPADDRESS_LENGTH

// same as SBF_IPSTATUS_1_0_GATEWAY_LENGTH

/*==Miscellaneous Blocks=====================================================*/
/*--ReceiverSetup_1_0_t : ---------------------------------------------------*/
const MARKERNAME_LENGTH = SBF_RECEIVERSETUP_1_0_MARKERNAME_LENGTH
const MARKERNBR_LENGTH = SBF_RECEIVERSETUP_1_0_MARKERNUMBER_LENGTH
const OBSERVER_LENGTH = SBF_RECEIVERSETUP_1_0_OBSERVER_LENGTH
const AGENCY_LENGTH = SBF_RECEIVERSETUP_1_0_AGENCY_LENGTH
const RXSERIALNBR_LENGTH = SBF_RECEIVERSETUP_1_0_RXSERIALNBR_LENGTH
const RXNAME_LENGTH = SBF_RECEIVERSETUP_1_0_RXNAME_LENGTH
const RXVERSION_LENGTH = SBF_RECEIVERSETUP_1_0_RXVERSION_LENGTH
const ANTSERIALNBR_LENGTH = SBF_RECEIVERSETUP_1_0_ANTSERIALNBR_LENGTH
const ANTTYPE_LENGTH = SBF_RECEIVERSETUP_1_0_ANTTYPE_LENGTH

/*--ReceiverSetup_1_1_t : ---------------------------------------------------*/
const MARKERTYPE_LENGTH = SBF_RECEIVERSETUP_1_1_MARKERTYPE_LENGTH

/*--ReceiverSetup_1_2_t : ---------------------------------------------------*/
const GNSSFWVERSION_LENGTH = SBF_RECEIVERSETUP_1_2_GNSSFWVERSION_LENGTH

/*--ReceiverSetup_1_3_t : ---------------------------------------------------*/
const PRODUCTNAME_LENGTH = SBF_RECEIVERSETUP_1_3_PRODUCTNAME_LENGTH

/*--ReceiverSetup_1_4_t : ---------------------------------------------------*/
const STATIONCODE_LENGTH = SBF_RECEIVERSETUP_1_4_STATIONCODE_LENGTH
const COUNTRYCODE_LENGTH = SBF_RECEIVERSETUP_1_4_COUNTRYCODE_LENGTH

/*--Commands_1_0_t : --------------------------------------------------------*/
const COMMANDS_LENGTH = SBF_COMMANDS_1_0_CMDDATA_LENGTH

/*--Comment_1_0_t : ---------------------------------------------------------*/
const COMMENT_LENGTH = SBF_COMMENT_1_0_COMMENT_LENGTH

/*--BBSamples_1_0_t : -------------------------------------------------------*/
const MAX_BBSAMPLES = SBF_BBSAMPLESDATA_1_0_SAMPLES_LENGTH

/*--ASCIIIn_1_0_t : ---------------------------------------------------------*/
const ASCII_STRING_LENGTH = SBF_ASCIIIN_1_0_ASCIISTRING_LENGTH

/*==TUR Specific Blocks======================================================*/

/*--TURPVTSatCorrections_1_0_t : ---------------------------------------------*/
const MAX_PVT_SATCORRECTIONS_DATA = SBF_TURPVTSATCORRECTIONS_1_0_DATA_LENGTH

/*--TURHPCAInfo_1_0_t : -----------------------------------------------------*/
const MAXSB_SATHPCA = SBF_TURHPCAINFO_1_0_HPCADATA_LENGTH

/*--CorrPeakSample_1_0_t : --------------------------------------------------*/
const MAXSB_CORRSAMPLE = SBF_CORRPEAKSAMPLE_1_0_CORRSAMPLEDATA_LENGTH

/*--GALIntegrity_1_0_t : ----------------------------------------------------*/
const MAX_GALINTEGRITY = SBF_GAINTEGRITY_1_0_INTEGRITYSVI_LENGTH

/*--SysTimeOffset_1_0_t : ---------------------------------------------------*/
/*--SysTimeOffset_1_1_t : ---------------------------------------------------*/
const MAXSB_TIMEOFFSET = SBF_SYSTIMEOFFSET_1_0_TIMEOFFSET_LENGTH

/*==Other Blocks=============================================================*/

/*--TrackerDatat_1_0_t : ----------------------------------------------------*/
//const SBF_TRACKERDATA_MAXSIZE = SBF_TRACKERDATA_1_0_DATA_LENGTH
//const SBF_TRACKERDATA_MAXMSGS = SBF_TRACKERDATA_1_0_SUBMSGINDEX_LENGTH
