// Peripheral: PWR_Periph  Power Control.
// Instances:
//  PWR  mmap.PWR_BASE
// Registers:
//  0x00 32  CR1   Power control register 1.
//  0x04 32  CR2   Power control register 2.
//  0x08 32  CR3   Power control register 3.
//  0x0C 32  CR4   Power control register 4.
//  0x10 32  SR1   Power status register 1.
//  0x14 32  SR2   Power status register 2.
//  0x18 32  SCR   Power status reset register.
//  0x20 32  PUCRA Pull_up control register of portA.
//  0x24 32  PDCRA Pull_Down control register of portA.
//  0x28 32  PUCRB Pull_up control register of portB.
//  0x2C 32  PDCRB Pull_Down control register of portB.
//  0x30 32  PUCRC Pull_up control register of portC.
//  0x34 32  PDCRC Pull_Down control register of portC.
//  0x38 32  PUCRD Pull_up control register of portD.
//  0x3C 32  PDCRD Pull_Down control register of portD.
//  0x40 32  PUCRE Pull_up control register of portE.
//  0x44 32  PDCRE Pull_Down control register of portE.
//  0x48 32  PUCRF Pull_up control register of portF.
//  0x4C 32  PDCRF Pull_Down control register of portF.
//  0x50 32  PUCRG Pull_up control register of portG.
//  0x54 32  PDCRG Pull_Down control register of portG.
//  0x58 32  PUCRH Pull_up control register of portH.
//  0x5C 32  PDCRH Pull_Down control register of portH.
// Import:
//  stm32/o/l476xx/mmap
package pwr

// DO NOT EDIT THIS FILE. GENERATED BY stm32xgen.

const (
	LPR           CR1 = 0x01 << 14 //+ Regulator low-power mode.
	VOS           CR1 = 0x03 << 9  //+ VOS[1:0] bits (Regulator voltage scaling output selection).
	DBP           CR1 = 0x01 << 8  //+ Disable Back-up domain Protection.
	LPMS          CR1 = 0x07 << 0  //+ Low-power mode selection field.
	LPMS_STOP0    CR1 = 0x00 << 0  //  Stop 0 mode.
	LPMS_STOP1    CR1 = 0x01 << 0  //  Stop 1 mode.
	LPMS_STOP2    CR1 = 0x02 << 0  //  Stop 2 mode.
	LPMS_STANDBY  CR1 = 0x03 << 0  //  Stand-by mode.
	LPMS_SHUTDOWN CR1 = 0x04 << 0  //  Shut-down mode.
)

const (
	LPRn  = 14
	VOSn  = 9
	DBPn  = 8
	LPMSn = 0
)

const (
	USV      CR2 = 0x01 << 10 //+ VDD USB Supply Valid.
	IOSV     CR2 = 0x01 << 9  //+ VDD IO2 independent I/Os Supply Valid.
	PVME     CR2 = 0x0F << 4  //+ PVM bits field.
	PVME4    CR2 = 0x08 << 4  //  PVM 4 Enable.
	PVME3    CR2 = 0x04 << 4  //  PVM 3 Enable.
	PVME2    CR2 = 0x02 << 4  //  PVM 2 Enable.
	PVME1    CR2 = 0x01 << 4  //  PVM 1 Enable.
	PLS      CR2 = 0x07 << 1  //+ PVD level selection.
	PLS_LEV0 CR2 = 0x00 << 1  //  PVD level 0.
	PLS_LEV1 CR2 = 0x01 << 1  //  PVD level 1.
	PLS_LEV2 CR2 = 0x02 << 1  //  PVD level 2.
	PLS_LEV3 CR2 = 0x03 << 1  //  PVD level 3.
	PLS_LEV4 CR2 = 0x04 << 1  //  PVD level 4.
	PLS_LEV5 CR2 = 0x05 << 1  //  PVD level 5.
	PLS_LEV6 CR2 = 0x06 << 1  //  PVD level 6.
	PLS_LEV7 CR2 = 0x07 << 1  //  PVD level 7.
	PVDE     CR2 = 0x01 << 0  //+ Power Voltage Detector Enable.
)

const (
	USVn  = 10
	IOSVn = 9
	PVMEn = 4
	PLSn  = 1
	PVDEn = 0
)

const (
	EIWF  CR3 = 0x01 << 15 //+ Enable Internal Wake-up line.
	APC   CR3 = 0x01 << 10 //+ Apply pull-up and pull-down configuration.
	RRS   CR3 = 0x01 << 8  //+ SRAM2 Retention in Stand-by mode.
	EWUP5 CR3 = 0x01 << 4  //+ Enable Wake-Up Pin 5.
	EWUP4 CR3 = 0x01 << 3  //+ Enable Wake-Up Pin 4.
	EWUP3 CR3 = 0x01 << 2  //+ Enable Wake-Up Pin 3.
	EWUP2 CR3 = 0x01 << 1  //+ Enable Wake-Up Pin 2.
	EWUP1 CR3 = 0x01 << 0  //+ Enable Wake-Up Pin 1.
	EWUP  CR3 = 0x1F << 0  //  Enable Wake-Up Pins.
)

