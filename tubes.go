package main

import (
	"fmt"
	"strings"
	"time"
)

type BahanMakanan struct {
	nama        string
	stok        int
	kadaluwarsa string
}

func tampilkanReminder(bahanMakanans []BahanMakanan) {
	// Hitung jumlah bahan yang perlu diingatkan
	var count int
	var bahanKadaluwarsa []BahanMakanan

	fmt.Println("\n🔔 Peringatan Kadaluwarsa Bahan Makanan:")
	fmt.Println("===============================================")

	for _, bahan := range bahanMakanans {
		kadaluwarsa, err := time.Parse("2006-01-02", bahan.kadaluwarsa)
		if err != nil {
			continue // Lewati jika format tanggal salah
		}

		selisihHari := int(kadaluwarsa.Sub(time.Now()).Hours() / 24)

		if selisihHari < 0 {
			// Bahan sudah kadaluwarsa
			bahanKadaluwarsa = append(bahanKadaluwarsa, bahan)
			fmt.Printf("❌ %-15s | SUDAH KADALUWARSA | %s\n",
				bahan.nama, bahan.kadaluwarsa)
			count++
		} else if selisihHari <= 7 {
			// Bahan akan kadaluwarsa dalam 7 hari
			bahanKadaluwarsa = append(bahanKadaluwarsa, bahan)
			fmt.Printf("⚠️  %-15s | %2d hari lagi     | %s\n",
				bahan.nama, selisihHari, bahan.kadaluwarsa)
			count++
		}
	}

	if count == 0 {
		fmt.Println("✅ Tidak ada bahan yang akan kadaluwarsa dalam 7 hari ke depan")
	} else {
		fmt.Printf("\nTotal: %d bahan perlu perhatian\n", count)
	}
	fmt.Println("===============================================")
}

func insertionSort(data []BahanMakanan) []BahanMakanan {
	result := make([]BahanMakanan, len(data))
	copy(result, data)

	for i := 1; i < len(result); i++ {
		key := result[i]
		j := i - 1
		for j >= 0 && result[j].stok > key.stok {
			result[j+1] = result[j]
			j--
		}
		result[j+1] = key
	}
	return result
}

func selectionSort(data []BahanMakanan) []BahanMakanan {
	result := make([]BahanMakanan, len(data))
	copy(result, data)

	for i := 0; i < len(result)-1; i++ {
		minIdx := i
		for j := i + 1; j < len(result); j++ {
			if result[j].kadaluwarsa < result[minIdx].kadaluwarsa {
				minIdx = j
			}
		}
		result[i], result[minIdx] = result[minIdx], result[i]
	}
	return result
}

func tampilkanData(data []BahanMakanan) {
	fmt.Println("\nDaftar Bahan Makanan:")
	fmt.Println("--------------------------------------------------")
	fmt.Printf("%-20s %-10s %-15s\n", "Nama", "Stok", "Kadaluwarsa")
	fmt.Println("--------------------------------------------------")
	for _, bahan := range data {
		fmt.Printf("%-20s %-10d %-15s\n", bahan.nama, bahan.stok, bahan.kadaluwarsa)
	}
	fmt.Println("--------------------------------------------------")
}

func daftarBahanMakanan(data []BahanMakanan) {
	var pilihDaftar int

	fmt.Println("\nMau lihat daftar bahan makanan berdasarkan apa?")
	fmt.Println("1. Nama Bahan Makanan")
	fmt.Println("2. Jumlah Stok")
	fmt.Println("3. Tanggal Kadaluwarsa")
	fmt.Print("Pilih: ")
	fmt.Scanln(&pilihDaftar)

	var hasil []BahanMakanan

	switch pilihDaftar {
	case 1:
		hasil = sortByNama(data)
	case 2:
		hasil = insertionSort(data)
	case 3:
		hasil = selectionSort(data)
	default:
		fmt.Println("Pilihan tidak valid, menampilkan data asli")
		hasil = data
	}

	tampilkanData(hasil)

}

