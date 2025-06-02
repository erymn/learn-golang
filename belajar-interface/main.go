package main

import "math"

type hitung2d interface {
	luas() float64
	keliling() float64
}

type hitung3d interface {
	volume() float64
}

// Persegi
type persegi struct {
	sisi float64
}

func (p persegi) luas() float64 {
	return p.sisi * p.sisi
}

func (p persegi) keliling() float64 {
	return p.sisi * 4
}

// lingkaran
type lingkaran struct {
	diameter float64
}

func (l lingkaran) jariJari() float64 {
	return l.diameter / 2
}
func (l lingkaran) luas() float64 {
	return 3.14 * l.jariJari() * l.jariJari()
}
func (l lingkaran) keliling() float64 {
	return 3.14 * l.diameter
}

// embedded interface
type hitung interface {
	hitung2d
	hitung3d
}

type kubus struct {
	sisi float64
}

func (k *kubus) volume() float64 {
	return math.Pow(k.sisi, 3)
}
func (k *kubus) luas() float64 {
	return math.Pow(k.sisi, 2) * 6
}
func (k *kubus) keliling() float64 {
	return k.sisi * 12
}

func main() {
	// var bangunDatar hitung

	// bangunDatar = persegi{10.0}
	// println("Luas Persegi:", bangunDatar.luas())
	// println("Keliling Persegi:", bangunDatar.keliling())

	// bangunDatar = lingkaran{14.5}
	// //cara casting object interface yang agak unik, karena jari-jari error
	// var bangunLingkaran lingkaran = bangunDatar.(lingkaran)

	// println("Luas Lingkaran:", bangunDatar.luas())
	// println("Keliling Lingkaran:", bangunDatar.keliling())
	// println("Jari-jari Lingkaran:", bangunLingkaran.jariJari())

	println("Hitung bangun ruang")
	var bangunRuang hitung = &kubus{5}
	println("Volume Kubus:", bangunRuang.volume())
	println("Luas Kubus:", bangunRuang.luas())
	println("Keliling Kubus:", bangunRuang.keliling())
}