const (
	EIWFn  = 15
	APCn   = 10
	RRSn   = 8
	EWUP5n = 4
	EWUP4n = 3
	EWUP3n = 2
	EWUP2n = 1
	EWUP1n = 0
)

const (
	VBRS CR4 = 0x01 << 9 //+ VBAT Battery charging Resistor Selection.
	VBE  CR4 = 0x01 << 8 //+ VBAT Battery charging Enable.
	WP5  CR4 = 0x01 << 4 //+ Wake-Up Pin 5 polarity.
	WP4  CR4 = 0x01 << 3 //+ Wake-Up Pin 4 polarity.
	WP3  CR4 = 0x01 << 2 //+ Wake-Up Pin 3 polarity.
	WP2  CR4 = 0x01 << 1 //+ Wake-Up Pin 2 polarity.
	WP1  CR4 = 0x01 << 0 //+ Wake-Up Pin 1 polarity.
)

const (
	VBRSn = 9
	VBEn  = 8
	WP5n  = 4
	WP4n  = 3
	WP3n  = 2
	WP2n  = 1
	WP1n  = 0
)

const (
	WUFI SR1 = 0x01 << 15 //+ Wake-Up Flag Internal.
	SBF  SR1 = 0x01 << 8  //+ Stand-By Flag.
	WUF  SR1 = 0x1F << 0  //+ Wake-up Flags.
	WUF5 SR1 = 0x10 << 0  //  Wake-up Flag 5.
	WUF4 SR1 = 0x08 << 0  //  Wake-up Flag 4.
	WUF3 SR1 = 0x04 << 0  //  Wake-up Flag 3.
	WUF2 SR1 = 0x02 << 0  //  Wake-up Flag 2.
	WUF1 SR1 = 0x01 << 0  //  Wake-up Flag 1.
)

const (
	WUFIn = 15
	SBFn  = 8
	WUFn  = 0
)

const (
	PVMO4  SR2 = 0x01 << 15 //+ Peripheral Voltage Monitoring Output 4.
	PVMO3  SR2 = 0x01 << 14 //+ Peripheral Voltage Monitoring Output 3.
	PVMO2  SR2 = 0x01 << 13 //+ Peripheral Voltage Monitoring Output 2.
	PVMO1  SR2 = 0x01 << 12 //+ Peripheral Voltage Monitoring Output 1.
	PVDO   SR2 = 0x01 << 11 //+ Power Voltage Detector Output.
	VOSF   SR2 = 0x01 << 10 //+ Voltage Scaling Flag.
	REGLPF SR2 = 0x01 << 9  //+ Low-power Regulator Flag.
	REGLPS SR2 = 0x01 << 8  //+ Low-power Regulator Started.
)

const (
	PVMO4n  = 15
	PVMO3n  = 14
	PVMO2n  = 13
	PVMO1n  = 12
	PVDOn   = 11
	VOSFn   = 10
	REGLPFn = 9
	REGLPSn = 8
)

const (
	CSBF  SCR = 0x01 << 8 //+ Clear Stand-By Flag.
	CWUF  SCR = 0x1F << 0 //+ Clear Wake-up Flags.
	CWUF5 SCR = 0x10 << 0 //  Clear Wake-up Flag 5.
	CWUF4 SCR = 0x08 << 0 //  Clear Wake-up Flag 4.
	CWUF3 SCR = 0x04 << 0 //  Clear Wake-up Flag 3.
	CWUF2 SCR = 0x02 << 0 //  Clear Wake-up Flag 2.
	CWUF1 SCR = 0x01 << 0 //  Clear Wake-up Flag 1.
)

const (
	CSBFn = 8
	CWUFn = 0
)

const (
	PUPA15 PUCRA = 0x01 << 15 //+ Port PA15 Pull-Up set.
	PUPA13 PUCRA = 0x01 << 13 //+ Port PA13 Pull-Up set.
	PUPA12 PUCRA = 0x01 << 12 //+ Port PA12 Pull-Up set.
	PUPA11 PUCRA = 0x01 << 11 //+ Port PA11 Pull-Up set.
	PUPA10 PUCRA = 0x01 << 10 //+ Port PA10 Pull-Up set.
	PUPA9  PUCRA = 0x01 << 9  //+ Port PA9 Pull-Up set.
	PUPA8  PUCRA = 0x01 << 8  //+ Port PA8 Pull-Up set.
	PUPA7  PUCRA = 0x01 << 7  //+ Port PA7 Pull-Up set.
	PUPA6  PUCRA = 0x01 << 6  //+ Port PA6 Pull-Up set.
	PUPA5  PUCRA = 0x01 << 5  //+ Port PA5 Pull-Up set.
	PUPA4  PUCRA = 0x01 << 4  //+ Port PA4 Pull-Up set.
	PUPA3  PUCRA = 0x01 << 3  //+ Port PA3 Pull-Up set.
	PUPA2  PUCRA = 0x01 << 2  //+ Port PA2 Pull-Up set.
	PUPA1  PUCRA = 0x01 << 1  //+ Port PA1 Pull-Up set.
	PUPA0  PUCRA = 0x01 << 0  //+ Port PA0 Pull-Up set.
)

