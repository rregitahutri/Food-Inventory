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

func tampilkanReminder(bahanMakanann []BahanMakanan) {
	var count int

	// Header
	fmt.Println("┌───────────────────────────────────────────────────────────┐")
	fmt.Println("│        ⚠️  PERINGATAN KADALUWARSA BAHAN MAKANAN ⚠️          │")
	fmt.Println("├──────┬────────────────┬──────────────────────┬────────────┤")
	fmt.Println("│ No.  │ Nama Bahan     │ Status               │ Kadaluwarsa│")
	fmt.Println("├──────┼────────────────┼──────────────────────┼────────────┤")

	today := time.Now().Truncate(24 * time.Hour)

	for i, bahan := range bahanMakanann {
		kadaluwarsa, err := time.Parse("2006-01-02", bahan.kadaluwarsa)
		if err != nil {
			continue
		}

		// Hitung selisih hari
		selisih := int(kadaluwarsa.Sub(today).Hours() / 24)
		status := ""

		if selisih < 0 {
			status = "LEWAT KADALUWARSA"
		} else if selisih == 0 {
			status = "KADALUWARSA HARI INI"
		} else if selisih <= 7 {
			status = fmt.Sprintf("%d HARI", selisih)
		} else {
			continue // Lewati bahan yang belum mendekati kadaluwarsa (> 7 hari)
		}

		fmt.Printf("│ %-4d │ %-14s │ %-20s │ %s │\n", i+1, bahan.nama, status, bahan.kadaluwarsa)
		count++
	}

	if count == 0 {
		fmt.Println("├──────┴────────────────┴──────────────────────┴────────────┤")
		fmt.Println("│ ✓ Tidak ada bahan yang kadaluwarsa hari ini/7 hari/lewat │")
	} else {
		fmt.Println("├──────┴────────────────┴──────────────────────┴────────────┤")
		fmt.Printf("│ Total: %d bahan perlu perhatian khusus                     │\n", count)
	}
	fmt.Println("└───────────────────────────────────────────────────────────┘")
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

func lihatDaftar(data []BahanMakanan) {
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

	// Input tanggal kadaluwarsa dengan validasi
	for {
		fmt.Print("Masukkan tanggal kadaluwarsa (format: yyyy-mm-dd): ")
		fmt.Scanln(&kadaluwarsa)
		_, err := time.Parse("2006-01-02", kadaluwarsa)
		if err != nil {
			fmt.Println("Error: Format tanggal tidak valid. Gunakan format yyyy-mm-dd.")
			continue
		}
		break
	}

	// Tambahkan data baru
	*data = append(*data, BahanMakanan{
		nama:        nama,
		stok:        stok,
		kadaluwarsa: kadaluwarsa,
	})

	fmt.Println("\nBahan makanan berhasil ditambahkan!")
	fmt.Printf("Nama: %s\nStok: %d\nKadaluwarsa: %s\n\n", nama, stok, kadaluwarsa)
}

func ubahData(data *[]BahanMakanan) {
	var nomor int
	fmt.Print("\n» Masukkan nomor bahan yang akan diubah: ")
	_, err := fmt.Scanln(&nomor)
	if err != nil || nomor < 1 || nomor > len(*data) {
		fmt.Println("Nomor tidak valid")
		return
	}

	bahan := &(*data)[nomor-1]
	fmt.Printf("\nData yang akan diubah:\nNama: %s\nStok: %d\nKadaluwarsa: %s\n",
		bahan.nama, bahan.stok, bahan.kadaluwarsa)

	// Input sementara
	var inputNama, inputStok, inputKadaluwarsa string

	fmt.Println("Masukkan data baru (beri tanda '-' jika tidak ingin mengubah):")

	// Input Nama
	fmt.Print("Nama : ")
	fmt.Scanln(&inputNama)
	if inputNama != "-" {
		bahan.nama = inputNama
	}

	// Input Stok
	for {
		fmt.Print("Jumlah Stok : ")
		fmt.Scanln(&inputStok)
		if inputStok == "-" {
			break
		}
		var stokBaru int
		_, err := fmt.Sscanf(inputStok, "%d", &stokBaru)
		if err == nil && stokBaru >= 0 {
			bahan.stok = stokBaru
			break
		} else {
			fmt.Println("Input stok tidak valid. Masukkan angka atau '-' untuk batal.")
		}
	}

	// Input Kadaluwarsa
	for {
		fmt.Print("Tanggal Kadaluwarsa (format yyyy-mm-dd): ")
		fmt.Scanln(&inputKadaluwarsa)
		if inputKadaluwarsa == "-" {
			break
		}
		if _, err := time.Parse("2025-05-25", inputKadaluwarsa); err == nil {
			bahan.kadaluwarsa = inputKadaluwarsa
			break
		} else {
			fmt.Println("Format tanggal tidak valid. Coba lagi atau masukkan '-' untuk batal.")
		}
	}

	fmt.Printf("\nData berhasil diubah:\nNama: %s\nStok: %d\nKadaluwarsa: %s\n\n",
		bahan.nama, bahan.stok, bahan.kadaluwarsa)
}

func hapusData(data *[]BahanMakanan) {
	// Cek jika data kosong
	if len(*data) == 0 {
		fmt.Println("\nTidak ada data bahan makanan yang bisa dihapus")
		return
	}

	// Input nomor yang akan dihapus
	var nomor int
	fmt.Print("Masukkan nomor bahan yang akan dihapus: ")
	_, err := fmt.Scanln(&nomor)
	if err != nil {
		fmt.Println("\nInput tidak valid")
		return
	}

	// Validasi nomor
	if nomor < 1 || nomor > len(*data) {
		fmt.Println("\nNomor tidak valid")
		return
	}

	// Proses penghapusan dengan pergeseran elemen
	bahanTerhapus := (*data)[nomor-1].nama
	for i := nomor - 1; i < len(*data)-1; i++ {
		(*data)[i] = (*data)[i+1]
	}
	*data = (*data)[:len(*data)-1]

	fmt.Printf("\nBahan makanan '%s' berhasil dihapus\n", bahanTerhapus)
}

func laporanStok(data []BahanMakanan) {
	total := len(data)
	digunakan := 0
	tersedia := 0

	// Hitung stok
	for i := 0; i < total; i++ {
		if data[i].stok == 0 {
			digunakan++
		} else {
			tersedia++
		}
	}

	// Tampilkan ringkasan
	fmt.Println("\nLaporan Stok Bahan Makanan")
	fmt.Println("===============================================")
	fmt.Printf("Total Bahan Makanan        : %d\n", total)
	fmt.Printf("Bahan Tersedia (stok > 0)  : %d\n", tersedia)
	fmt.Printf("Bahan Telah Digunakan      : %d\n", digunakan)
	fmt.Println("===============================================\n")

	if total == 0 {
		fmt.Println("Tidak ada data bahan makanan.")
		return
	}

	// Tampilkan bahan tersedia
	fmt.Println("Daftar Bahan Tersedia:")
	fmt.Printf("%-5s %-20s %-10s %-15s\n", "No", "Nama", "Stok", "Kadaluwarsa")
	fmt.Println("------------------------------------------------------------")
	no := 1
	for i := 0; i < total; i++ {
		if data[i].stok > 0 {
			fmt.Printf("%-5d %-20s %-10d %-15s\n", no, data[i].nama, data[i].stok, data[i].kadaluwarsa)
			no++
		}
	}
	fmt.Println("------------------------------------------------------------\n")
	fmt.Println("Daftar Bahan Telah Digunakan (Stok = 0):")
	fmt.Printf("%-5s %-20s %-10s %-15s\n", "No", "Nama", "Stok", "Kadaluwarsa")
	fmt.Println("------------------------------------------------------------")
	no = 1
	for i := 0; i < total; i++ {
		if data[i].stok == 0 {
			fmt.Printf("%-5d %-20s %-10d %-15s\n", no, data[i].nama, data[i].stok, data[i].kadaluwarsa)
			no++
		}
	}
	fmt.Println("------------------------------------------------------------")
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

func binarySearch(data []BahanMakanan) {
	// Urutkan data terlebih dahulu
	sorted := sortByNama(data)

	// Tampilkan data terurut
	fmt.Println("\nData Bahan Makanan (Terurut berdasarkan Nama):")
	fmt.Println("===============================================")
	fmt.Printf("%-3s | %-15s | %-5s | %-15s\n", "No.", "Nama Bahan", "Stok", "Kadaluwarsa")
	fmt.Println("===============================================")
	for i := 0; i < len(sorted); i++ {
		fmt.Printf("%-3d | %-15s | %-5d | %-15s\n",
			i+1, sorted[i].nama, sorted[i].stok, sorted[i].kadaluwarsa)
	}
	fmt.Println("===============================================")

	// Input nama yang dicari
	var cariNama string
	fmt.Print("\nMasukkan nama bahan makanan yang dicari: ")
	fmt.Scan(&cariNama)

	// Lakukan pencarian
	index := pencarianCepat(sorted, cariNama)

	// Tampilkan hasil
	if index != -1 {
		fmt.Println("\nHasil Pencarian")
		fmt.Println("---------------------------")
		fmt.Printf("Nama Bahan Makanan : %s\n", sorted[index].nama)
		fmt.Printf("Stok               : %d\n", sorted[index].stok)
		fmt.Printf("Kadaluwarsa        : %s\n\n", sorted[index].kadaluwarsa)
	} else {
		fmt.Println("Data tidak ditemukan.")
	}
}

func cariData(data []BahanMakanan) {
	var pilihCari int

	fmt.Println("\nPilih opsi pencarian")
	fmt.Println("1. Pencarian Cepat & Terurut (Binary Search)")
	fmt.Println("2. Pencarian Biasa (Sequential Search)")
	fmt.Println("3. Kembali ke Menu Utama")
	fmt.Print("Pilih : ")
	fmt.Scan(&pilihCari)

	switch pilihCari {
	case 1:
		binarySearch(data)
	case 2:
		var cariNama string
		fmt.Print("\nMasukkan nama bahan makanan yang dicari: ")
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
		} else {
			fmt.Println("Input tidak valid, silakan masukkan y atau n")
		}
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

		fmt.Println("╔───────────────────────────────────────────────────────╗")
		fmt.Println("                Daftar Bahan Makanan                    ")
		fmt.Println("╚═══════════════════════════════════════════════════════╝")
		fmt.Println("┌─────┬──────────────────┬────────┬───────────────────┐")
		fmt.Printf("│ %-3s │ %-16s │ %-5s  │ %-17s │\n",
			"No.", "Nama Bahan", "Stok", "Kadaluwarsa")
		fmt.Println("├─────┼──────────────────┼────────┼───────────────────┤")
		for i := 0; i < len(data); i++ {
			fmt.Printf("│ %-3d │ %-16s │ %-5d  │ %-17s │\n",
				i+1, data[i].nama, data[i].stok, data[i].kadaluwarsa)
		}
		fmt.Println("└─────┴──────────────────┴────────┴───────────────────┘")

		fmt.Println("╔════════════════════════════════════════════╗")
		fmt.Println("║              FITUR APLIKASI                ║")
		fmt.Println("╠════════════════════════════════════════════╣")
		fmt.Println("║   📋 Lihat Daftar Bahan Makanan            ║")
		fmt.Println("║   ➕ Tambah Data Bahan Makanan             ║")
		fmt.Println("║   ✏️  Ubah Data Bahan Makanan               ║")
		fmt.Println("║   🗑️  Hapus Bahan Makanan                   ║")
		fmt.Println("║   🔍 Cari Bahan Makanan                    ║")
		fmt.Println("║   📊 Laporan Stok Bahan Makanan            ║")
		fmt.Println("╠════════════════════════════════════════════╣")
		fmt.Println("║   🚪 Keluar                                ║")
		fmt.Println("╚════════════════════════════════════════════╝")

		tampilkanReminder(data)
		fmt.Print("Pilih menu: ")
		fmt.Scanln(&pilihMenu)

		switch pilihMenu {
		case 1:
			lihatDaftar(data)
		case 2:
			tambahData(&data)
		case 3:
			ubahData(&data)
		case 4:
			hapusData(&data)
		case 5:
			cariData(data)
		case 6:
			laporanStok(data)
		default:
			fmt.Println("Menu tidak valid")
		}

		if !konfirmasiKembali() {
			fmt.Println("Terima kasih telah menggunakan aplikasi ini 🙌")
			return
		}
	}
}
