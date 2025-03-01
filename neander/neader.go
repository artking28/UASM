package neander

import "fmt"

type Result struct {
	Ac, Pc byte
}

const (
	NOP = 0   // Nenhuma operação
	STA = 16  // Armazena Acumulador no endereço “end” da memória
	LDA = 32  // Carrega o Acumulador com o conteúdo do endereço “end” da memória
	ADD = 48  // Soma o conteúdo do endereço “end” da memória ao Acumulador
	OR  = 64  // Efetua operação lógica “OU” do conteúdo do endereço “end” da memória ao Acumulador
	AND = 80  // Efetua operação lógica “E” do conteúdo do endereço “end” da memória ao Acumulador
	NOT = 96  // Inverte todos os bits do Acumulador
	JMP = 128 // Desvio incondicional para o endereço “end” da memória
	JN  = 144 // Desvio condicional, se “Ac!=0”, para o endereço “end” da memória
	JZ  = 160 // Desvio condicional, se "Ac==0", para o endereço “end” da memória
	HLT = 240 // Encerra o ciclo de busca-decodificação-execução
)

func RunProgram(program []byte) (Result, []byte) {
	padding := 4
	var result Result
	for i := padding; i < len(program); i += padding {
		mnemonic := program[i]
		addr := program[i+2]
		addrValueIndex := int(addr)*2 + padding
		switch mnemonic {
		case NOP:
			result.Pc += 1
			continue
		case STA:
			result.Pc += 2
			program[addrValueIndex] = result.Ac
			break
		case LDA:
			result.Pc += 2
			result.Ac = program[addrValueIndex]
			break
		case ADD:
			result.Pc += 2
			result.Ac += program[addrValueIndex]
			break
		case OR:
			result.Pc += 2
			result.Ac |= program[addrValueIndex]
			break
		case AND:
			result.Pc += 2
			result.Ac &= program[addrValueIndex]
			break
		case NOT:
			result.Pc += 1
			result.Ac ^= result.Ac
			break
		case JMP:
			result.Pc += 2
			i = int(program[addrValueIndex])
			break
		case JN:
			result.Pc += 2
			if result.Ac != 0 {
				i = int(program[addrValueIndex])
			}
			continue
		case JZ:
			result.Pc += 2
			if result.Ac == 0 {
				i = int(program[addrValueIndex])
			}
			continue
		case HLT:
			result.Pc += 1
			i = len(program)
			break
		}
	}
	return result, program
}

func PrintProgram(program []byte) {
	padding := 4
	print("\nProgram:\n")
	for i := padding; i < len(program); i += padding {
		mnemonic := program[i]
		addr := int(program[i+2])
		addrV := program[addr*2+padding]
		fmt.Printf("[%.3d]", i/4)
		switch mnemonic {
		case NOP:
			fmt.Printf("\tNOP\n")
			break
		case STA:
			fmt.Printf("\tSTA %d(value=%d)\n", addr, addrV)
			break
		case LDA:
			fmt.Printf("\tLDA %d(value=%d)\n", addr, addrV)
			break
		case ADD:
			fmt.Printf("\tADD %d(value=%d)\n", addr, addrV)
			break
		case OR:
			fmt.Printf("\tOR %d(value=%d)\n", addr, addrV)
			break
		case AND:
			fmt.Printf("\tAND %d(value=%d)\n", addr, addrV)
			break
		case NOT:
			fmt.Printf("\tNOT\n")
			break
		case JMP:
			fmt.Printf("\tJMP %d(line=%d)\n", addr, addrV)
			break
		case JN:
			fmt.Printf("\tJN  %d(line=%d)\n", addr, addrV)
			break
		case JZ:
			fmt.Printf("\tJZ  %d(line=%d)\n", addr, addrV)
			break
		case HLT:
			fmt.Printf("\tHLT\n")
			i = len(program)
			break
		}
	}
}