const (
	PUPA15n = 15
	PUPA13n = 13
	PUPA12n = 12
	PUPA11n = 11
	PUPA10n = 10
	PUPA9n  = 9
	PUPA8n  = 8
	PUPA7n  = 7
	PUPA6n  = 6
	PUPA5n  = 5
	PUPA4n  = 4
	PUPA3n  = 3
	PUPA2n  = 2
	PUPA1n  = 1
	PUPA0n  = 0
)

const (
	PDPA14 PDCRA = 0x01 << 14 //+ Port PA14 Pull-Down set.
	PDPA12 PDCRA = 0x01 << 12 //+ Port PA12 Pull-Down set.
	PDPA11 PDCRA = 0x01 << 11 //+ Port PA11 Pull-Down set.
	PDPA10 PDCRA = 0x01 << 10 //+ Port PA10 Pull-Down set.
	PDPA9  PDCRA = 0x01 << 9  //+ Port PA9 Pull-Down set.
	PDPA8  PDCRA = 0x01 << 8  //+ Port PA8 Pull-Down set.
	PDPA7  PDCRA = 0x01 << 7  //+ Port PA7 Pull-Down set.
	PDPA6  PDCRA = 0x01 << 6  //+ Port PA6 Pull-Down set.
	PDPA5  PDCRA = 0x01 << 5  //+ Port PA5 Pull-Down set.
	PDPA4  PDCRA = 0x01 << 4  //+ Port PA4 Pull-Down set.
	PDPA3  PDCRA = 0x01 << 3  //+ Port PA3 Pull-Down set.
	PDPA2  PDCRA = 0x01 << 2  //+ Port PA2 Pull-Down set.
	PDPA1  PDCRA = 0x01 << 1  //+ Port PA1 Pull-Down set.
	PDPA0  PDCRA = 0x01 << 0  //+ Port PA0 Pull-Down set.
)

const (
	PDPA14n = 14
	PDPA12n = 12
	PDPA11n = 11
	PDPA10n = 10
	PDPA9n  = 9
	PDPA8n  = 8
	PDPA7n  = 7
	PDPA6n  = 6
	PDPA5n  = 5
	PDPA4n  = 4
	PDPA3n  = 3
	PDPA2n  = 2
	PDPA1n  = 1
	PDPA0n  = 0
)

const (
	PUPB15 PUCRB = 0x01 << 15 //+ Port PB15 Pull-Up set.
	PUPB14 PUCRB = 0x01 << 14 //+ Port PB14 Pull-Up set.
	PUPB13 PUCRB = 0x01 << 13 //+ Port PB13 Pull-Up set.
	PUPB12 PUCRB = 0x01 << 12 //+ Port PB12 Pull-Up set.
	PUPB11 PUCRB = 0x01 << 11 //+ Port PB11 Pull-Up set.
	PUPB10 PUCRB = 0x01 << 10 //+ Port PB10 Pull-Up set.
	PUPB9  PUCRB = 0x01 << 9  //+ Port PB9 Pull-Up set.
	PUPB8  PUCRB = 0x01 << 8  //+ Port PB8 Pull-Up set.
	PUPB7  PUCRB = 0x01 << 7  //+ Port PB7 Pull-Up set.
	PUPB6  PUCRB = 0x01 << 6  //+ Port PB6 Pull-Up set.
	PUPB5  PUCRB = 0x01 << 5  //+ Port PB5 Pull-Up set.
	PUPB4  PUCRB = 0x01 << 4  //+ Port PB4 Pull-Up set.
	PUPB3  PUCRB = 0x01 << 3  //+ Port PB3 Pull-Up set.
	PUPB2  PUCRB = 0x01 << 2  //+ Port PB2 Pull-Up set.
	PUPB1  PUCRB = 0x01 << 1  //+ Port PB1 Pull-Up set.
	PUPB0  PUCRB = 0x01 << 0  //+ Port PB0 Pull-Up set.
)

