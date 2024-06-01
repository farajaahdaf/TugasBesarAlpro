package main

import (
	"fmt"
)

const NMAX = 1000

type tabData struct {
	Name     string
	Price    float64
	Quantity int
}

type Transaksi struct {
	jumlahTerjual float64
	Price         float64
}

type dataTab [NMAX]tabData
type dataTransaksi [NMAX]Transaksi

var keluarAplikasi bool

func main() {
	var barang dataTab
	var transaksi dataTransaksi
	var index int
	var pendapatan float64

	keluarAplikasi = false
	for !keluarAplikasi {
		fmt.Println("*** ------------------------------------------- ***")
		fmt.Println("***            Aplikasi Kasir Minimart          ***")
		fmt.Println("***               Created by :                  ***")
		fmt.Println("***      - Mohammad Faraja Ahdaf (103012300178) ***")
		fmt.Println("***      - Jiyad Arsal Asari  (103012300351)    ***")
		fmt.Println("*** ------------------------------------------- ***")
		fmt.Println("1. Tambah Barang")
		fmt.Println("2. Ubah Barang")
		fmt.Println("3. Hapus Barang")
		fmt.Println("4. Lihat Daftar Barang")
		fmt.Println("5. Catat Transaksi")
		fmt.Println("6. Lihat Daftar Transaksi dan Omset Harian")
		fmt.Println("7. Keluar")
		fmt.Print("Pilih menu: ")

		var choice int
		fmt.Scanln(&choice)

		if choice == 1 {
			tambahBarang(&barang, index)
		} else if choice == 2 {
			ubahBarang(&barang, index)
		} else if choice == 3 {
			hapusBarang(&barang, index, NMAX)
		} else if choice == 4 {
			tampilkanDaftarBarang(&barang)
		} else if choice == 5 {
			catatTransaksi(&barang, &transaksi, &pendapatan, index)
		} else if choice == 6 {
			tampilkanDaftarTransaksi(barang, &transaksi, pendapatan)
		} else if choice == 7 {
			fmt.Println("Berhasil keluar aplikasi.")
			keluarAplikasi = true
		}
	}
}

func kembaliKeMenu() {
	fmt.Println("Klik Enter untuk kembali ke menu utama...")
	var input string
	fmt.Scanln(&input)
	if input != "" {
		keluarAplikasi = true
	}

}

func tambahBarang(A *dataTab, index int) {
	var i int
	// Mencari indeks yang tersedia untuk menambahkan barang baru
	index = -1
	for i = 0; i < NMAX; i++ {
		if A[i].Name == "" {
			index = i
			break
		}
	}

	if index == -1 {
		fmt.Println("Tidak bisa menambahkan barang baru, kapasitas penuh.")
		kembaliKeMenu() 
		return
	}

	// Memasukkan data barang baru
	fmt.Println("Masukkan data barang baru:")
	fmt.Print("Nama: ")
	var nama string
	fmt.Scanln(&nama)
	fmt.Print("Harga: ")
	var harga float64
	fmt.Scanln(&harga)
	fmt.Print("Kuantitas: ")
	var kuantitas int
	fmt.Scanln(&kuantitas)

	// Mengecek apakah barang sudah ada dalam daftar
	if sequentialSearchItem(nama, A) != -1 {
		fmt.Println("Barang sudah ada dalam daftar.")
		fmt.Print("Input ulang barang ? (y/n): ")
		var choice string
		fmt.Scanln(&choice)
		if choice != "y" && choice != "Y" {
			kembaliKeMenu()
			return
		}
	} else {
		// Menambahkan barang ke dalam daftar
		A[index].Name = nama
		A[index].Price = harga
		A[index].Quantity = kuantitas
		fmt.Println("Barang berhasil ditambahkan.")
	}
	kembaliKeMenu()
}

func ubahBarang(A *dataTab, index int) {
	fmt.Print("Masukkan nama barang yang mau diubah: ")
	var namaBarang string
	fmt.Scanln(&namaBarang)

	index = binarySearchItem(namaBarang, A)

	if index != -1 {
		fmt.Print("Masukkan harga baru: ")
		var hargaBaru float64
		fmt.Scanln(&hargaBaru)
		fmt.Print("Masukkan kuantitas baru: ")
		var kuantitasBaru int
		fmt.Scanln(&kuantitasBaru)

		A[index].Price = hargaBaru
		A[index].Quantity = kuantitasBaru
		fmt.Println("Barang berhasil diubah.")
	} else {
		fmt.Println("Nama produk tidak ditemukan.")
	}
	kembaliKeMenu()
}

func hapusBarang(A *dataTab, index int, len int) {
	fmt.Print("Masukkan nama barang yang ingin dihapus: ")
	var namaBarang string
	fmt.Scanln(&namaBarang)

	index = sequentialSearchItem(namaBarang, A)

	if index != -1 {
		// Menggeser elemen-elemen setelah elemen yang dihapus satu posisi ke kiri
		for i := index; i < len-1; i++ {
			A[i] = A[i+1]
		}
		// Mengosongkan elemen terakhir
		len--

		fmt.Println("Barang berhasil dihapus.")
	} else {
		fmt.Println("Barang tidak ditemukan.")
	}
	kembaliKeMenu()
}

