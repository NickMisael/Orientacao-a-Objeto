package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"os/exec"
	"runtime"
	"strconv"
	"strings"
	"time"
)

var so = runtime.GOOS

// LIMPATELA 
func limpa() {
	if so == "windows" {
		clear := exec.Command("cmd", "/c", "cls")
		clear.Stdout = os.Stdout
		clear.Run()
	} else if so == "linux" {
		clear := exec.Command("clear")
		clear.Stdout = os.Stdout
		clear.Run()
	}
}

// TELA INICIAL 
func Menu() {
	fmt.Println("\t0- Sair ")
	fmt.Println("\t1- Depositar   ")
	fmt.Println("\t2- Sacar  ")
	fmt.Println("\t3- Extrato")
	fmt.Printf("\t>> ")
}

// CLASSE CONTA
type Conta struct {
	Num     int
	Titular string
	Saldo   float64
	Limite  float64
}

// FUNCAO PARA CRIAR CONTA
func (c *Conta) cria_Conta() {
	scanner := bufio.NewScanner(os.Stdin)

	// DANDO UM ID
	rand.Seed(time.Now().UnixNano())
	c.Num = rand.Intn(1000)

	// VALIDANDO O NOME
	for {
		fmt.Printf("\tDigite seu nome: ")
		for scanner.Scan() {
			c.Titular = scanner.Text()
			break
		}
		if scanner.Err() != nil {
			panic(scanner.Err())
		}
		if len(c.Titular) < 3 {
			fmt.Println("\tError: Número de Caracteres Inválido!!")
			time.Sleep(time.Second + 3)
			limpa()
			continue
		}
		bin := []byte(c.Titular)
		count := 0
		for _, bit := range bin {
			if bit > 0x40 && bit < 0x5B || bit > 0x60 && bit < 0x7B || bit == 0x20 {
				count -= -1
			}
		}

		if count != len(bin) {
			fmt.Println("\tError: Caracter Inválido!!")
			time.Sleep(time.Second + 3)
			limpa()
			continue
		}
		break
	}

	// VALIDANDO O SALDO
	for {
		var saldo string
		fmt.Printf("\tDigite seu saldo inicial: ")
		for scanner.Scan() {
			saldo = scanner.Text()
			break
		}
		if scanner.Err() != nil {
			panic(scanner.Err())
		}
		if len(c.Titular) < 1 {
			fmt.Println("\tError: Número de Caracteres Inválido!!")
			time.Sleep(time.Second + 3)
			limpa()
			continue
		}
		bin := []byte(saldo)
		count := 0
		for _, bit := range bin {
			if bit >= 0x30 && bit <= 0x39 || bit == 0x2E {
				count -= -1
			}
		}

		if count != len(bin) {
			fmt.Println("\tError: Caracter Inválido!!")
			time.Sleep(time.Second + 3)
			limpa()
			continue
		}
		c.Saldo, _ = strconv.ParseFloat(saldo, 64)
		if c.Saldo < 20 {
			fmt.Println("\tError: Valor Inválido!")
			time.Sleep(time.Second + 3)
			limpa()
			continue
		}
		break
	}

	// Definindo Limite
	if c.Saldo > 3000 {
		c.Limite = 3000
	} else {
		c.Limite = c.Saldo * 0.07
	}
}

// FUNCAO PARA CONSULTAR O EXTRATO
func (c Conta) Extrato() {
	fmt.Printf("\tExtrato:\n")
	fmt.Printf("\tID: %d\n", c.Num)
	saldo := fmt.Sprintf("%.2f", c.Saldo)
	saldo = strings.Replace(saldo, ".", ",", 1)
	limite := fmt.Sprintf("%.2f", c.Limite)
	limite = strings.Replace(limite, ".", ",", 1)
	fmt.Printf("\tNome: %s\n", c.Titular)
	fmt.Printf("\tSaldo: R$ %s\n", saldo)
	fmt.Printf("\tLimite: R$ %s\n", limite)
}

func main() {
	limpa()
	// Declarando variáveis
	scanner := bufio.NewScanner(os.Stdin)
	c := Conta{}

	// CRIANDO UMA CONTA
	c.cria_Conta()
	limpa()
	fmt.Printf("\tEstamos Carregando o Sistema...")
	fmt.Printf("\a\a\a")
	time.Sleep(time.Second + 3)
	limpa()

	for {
		var esc string
		Menu()
		for scanner.Scan() {
			esc = scanner.Text()
			break
		}
		if scanner.Err() != nil {
			panic(scanner.Err())
		}

		if esc == "0" {
			fmt.Println("\tObrigado por utilizar nosso programa!!")
			fmt.Println("\tVolte Sempre!")
			break
		}

		if len(esc) < 1 || len(esc) > 1 {
			fmt.Println("\tError: Engraçadão Você hein!!")
			time.Sleep(time.Second + 3)
			limpa()
			continue
		}

		bin := []byte(esc)
		ver := true
		for _, bit := range bin {
			if bit < 0x30 || bit > 0x33 {
				ver = false
			}
		}
		if ver == false {
			fmt.Println("\tError: Engraçadão Você hein!!")
			time.Sleep(time.Second + 3)
			limpa()
			continue
		}
		limpa()
		switch esc {
		case "1":
			fmt.Println("\tDepositar()")
			fmt.Printf("\tTecle para continuar...")
			fmt.Scanln(&esc)
			time.Sleep(time.Second + 2)
			limpa()
		case "2":
			fmt.Println("\tSacar()")
			fmt.Printf("\tTecle para continuar...")
			fmt.Scanln(&esc)
			time.Sleep(time.Second + 2)
			limpa()
		case "3":
			c.Extrato()
			fmt.Printf("\tTecle para continuar...")
			fmt.Scanln(&esc)
			time.Sleep(time.Second + 2)
			limpa()
		}
	}
}