const (
	PUPB15n = 15
	PUPB14n = 14
	PUPB13n = 13
	PUPB12n = 12
	PUPB11n = 11
	PUPB10n = 10
	PUPB9n  = 9
	PUPB8n  = 8
	PUPB7n  = 7
	PUPB6n  = 6
	PUPB5n  = 5
	PUPB4n  = 4
	PUPB3n  = 3
	PUPB2n  = 2
	PUPB1n  = 1
	PUPB0n  = 0
)

const (
	PDPB15 PDCRB = 0x01 << 15 //+ Port PB15 Pull-Down set.
	PDPB14 PDCRB = 0x01 << 14 //+ Port PB14 Pull-Down set.
	PDPB13 PDCRB = 0x01 << 13 //+ Port PB13 Pull-Down set.
	PDPB12 PDCRB = 0x01 << 12 //+ Port PB12 Pull-Down set.
	PDPB11 PDCRB = 0x01 << 11 //+ Port PB11 Pull-Down set.
	PDPB10 PDCRB = 0x01 << 10 //+ Port PB10 Pull-Down set.
	PDPB9  PDCRB = 0x01 << 9  //+ Port PB9 Pull-Down set.
	PDPB8  PDCRB = 0x01 << 8  //+ Port PB8 Pull-Down set.
	PDPB7  PDCRB = 0x01 << 7  //+ Port PB7 Pull-Down set.
	PDPB6  PDCRB = 0x01 << 6  //+ Port PB6 Pull-Down set.
	PDPB5  PDCRB = 0x01 << 5  //+ Port PB5 Pull-Down set.
	PDPB3  PDCRB = 0x01 << 3  //+ Port PB3 Pull-Down set.
	PDPB2  PDCRB = 0x01 << 2  //+ Port PB2 Pull-Down set.
	PDPB1  PDCRB = 0x01 << 1  //+ Port PB1 Pull-Down set.
	PDPB0  PDCRB = 0x01 << 0  //+ Port PB0 Pull-Down set.
)

const (
	PDPB15n = 15
	PDPB14n = 14
	PDPB13n = 13
	PDPB12n = 12
	PDPB11n = 11
	PDPB10n = 10
	PDPB9n  = 9
	PDPB8n  = 8
	PDPB7n  = 7
	PDPB6n  = 6
	PDPB5n  = 5
	PDPB3n  = 3
	PDPB2n  = 2
	PDPB1n  = 1
	PDPB0n  = 0
)

const (
	PUPC15 PUCRC = 0x01 << 15 //+ Port PC15 Pull-Up set.
	PUPC14 PUCRC = 0x01 << 14 //+ Port PC14 Pull-Up set.
	PUPC13 PUCRC = 0x01 << 13 //+ Port PC13 Pull-Up set.
	PUPC12 PUCRC = 0x01 << 12 //+ Port PC12 Pull-Up set.
	PUPC11 PUCRC = 0x01 << 11 //+ Port PC11 Pull-Up set.
	PUPC10 PUCRC = 0x01 << 10 //+ Port PC10 Pull-Up set.
	PUPC9  PUCRC = 0x01 << 9  //+ Port PC9 Pull-Up set.
	PUPC8  PUCRC = 0x01 << 8  //+ Port PC8 Pull-Up set.
	PUPC7  PUCRC = 0x01 << 7  //+ Port PC7 Pull-Up set.
	PUPC6  PUCRC = 0x01 << 6  //+ Port PC6 Pull-Up set.
	PUPC5  PUCRC = 0x01 << 5  //+ Port PC5 Pull-Up set.
	PUPC4  PUCRC = 0x01 << 4  //+ Port PC4 Pull-Up set.
	PUPC3  PUCRC = 0x01 << 3  //+ Port PC3 Pull-Up set.
	PUPC2  PUCRC = 0x01 << 2  //+ Port PC2 Pull-Up set.
	PUPC1  PUCRC = 0x01 << 1  //+ Port PC1 Pull-Up set.
	PUPC0  PUCRC = 0x01 << 0  //+ Port PC0 Pull-Up set.
)

const (
	PUPC15n = 15
	PUPC14n = 14
	PUPC13n = 13
	PUPC12n = 12
	PUPC11n = 11
	PUPC10n = 10
	PUPC9n  = 9
	PUPC8n  = 8
	PUPC7n  = 7
	PUPC6n  = 6
	PUPC5n  = 5
	PUPC4n  = 4
	PUPC3n  = 3
	PUPC2n  = 2
	PUPC1n  = 1
	PUPC0n  = 0
)