func tambahData(data *[]BahanMakanan) {
	var (
		nama        string
		stok        int
		kadaluwarsa string
	)

	// Input nama bahan
	fmt.Print("Masukkan nama bahan makanan: ")
	fmt.Scanln(&nama)

	// Input stok dengan validasi
	for {
		fmt.Print("Masukkan jumlah stok: ")
		_, err := fmt.Scanln(&stok)
		if err == nil {
			break
		}
		fmt.Println("Error: Stok harus berupa angka!")
		fmt.Scanln() // Membersihkan buffer input
	}

	// Input tanggal kadaluwarsa
	fmt.Print("Masukkan tanggal kadaluwarsa (format: yyyy-mm-dd): ")
	fmt.Scanln(&kadaluwarsa)

	// Tambahkan data baru
	*data = append(*data, BahanMakanan{
		nama:        nama,
		stok:        stok,
		kadaluwarsa: kadaluwarsa,
	})

	fmt.Println("\n✅ Bahan makanan berhasil ditambahkan!")
	fmt.Printf("Nama: %s\nStok: %d\nKadaluwarsa: %s\n\n", nama, stok, kadaluwarsa)
}

func ubahData(data *[]BahanMakanan) {
	// Input nomor yang akan diubah
	var nomor int
	fmt.Print("\n» Masukkan nomor bahan yang akan diubah: ")
	_, err := fmt.Scanln(&nomor)
	if err != nil || nomor < 1 || nomor > len(*data) {
		fmt.Println("❌ Nomor tidak valid")
		return
	}

	// Tampilkan data yang akan diubah
	bahan := &(*data)[nomor-1]
	fmt.Printf("\nData yang akan diubah:\nNama: %s\nStok: %d\nKadaluwarsa: %s\n",
		bahan.nama, bahan.stok, bahan.kadaluwarsa)

	// Menu pilihan yang akan diubah
	var pilihan int
	fmt.Println("\nPilih yang akan diubah:")
	fmt.Println("1. Nama")
	fmt.Println("2. Stok")
	fmt.Println("3. Tanggal Kadaluwarsa")
	fmt.Println("4. Ubah Ketiganya")
	fmt.Print("Pilihan: ")
	fmt.Scanln(&pilihan)

	// Proses perubahan
	switch pilihan {
	case 1:
		fmt.Print("Masukkan nama baru: ")
		fmt.Scanln(&bahan.nama)
	case 2:
		for {
			fmt.Print("Masukkan stok baru: ")
			_, err := fmt.Scanln(&bahan.stok)
			if err == nil {
				break
			}
			fmt.Println("Error: Stok harus angka!")
			fmt.Scanln() // Clear buffer
		}
	case 3:
		fmt.Print("Masukkan tanggal baru (yyyy-mm-dd): ")
		fmt.Scanln(&bahan.kadaluwarsa)
	case 4:
		fmt.Print("Masukkan nama baru: ")
		fmt.Scanln(&bahan.nama)

		for {
			fmt.Print("Masukkan stok baru: ")
			_, err := fmt.Scanln(&bahan.stok)
			if err == nil {
				break
			}
			fmt.Println("Error: Stok harus angka!")
			fmt.Scanln()
		}

		fmt.Print("Masukkan tanggal baru (yyyy-mm-dd): ")
		fmt.Scanln(&bahan.kadaluwarsa)
	default:
		fmt.Println("Pilihan tidak valid")
		return
	}

	fmt.Printf("\n✅ Data berhasil diubah:\nNama: %s\nStok: %d\nKadaluwarsa: %s\n\n",
		bahan.nama, bahan.stok, bahan.kadaluwarsa)
}