func tampilkanDaftarBarang(A *dataTab) {
	var pilihan int
	fmt.Println("Pilih kategori pengurutan:")
	fmt.Println("1. Tampilkan Daftar Barang dari Harga Tertinggi (Descending)")
	fmt.Println("2. Tampilkan Daftar Barang dari Harga Terendah (Ascending)")
	fmt.Println("3. Tampilkan Daftar Barang dari Kuantitas Terbanyak (Descending)")
	fmt.Println("4. Tampilkan Daftar Barang dari Kuantitas Terendah (Ascending)")
	fmt.Print("Pilihan: ")
	fmt.Scanln(&pilihan)

	if pilihan == 1 {
    	selectionSort(A, "price", "desc")
	} else if pilihan == 2 {
    	selectionSort(A, "price", "asc")
	} else if pilihan == 3 {
    	insertionSort(A, "quantity", "desc")
	} else if pilihan == 4 {
    	insertionSort(A, "quantity", "asc")
	} else {
    	fmt.Println("Pilihan tidak valid.")
    	kembaliKeMenu()
    	return
	}

	fmt.Println("Daftar Barang :")
	fmt.Println("Nama | Harga | Kuantitas")
	for i := 0; i < NMAX && A[i].Name != ""; i++ {
		fmt.Printf("%s | Rp. %.2f | %d\n", A[i].Name, A[i].Price, A[i].Quantity)
	}

	kembaliKeMenu()
}

func catatTransaksi(A *dataTab, transaksi *dataTransaksi, pendapatan *float64, index int) {
	fmt.Print("Masukkan nama barang: ")
	var namaBarang string
	fmt.Scanln(&namaBarang)

	index = sequentialSearchItem(namaBarang, A)

	if index != -1 {
		fmt.Print("Jumlah yang terjual: ")
		var jumlahTerjual float64
		fmt.Scanln(&jumlahTerjual)

		if jumlahTerjual > float64(A[index].Quantity) {
			fmt.Println("Jumlah terjual melebihi kuantitas yang ada.")
		} else {
			// Mengupdate transaksi
			transaksi[index].jumlahTerjual += jumlahTerjual
			transaksi[index].Price = A[index].Price

			// Mengupdate pendapatan
			*pendapatan += A[index].Price * jumlahTerjual

			// Mengupdate stok barang
			A[index].Quantity -= int(jumlahTerjual)

			fmt.Println("Transaksi berhasil dicatat.")
		}
	} else {
		fmt.Println("Barang tidak ditemukan.")
	}
	kembaliKeMenu()
}

func tampilkanDaftarTransaksi(A dataTab, transaksi *dataTransaksi, pendapatan float64) {
	fmt.Println("=== Daftar Transaksi ===")
	var totalHarga float64
	var i int
	if pendapatan == 0 {
		fmt.Println("Belum ada transaksi yang dicatat.")
	} else {
		fmt.Println("No | Nama Barang | Jumlah Terjual | Harga Sebelum Diubah | Total Harga")
		for i = 0; i < NMAX; i++ {
			if transaksi[i].jumlahTerjual > 0 {
				totalHarga = transaksi[i].jumlahTerjual * transaksi[i].Price
				fmt.Printf("%d. %s | %.0f Item | Rp. %.2f | Rp. %.2f\n", i+1, A[i].Name, transaksi[i].jumlahTerjual, transaksi[i].Price, totalHarga)
			}
		}
		fmt.Printf("Total Omset hari ini: Rp. %.2f\n", pendapatan)
	}
	kembaliKeMenu()
}

// Tempat fungsi Sequential Search, BinarySearch dan Sorting
func sequentialSearchItem(name string, A *dataTab) int {
	var i int
	for i = 0; i < NMAX; i++ {
		if A[i].Name == name {
			return i
		}
	}
	return -1
}

func binarySearchItem(name string, A *dataTab) int {
	// Mencari batas array yang terisi
	n := 0
	for n < NMAX && A[n].Name != "" {
		n++
	}

	left, right := 0, n-1
	for left <= right {
		mid := (left + right) / 2
		if A[mid].Name == name {
			return mid
		}
		if A[mid].Name < name {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return -1
}

func selectionSort(A *dataTab, category string, order string) {
	n := 0
	for n < NMAX && A[n].Name != "" {
		n++
	}

	for pass := 0; pass < n-1; pass++ {
		idx := pass
		for i := pass + 1; i < n; i++ {
			if (order == "asc" && ((category == "price" && A[i].Price < A[idx].Price) || (category == "quantity" && A[i].Quantity < A[idx].Quantity))) ||
				(order == "desc" && ((category == "price" && A[i].Price > A[idx].Price) || (category == "quantity" && A[i].Quantity > A[idx].Quantity))) {
				idx = i
			}
		}
		A[pass], A[idx] = A[idx], A[pass]
	}
}

func insertionSort(A *dataTab, category string, order string) {
	n := 0
	for n < NMAX && A[n].Name != "" {
		n++
	}

	for i := 1; i < n; i++ {
		key := A[i]
		j := i - 1

		for j >= 0 && ((order == "asc" && ((category == "price" && A[j].Price > key.Price) || (category == "quantity" && A[j].Quantity > key.Quantity))) ||
			(order == "desc" && ((category == "price" && A[j].Price < key.Price) || (category == "quantity" && A[j].Quantity < key.Quantity)))) {
			A[j+1] = A[j]
			j--
		}
		A[j+1] = key
	}
}