const (
	PDPC15 PDCRC = 0x01 << 15 //+ Port PC15 Pull-Down set.
	PDPC14 PDCRC = 0x01 << 14 //+ Port PC14 Pull-Down set.
	PDPC13 PDCRC = 0x01 << 13 //+ Port PC13 Pull-Down set.
	PDPC12 PDCRC = 0x01 << 12 //+ Port PC12 Pull-Down set.
	PDPC11 PDCRC = 0x01 << 11 //+ Port PC11 Pull-Down set.
	PDPC10 PDCRC = 0x01 << 10 //+ Port PC10 Pull-Down set.
	PDPC9  PDCRC = 0x01 << 9  //+ Port PC9 Pull-Down set.
	PDPC8  PDCRC = 0x01 << 8  //+ Port PC8 Pull-Down set.
	PDPC7  PDCRC = 0x01 << 7  //+ Port PC7 Pull-Down set.
	PDPC6  PDCRC = 0x01 << 6  //+ Port PC6 Pull-Down set.
	PDPC5  PDCRC = 0x01 << 5  //+ Port PC5 Pull-Down set.
	PDPC4  PDCRC = 0x01 << 4  //+ Port PC4 Pull-Down set.
	PDPC3  PDCRC = 0x01 << 3  //+ Port PC3 Pull-Down set.
	PDPC2  PDCRC = 0x01 << 2  //+ Port PC2 Pull-Down set.
	PDPC1  PDCRC = 0x01 << 1  //+ Port PC1 Pull-Down set.
	PDPC0  PDCRC = 0x01 << 0  //+ Port PC0 Pull-Down set.
)

const (
	PDPC15n = 15
	PDPC14n = 14
	PDPC13n = 13
	PDPC12n = 12
	PDPC11n = 11
	PDPC10n = 10
	PDPC9n  = 9
	PDPC8n  = 8
	PDPC7n  = 7
	PDPC6n  = 6
	PDPC5n  = 5
	PDPC4n  = 4
	PDPC3n  = 3
	PDPC2n  = 2
	PDPC1n  = 1
	PDPC0n  = 0
)

const (
	PUPD15 PUCRD = 0x01 << 15 //+ Port PD15 Pull-Up set.
	PUPD14 PUCRD = 0x01 << 14 //+ Port PD14 Pull-Up set.
	PUPD13 PUCRD = 0x01 << 13 //+ Port PD13 Pull-Up set.
	PUPD12 PUCRD = 0x01 << 12 //+ Port PD12 Pull-Up set.
	PUPD11 PUCRD = 0x01 << 11 //+ Port PD11 Pull-Up set.
	PUPD10 PUCRD = 0x01 << 10 //+ Port PD10 Pull-Up set.
	PUPD9  PUCRD = 0x01 << 9  //+ Port PD9 Pull-Up set.
	PUPD8  PUCRD = 0x01 << 8  //+ Port PD8 Pull-Up set.
	PUPD7  PUCRD = 0x01 << 7  //+ Port PD7 Pull-Up set.
	PUPD6  PUCRD = 0x01 << 6  //+ Port PD6 Pull-Up set.
	PUPD5  PUCRD = 0x01 << 5  //+ Port PD5 Pull-Up set.
	PUPD4  PUCRD = 0x01 << 4  //+ Port PD4 Pull-Up set.
	PUPD3  PUCRD = 0x01 << 3  //+ Port PD3 Pull-Up set.
	PUPD2  PUCRD = 0x01 << 2  //+ Port PD2 Pull-Up set.
	PUPD1  PUCRD = 0x01 << 1  //+ Port PD1 Pull-Up set.
	PUPD0  PUCRD = 0x01 << 0  //+ Port PD0 Pull-Up set.
)

const (
	PUPD15n = 15
	PUPD14n = 14
	PUPD13n = 13
	PUPD12n = 12
	PUPD11n = 11
	PUPD10n = 10
	PUPD9n  = 9
	PUPD8n  = 8
	PUPD7n  = 7
	PUPD6n  = 6
	PUPD5n  = 5
	PUPD4n  = 4
	PUPD3n  = 3
	PUPD2n  = 2
	PUPD1n  = 1
	PUPD0n  = 0
)

const (
	PDPD15 PDCRD = 0x01 << 15 //+ Port PD15 Pull-Down set.
	PDPD14 PDCRD = 0x01 << 14 //+ Port PD14 Pull-Down set.
	PDPD13 PDCRD = 0x01 << 13 //+ Port PD13 Pull-Down set.
	PDPD12 PDCRD = 0x01 << 12 //+ Port PD12 Pull-Down set.
	PDPD11 PDCRD = 0x01 << 11 //+ Port PD11 Pull-Down set.
	PDPD10 PDCRD = 0x01 << 10 //+ Port PD10 Pull-Down set.
	PDPD9  PDCRD = 0x01 << 9  //+ Port PD9 Pull-Down set.
	PDPD8  PDCRD = 0x01 << 8  //+ Port PD8 Pull-Down set.
	PDPD7  PDCRD = 0x01 << 7  //+ Port PD7 Pull-Down set.
	PDPD6  PDCRD = 0x01 << 6  //+ Port PD6 Pull-Down set.
	PDPD5  PDCRD = 0x01 << 5  //+ Port PD5 Pull-Down set.
	PDPD4  PDCRD = 0x01 << 4  //+ Port PD4 Pull-Down set.
	PDPD3  PDCRD = 0x01 << 3  //+ Port PD3 Pull-Down set.
	PDPD2  PDCRD = 0x01 << 2  //+ Port PD2 Pull-Down set.
	PDPD1  PDCRD = 0x01 << 1  //+ Port PD1 Pull-Down set.
	PDPD0  PDCRD = 0x01 << 0  //+ Port PD0 Pull-Down set.
)