func hapusData(data *[]BahanMakanan) {
	// Cek jika data kosong
	if len(*data) == 0 {
		fmt.Println("\n❌ Tidak ada data bahan makanan yang bisa dihapus")
		return
	}

	// Input nomor yang akan dihapus
	var nomor int
	fmt.Print("Masukkan nomor bahan yang akan dihapus: ")
	_, err := fmt.Scanln(&nomor)
	if err != nil {
		fmt.Println("\n❌ Input tidak valid")
		return
	}

	// Validasi nomor
	if nomor < 1 || nomor > len(*data) {
		fmt.Println("\n❌ Nomor tidak valid")
		return
	}

	// Proses penghapusan dengan pergeseran elemen
	bahanTerhapus := (*data)[nomor-1].nama
	for i := nomor - 1; i < len(*data)-1; i++ {
		(*data)[i] = (*data)[i+1]
	}
	*data = (*data)[:len(*data)-1]

	fmt.Printf("\n✅ Bahan makanan '%s' berhasil dihapus\n", bahanTerhapus)
}

func laporanStok(data []BahanMakanan) {
	total := len(data)
	digunakan := 0
	tersedia := 0

	for _, bahan := range data {
		if bahan.stok == 0 {
			digunakan++
		} else {
			tersedia++
		}
	}

	fmt.Println("\n📦 Laporan Stok Bahan Makanan")
	fmt.Println("===============================================")
	fmt.Printf("Total Bahan Makanan        : %d\n", total)
	fmt.Printf("Bahan Tersedia (stok > 0)  : %d\n", tersedia)
	fmt.Printf("Bahan Telah Digunakan      : %d\n", digunakan)
	fmt.Println("===============================================")
}

func sortByNama(data []BahanMakanan) []BahanMakanan {
	hasil := make([]BahanMakanan, len(data))
	copy(hasil, data)

	for i := 0; i < len(hasil)-1; i++ {
		min := i
		for j := i + 1; j < len(hasil); j++ {
			if strings.ToLower(hasil[j].nama) < strings.ToLower(hasil[min].nama) {
				min = j
			}
		}
		hasil[i], hasil[min] = hasil[min], hasil[i]
	}
	return hasil
}

func pencarianCepat(data []BahanMakanan, nama string) int {
	left := 0
	right := len(data) - 1
	found := -1

	for left <= right && found == -1 {
		mid := (left + right) / 2
		if strings.ToLower(nama) < strings.ToLower(data[mid].nama) {
			right = mid - 1
		} else if strings.ToLower(nama) > strings.ToLower(data[mid].nama) {
			left = mid + 1
		} else {
			found = mid
		}
	}
	return found
}

func sequentialSearch(data []BahanMakanan, nama string) int {
	for i, bahan := range data {
		if strings.ToLower(bahan.nama) == strings.ToLower(nama) {
			return i
		}
	}
	return -1
}

func binerSearch(data []BahanMakanan) {
	var pilihCari int
	fmt.Println("Pilih opsi pencarian")
	fmt.Println("1. Pencarian Cepat (Binary Search)")
	fmt.Println("2. Pencarian Biasa (Sequential Search)")
	fmt.Println("3. Kembali ke Menu Utama")
	fmt.Print("Pilih : ")
	fmt.Scan(&pilihCari)

	switch pilihCari {
	case 1:
		// Urutkan data terlebih dahulu
		sorted := sortByNama(data)
		fmt.Println("\nData Bahan Makanan (Terurut berdasarkan Nama):")
		fmt.Println("===============================================")
		fmt.Printf("%-3s | %-15s | %-5s | %-15s\n", "No.", "Nama Bahan", "Stok", "Kadaluwarsa")
		fmt.Println("===============================================")
		for i, bahan := range sorted {
			fmt.Printf("%-3d | %-15s | %-5d | %-15s\n", i+1, bahan.nama, bahan.stok, bahan.kadaluwarsa)
		}
		fmt.Println("===============================================")

		var cariNama string
		fmt.Print("\nMasukkan nama bahan makanan yang dicari: ")
		fmt.Scan(&cariNama)

		index := pencarianCepat(sorted, cariNama)

		if index != -1 {
			fmt.Println("\nHasil Pencarian")
			fmt.Println("---------------------------")
			fmt.Printf("Nama Bahan Makanan : %s\n", sorted[index].nama)
			fmt.Printf("Stok               : %d\n", sorted[index].stok)
			fmt.Printf("Kadaluwarsa        : %s\n\n", sorted[index].kadaluwarsa)
		} else {
			fmt.Println("Data tidak ditemukan.")
		}

	case 2:
		var cariNama string
		fmt.Print("Masukkan nama bahan makanan yang dicari: ")
		fmt.Scan(&cariNama)

		index := sequentialSearch(data, cariNama)

		if index != -1 {
			fmt.Println("\nHasil Pencarian")
			fmt.Println("---------------------------")
			fmt.Printf("Nama Bahan Makanan : %s\n", data[index].nama)
			fmt.Printf("Stok               : %d\n", data[index].stok)
			fmt.Printf("Kadaluwarsa        : %s\n\n", data[index].kadaluwarsa)
		} else {
			fmt.Println("Data tidak ditemukan.")
		}
	case 3:
		return
	default:
		fmt.Println("Pilihan tidak valid")
	}
}

