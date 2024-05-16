package main

import (
	"fmt"
)

type Ban interface {
	Jenis() string
}

type BanKaret struct{}

func (b BanKaret) Jenis() string {
	return "Ban Karet"
}

type BanKayu struct{}

func (b BanKayu) Jenis() string {
	return "Ban Kayu"
}

type BanBesi struct{}

func (b BanBesi) Jenis() string {
	return "Ban Besi"
}

type Pintu struct {
	Sisi     string
	KetukSuara string
	BukaSuara string
}

func (p Pintu) Ketuk() {
	fmt.Println(p.KetukSuara)
}

func (p Pintu) Buka() {
	fmt.Println(p.BukaSuara)
}

type Mobil struct {
	Roda [4]Ban
	PintuKanan Pintu
	PintuKiri  Pintu
}

func NewMobil() *Mobil {
	return &Mobil{
		PintuKanan: Pintu{Sisi: "Kanan", KetukSuara: "Knock", BukaSuara: "beep"},
		PintuKiri:  Pintu{Sisi: "Kiri", KetukSuara: "beep", BukaSuara: "Knock"},
	}
}

func (m *Mobil) GantiBan(index int, ban Ban) {
	if index < 0 || index >= len(m.Roda) {
		fmt.Println("Index roda tidak valid")
		return
	}
	m.Roda[index] = ban
}

func main() {
	mobil := NewMobil()
	mobil.GantiBan(0, BanKaret{})
	mobil.GantiBan(1, BanKayu{})
	mobil.GantiBan(2, BanBesi{})
	mobil.GantiBan(3, BanKaret{})

	fmt.Print("Ketuk pintu kanan: ")
	mobil.PintuKanan.Ketuk()

	fmt.Print("Buka pintu kanan: ")
	mobil.PintuKanan.Buka()

	fmt.Print("Ketuk pintu kiri: ")
	mobil.PintuKiri.Ketuk()

	fmt.Print("Buka pintu kiri: ")
	mobil.PintuKiri.Buka()
}