const (
	PDPD15n = 15
	PDPD14n = 14
	PDPD13n = 13
	PDPD12n = 12
	PDPD11n = 11
	PDPD10n = 10
	PDPD9n  = 9
	PDPD8n  = 8
	PDPD7n  = 7
	PDPD6n  = 6
	PDPD5n  = 5
	PDPD4n  = 4
	PDPD3n  = 3
	PDPD2n  = 2
	PDPD1n  = 1
	PDPD0n  = 0
)

const (
	PUPE15 PUCRE = 0x01 << 15 //+ Port PE15 Pull-Up set.
	PUPE14 PUCRE = 0x01 << 14 //+ Port PE14 Pull-Up set.
	PUPE13 PUCRE = 0x01 << 13 //+ Port PE13 Pull-Up set.
	PUPE12 PUCRE = 0x01 << 12 //+ Port PE12 Pull-Up set.
	PUPE11 PUCRE = 0x01 << 11 //+ Port PE11 Pull-Up set.
	PUPE10 PUCRE = 0x01 << 10 //+ Port PE10 Pull-Up set.
	PUPE9  PUCRE = 0x01 << 9  //+ Port PE9 Pull-Up set.
	PUPE8  PUCRE = 0x01 << 8  //+ Port PE8 Pull-Up set.
	PUPE7  PUCRE = 0x01 << 7  //+ Port PE7 Pull-Up set.
	PUPE6  PUCRE = 0x01 << 6  //+ Port PE6 Pull-Up set.
	PUPE5  PUCRE = 0x01 << 5  //+ Port PE5 Pull-Up set.
	PUPE4  PUCRE = 0x01 << 4  //+ Port PE4 Pull-Up set.
	PUPE3  PUCRE = 0x01 << 3  //+ Port PE3 Pull-Up set.
	PUPE2  PUCRE = 0x01 << 2  //+ Port PE2 Pull-Up set.
	PUPE1  PUCRE = 0x01 << 1  //+ Port PE1 Pull-Up set.
	PUPE0  PUCRE = 0x01 << 0  //+ Port PE0 Pull-Up set.
)

const (
	PUPE15n = 15
	PUPE14n = 14
	PUPE13n = 13
	PUPE12n = 12
	PUPE11n = 11
	PUPE10n = 10
	PUPE9n  = 9
	PUPE8n  = 8
	PUPE7n  = 7
	PUPE6n  = 6
	PUPE5n  = 5
	PUPE4n  = 4
	PUPE3n  = 3
	PUPE2n  = 2
	PUPE1n  = 1
	PUPE0n  = 0
)

const (
	PDPE15 PDCRE = 0x01 << 15 //+ Port PE15 Pull-Down set.
	PDPE14 PDCRE = 0x01 << 14 //+ Port PE14 Pull-Down set.
	PDPE13 PDCRE = 0x01 << 13 //+ Port PE13 Pull-Down set.
	PDPE12 PDCRE = 0x01 << 12 //+ Port PE12 Pull-Down set.
	PDPE11 PDCRE = 0x01 << 11 //+ Port PE11 Pull-Down set.
	PDPE10 PDCRE = 0x01 << 10 //+ Port PE10 Pull-Down set.
	PDPE9  PDCRE = 0x01 << 9  //+ Port PE9 Pull-Down set.
	PDPE8  PDCRE = 0x01 << 8  //+ Port PE8 Pull-Down set.
	PDPE7  PDCRE = 0x01 << 7  //+ Port PE7 Pull-Down set.
	PDPE6  PDCRE = 0x01 << 6  //+ Port PE6 Pull-Down set.
	PDPE5  PDCRE = 0x01 << 5  //+ Port PE5 Pull-Down set.
	PDPE4  PDCRE = 0x01 << 4  //+ Port PE4 Pull-Down set.
	PDPE3  PDCRE = 0x01 << 3  //+ Port PE3 Pull-Down set.
	PDPE2  PDCRE = 0x01 << 2  //+ Port PE2 Pull-Down set.
	PDPE1  PDCRE = 0x01 << 1  //+ Port PE1 Pull-Down set.
	PDPE0  PDCRE = 0x01 << 0  //+ Port PE0 Pull-Down set.
)