func konfirmasiKembali() bool {
	var input string
	for {
		fmt.Print("\nApakah Anda ingin kembali ke menu utama? (y/n): ")
		fmt.Scanln(&input)
		if strings.ToLower(input) == "y" {
			return true
		} else if strings.ToLower(input) == "n" {
			return false
		}
		fmt.Println("Input tidak valid, silakan masukkan y atau n")
	}
}

func main() {
	var pilihMenu int

	data := []BahanMakanan{
		{"Minyak", 3, "2026-03-02"},
		{"Gula", 2, "2025-05-15"},
		{"Kopi", 12, "2026-09-09"},
		{"Minyak", 6, "2025-05-27"},
		{"Telur", 0, "2025-05-27"},
		{"Beras", 10, "2025-05-09"},
	}

	for {
		fmt.Println("\nSelamat Datang di ✨ Aplikasi Manajeman Stok Bahan Makanan ✨")

		fmt.Println("\nDaftar Bahan Makanan")
		fmt.Println("===============================================")
		fmt.Printf("%-3s | %-15s | %-5s | %-15s\n",
			"No.", "Nama Bahan", "Stok", "Kadaluwarsa")
		fmt.Println("===============================================")
		for i := 0; i < len(data); i++ {
			fmt.Printf("%-3d | %-15s | %-5d | %-15s\n",
				i+1, data[i].nama, data[i].stok, data[i].kadaluwarsa)
		}
		fmt.Println("===============================================")

		fmt.Println("Daftar Fitur:")
		fmt.Println("1. Lihat Daftar Bahan Makanan")
		fmt.Println("2. Tambah Data Bahan Makanan")
		fmt.Println("3. Ubah Data Bahan Makanan")
		fmt.Println("4. Hapus Bahan Makanan")
		fmt.Println("5. Cari Bahan Makanan")
		fmt.Println("6. Laporan Stok Bahan Makanan")
		tampilkanReminder(data)
		fmt.Print("Pilih menu: ")
		fmt.Scanln(&pilihMenu)

		switch pilihMenu {
		case 1:
			daftarBahanMakanan(data)
			if !konfirmasiKembali() {
				continue
			}
		case 2:
			tambahData(&data)
			if !konfirmasiKembali() {
				continue
			}
		case 3:
			ubahData(&data)
			if !konfirmasiKembali() {
				continue
			}
		case 4:
			hapusData(&data)
			if !konfirmasiKembali() {
				continue
			}
		case 5:
			binerSearch(data)
			if !konfirmasiKembali() {
				continue
			}
		case 6:
			laporanStok(data)
			if !konfirmasiKembali() {
				continue
			}
		default:
			fmt.Println("Menu tidak valid")
			if !konfirmasiKembali() {
				continue
			}
		}
	}
}
