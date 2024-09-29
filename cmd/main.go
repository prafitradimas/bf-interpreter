package main

import "fmt"

const (
	INC_PTR    = '>' // increment pointer
	DEC_PTR    = '<' // decrement pointer
	INC_VAL    = '+' // increment value at pointer
	DEC_VAL    = '-' // decrement value at pointer
	INPUT      = ',' // read one character from input into value at pointer
	OUTPUT     = '.' // print value at pointer to output as character
	BEGIN_LOOP = '[' // begin loop (while value at pointer is non-zero)
	END_LOOP   = ']' // end loop (while value at pointer is non-zero)
)

func main() {
	input := "++++++++++[>+>+++>+++++++>++++++++++<<<<-]>>>++.>+.+++++++..+++.<<++++++++++++++.------------.>+++++++++++++++.>.+++.------.--------.<<+."
	mem := make([]rune, 1024)

	// var parse func(s string, pos int)
	parse := func(s string, pos int) {
		for i := 0; len(s) > i; i++ {
			switch s[i] {
			case INC_PTR:
				if len(mem) > pos {
					pos += 1
				}
			case DEC_PTR:
				if pos > 0 {
					pos -= 1
				}
			case INC_VAL:
				mem[pos]++
			case DEC_VAL:
				mem[pos]--
			case INPUT:
			case OUTPUT:
				fmt.Printf("%s", string(mem[pos]))
			case BEGIN_LOOP:
				if mem[pos] == 0 {
					i++
					count := 0 // count '[' in each encounter
					for count > 0 || s[i] != END_LOOP {
						switch s[i] {
						case BEGIN_LOOP:
							count++
						case END_LOOP:
							count--
						}
						i++
					}
				}
			case END_LOOP:
				if mem[pos] != 0 {
					i--
					count := 0 // count ']' in each encounter
					for count > 0 || s[i] != BEGIN_LOOP {
						switch s[i] {
						case BEGIN_LOOP:
							count--
						case END_LOOP:
							count++
						}
						i--
					}
				}
			}
		}
	}

	parse(input, 0)
}