const (
	PDPE15n = 15
	PDPE14n = 14
	PDPE13n = 13
	PDPE12n = 12
	PDPE11n = 11
	PDPE10n = 10
	PDPE9n  = 9
	PDPE8n  = 8
	PDPE7n  = 7
	PDPE6n  = 6
	PDPE5n  = 5
	PDPE4n  = 4
	PDPE3n  = 3
	PDPE2n  = 2
	PDPE1n  = 1
	PDPE0n  = 0
)

const (
	PUPF15 PUCRF = 0x01 << 15 //+ Port PF15 Pull-Up set.
	PUPF14 PUCRF = 0x01 << 14 //+ Port PF14 Pull-Up set.
	PUPF13 PUCRF = 0x01 << 13 //+ Port PF13 Pull-Up set.
	PUPF12 PUCRF = 0x01 << 12 //+ Port PF12 Pull-Up set.
	PUPF11 PUCRF = 0x01 << 11 //+ Port PF11 Pull-Up set.
	PUPF10 PUCRF = 0x01 << 10 //+ Port PF10 Pull-Up set.
	PUPF9  PUCRF = 0x01 << 9  //+ Port PF9 Pull-Up set.
	PUPF8  PUCRF = 0x01 << 8  //+ Port PF8 Pull-Up set.
	PUPF7  PUCRF = 0x01 << 7  //+ Port PF7 Pull-Up set.
	PUPF6  PUCRF = 0x01 << 6  //+ Port PF6 Pull-Up set.
	PUPF5  PUCRF = 0x01 << 5  //+ Port PF5 Pull-Up set.
	PUPF4  PUCRF = 0x01 << 4  //+ Port PF4 Pull-Up set.
	PUPF3  PUCRF = 0x01 << 3  //+ Port PF3 Pull-Up set.
	PUPF2  PUCRF = 0x01 << 2  //+ Port PF2 Pull-Up set.
	PUPF1  PUCRF = 0x01 << 1  //+ Port PF1 Pull-Up set.
	PUPF0  PUCRF = 0x01 << 0  //+ Port PF0 Pull-Up set.
)

const (
	PUPF15n = 15
	PUPF14n = 14
	PUPF13n = 13
	PUPF12n = 12
	PUPF11n = 11
	PUPF10n = 10
	PUPF9n  = 9
	PUPF8n  = 8
	PUPF7n  = 7
	PUPF6n  = 6
	PUPF5n  = 5
	PUPF4n  = 4
	PUPF3n  = 3
	PUPF2n  = 2
	PUPF1n  = 1
	PUPF0n  = 0
)

const (
	PDPF15 PDCRF = 0x01 << 15 //+ Port PF15 Pull-Down set.
	PDPF14 PDCRF = 0x01 << 14 //+ Port PF14 Pull-Down set.
	PDPF13 PDCRF = 0x01 << 13 //+ Port PF13 Pull-Down set.
	PDPF12 PDCRF = 0x01 << 12 //+ Port PF12 Pull-Down set.
	PDPF11 PDCRF = 0x01 << 11 //+ Port PF11 Pull-Down set.
	PDPF10 PDCRF = 0x01 << 10 //+ Port PF10 Pull-Down set.
	PDPF9  PDCRF = 0x01 << 9  //+ Port PF9 Pull-Down set.
	PDPF8  PDCRF = 0x01 << 8  //+ Port PF8 Pull-Down set.
	PDPF7  PDCRF = 0x01 << 7  //+ Port PF7 Pull-Down set.
	PDPF6  PDCRF = 0x01 << 6  //+ Port PF6 Pull-Down set.
	PDPF5  PDCRF = 0x01 << 5  //+ Port PF5 Pull-Down set.
	PDPF4  PDCRF = 0x01 << 4  //+ Port PF4 Pull-Down set.
	PDPF3  PDCRF = 0x01 << 3  //+ Port PF3 Pull-Down set.
	PDPF2  PDCRF = 0x01 << 2  //+ Port PF2 Pull-Down set.
	PDPF1  PDCRF = 0x01 << 1  //+ Port PF1 Pull-Down set.
	PDPF0  PDCRF = 0x01 << 0  //+ Port PF0 Pull-Down set.
)

const (
	PDPF15n = 15
	PDPF14n = 14
	PDPF13n = 13
	PDPF12n = 12
	PDPF11n = 11
	PDPF10n = 10
	PDPF9n  = 9
	PDPF8n  = 8
	PDPF7n  = 7
	PDPF6n  = 6
	PDPF5n  = 5
	PDPF4n  = 4
	PDPF3n  = 3
	PDPF2n  = 2
	PDPF1n  = 1
	PDPF0n  = 0
)

