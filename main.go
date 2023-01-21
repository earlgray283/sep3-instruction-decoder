package main

import (
	"flag"
	"fmt"
	"log"
	"strconv"
)

const version string = "1.0.0"

var versionOpt bool

func main() {
	flag.Parse()

	flag.BoolVar(&versionOpt, "version", false, "show version")
	if versionOpt {
		fmt.Println(version)
		return
	}

	for {
		fmt.Print(">>> ")
		var s string
		fmt.Scan(&s)

		num, err := strconv.ParseUint(s, 16, 16)
		if err != nil {
			log.Println(err)
			continue
		}
		bin := fmt.Sprintf("%016b", num)

		rawInstruction, rawFromOp, rawToOp := bin[:6], bin[6:11], bin[11:]
		instruction, fromOp, toOp := instructionToString(rawInstruction), operandToString(rawFromOp), operandToString(rawToOp)
		fmt.Printf("%v %v, %v [%v %v %v]\n", instruction, fromOp, toOp, rawInstruction, rawFromOp, rawToOp)
	}
}

type addressMode int

const (
	addressModeD  addressMode = iota // 直接(00)
	addressModeI                     // 間接(01)
	addressModeMI                    // プレデクリメント間接(10)
	addressModeIP                    // ポストインクリメント間接(11)
)

func (a addressMode) String() string {
	textList := []string{"D", "I", "MI", "IP"}
	return textList[a]
}

func operandToString(bin string) string {
	const (
		addressModeBitNum = 2
		registerBitNum    = 3
	)
	addressModeList := []addressMode{addressModeD, addressModeI, addressModeMI, addressModeIP}
	rawAddressMode, _ := strconv.ParseUint(bin[:2], 2, addressModeBitNum)
	am := addressModeList[rawAddressMode]
	registerNum, _ := strconv.ParseUint(bin[2:5], 2, registerBitNum)
	switch am {
	case addressModeD:
		return fmt.Sprintf("R%v", registerNum)
	case addressModeI:
		return fmt.Sprintf("(R%v)", registerNum)
	case addressModeMI:
		return fmt.Sprintf("-(R%v)", registerNum)
	case addressModeIP:
		return fmt.Sprintf("(R%v)+", registerNum)
	}
	panic("unreachable")
}

func instructionToString(bin string) string {
	return instructionByHex[bin]
}

var instructionByHex map[string]string = map[string]string{
	"000000": "HLT",
	"000001": "HLT",
	"000010": "HLT",
	"000011": "HLT",
	"000100": "CLR",
	"001000": "ASL",
	"001001": "ASR",
	"001100": "LSL",
	"001101": "LSR",
	"001110": "ROL",
	"001111": "ROR",
	"010000": "MOV",
	"010001": "JMP",
	"010010": "RET",
	"010011": "RIT",
	"010100": "ADD",
	"010101": "RJP",
	"011000": "SUB",
	"011011": "CMP",
	"011100": "NOP",
	"011101": "NOP",
	"011110": "NOP",
	"011111": "NOP",
	"100000": "OR",
	"100001": "XOR",
	"100010": "AND",
	"100011": "BIT",
	"101100": "JSR",
	"101101": "RJS",
	"101110": "SVC",
	"110000": "BRN",
	"110001": "BRZ",
	"110010": "BRV",
	"110011": "BRC",
	"111000": "BRN",
	"111001": "BRZ",
	"111010": "BRV",
	"111011": "BRC",
}