const (
	PUPG15 PUCRG = 0x01 << 15 //+ Port PG15 Pull-Up set.
	PUPG14 PUCRG = 0x01 << 14 //+ Port PG14 Pull-Up set.
	PUPG13 PUCRG = 0x01 << 13 //+ Port PG13 Pull-Up set.
	PUPG12 PUCRG = 0x01 << 12 //+ Port PG12 Pull-Up set.
	PUPG11 PUCRG = 0x01 << 11 //+ Port PG11 Pull-Up set.
	PUPG10 PUCRG = 0x01 << 10 //+ Port PG10 Pull-Up set.
	PUPG9  PUCRG = 0x01 << 9  //+ Port PG9 Pull-Up set.
	PUPG8  PUCRG = 0x01 << 8  //+ Port PG8 Pull-Up set.
	PUPG7  PUCRG = 0x01 << 7  //+ Port PG7 Pull-Up set.
	PUPG6  PUCRG = 0x01 << 6  //+ Port PG6 Pull-Up set.
	PUPG5  PUCRG = 0x01 << 5  //+ Port PG5 Pull-Up set.
	PUPG4  PUCRG = 0x01 << 4  //+ Port PG4 Pull-Up set.
	PUPG3  PUCRG = 0x01 << 3  //+ Port PG3 Pull-Up set.
	PUPG2  PUCRG = 0x01 << 2  //+ Port PG2 Pull-Up set.
	PUPG1  PUCRG = 0x01 << 1  //+ Port PG1 Pull-Up set.
	PUPG0  PUCRG = 0x01 << 0  //+ Port PG0 Pull-Up set.
)

const (
	PUPG15n = 15
	PUPG14n = 14
	PUPG13n = 13
	PUPG12n = 12
	PUPG11n = 11
	PUPG10n = 10
	PUPG9n  = 9
	PUPG8n  = 8
	PUPG7n  = 7
	PUPG6n  = 6
	PUPG5n  = 5
	PUPG4n  = 4
	PUPG3n  = 3
	PUPG2n  = 2
	PUPG1n  = 1
	PUPG0n  = 0
)

const (
	PDPG15 PDCRG = 0x01 << 15 //+ Port PG15 Pull-Down set.
	PDPG14 PDCRG = 0x01 << 14 //+ Port PG14 Pull-Down set.
	PDPG13 PDCRG = 0x01 << 13 //+ Port PG13 Pull-Down set.
	PDPG12 PDCRG = 0x01 << 12 //+ Port PG12 Pull-Down set.
	PDPG11 PDCRG = 0x01 << 11 //+ Port PG11 Pull-Down set.
	PDPG10 PDCRG = 0x01 << 10 //+ Port PG10 Pull-Down set.
	PDPG9  PDCRG = 0x01 << 9  //+ Port PG9 Pull-Down set.
	PDPG8  PDCRG = 0x01 << 8  //+ Port PG8 Pull-Down set.
	PDPG7  PDCRG = 0x01 << 7  //+ Port PG7 Pull-Down set.
	PDPG6  PDCRG = 0x01 << 6  //+ Port PG6 Pull-Down set.
	PDPG5  PDCRG = 0x01 << 5  //+ Port PG5 Pull-Down set.
	PDPG4  PDCRG = 0x01 << 4  //+ Port PG4 Pull-Down set.
	PDPG3  PDCRG = 0x01 << 3  //+ Port PG3 Pull-Down set.
	PDPG2  PDCRG = 0x01 << 2  //+ Port PG2 Pull-Down set.
	PDPG1  PDCRG = 0x01 << 1  //+ Port PG1 Pull-Down set.
	PDPG0  PDCRG = 0x01 << 0  //+ Port PG0 Pull-Down set.
)

const (
	PDPG15n = 15
	PDPG14n = 14
	PDPG13n = 13
	PDPG12n = 12
	PDPG11n = 11
	PDPG10n = 10
	PDPG9n  = 9
	PDPG8n  = 8
	PDPG7n  = 7
	PDPG6n  = 6
	PDPG5n  = 5
	PDPG4n  = 4
	PDPG3n  = 3
	PDPG2n  = 2
	PDPG1n  = 1
	PDPG0n  = 0
)

const (
	PUPH1 PUCRH = 0x01 << 1 //+ Port PH1 Pull-Up set.
	PUPH0 PUCRH = 0x01 << 0 //+ Port PH0 Pull-Up set.
)

const (
	PUPH1n = 1
	PUPH0n = 0
)

const (
	PDPH1 PDCRH = 0x01 << 1 //+ Port PH1 Pull-Down set.
	PDPH0 PDCRH = 0x01 << 0 //+ Port PH0 Pull-Down set.
)

const (
	PDPH1n = 1
	PDPH0n = 0
)
