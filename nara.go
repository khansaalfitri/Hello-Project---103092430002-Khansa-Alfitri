// ini merupakan tugas besar mata kuliah alpro, dan komentar untuk tugas Git Wawasan Global TIK
package main

import (
	"fmt"
)

const MAX = 100

type ppasien struct {
	namaPasien, tanggal, hasilMCU string
	IDpasien                      int
}

type paketMCU struct {
	namaPaket      string
	IDpaket, harga int
}

var datapaket [MAX]paketMCU
var datapasien [MAX]ppasien
var jumlahpaket, jumlahpasien int

func main() {
	var pilih int
	var jalan bool

	jalan = true
	for jalan {
		fmt.Println("\n=== MENU MEDICAL CHECK-UP ===")
		fmt.Println("1. Kelola Data")
		fmt.Println("2. Laporan Pemasukan Harian")
		fmt.Println("3. Pencarian Data Pasien")
		fmt.Println("4. Urutkan Data Pasien")
		fmt.Println("5. Pencarian berdasarkan Tanggal MCU (Binary Search)")
		fmt.Println("0. Keluar")
		fmt.Print("Pilih menu: ")
		fmt.Scan(&pilih)

		if pilih == 1 {
			KelolaData()
		} else if pilih == 2 {
			var tanggal, bulan, tahun string
			// Minta 3 input terpisah dari pengguna
			fmt.Print("Masukkan Tanggal Laporan (DD): ")
			fmt.Scan(&tanggal)
			fmt.Print("Masukkan Bulan Laporan (MM): ")
			fmt.Scan(&bulan)
			fmt.Print("Masukkan Tahun Laporan (YYYY): ")
			fmt.Scan(&tahun)

			// Panggil fungsi dengan tiga argumen
			laporanPemasukan(tanggal, bulan, tahun)

		} else if pilih == 3 {
			var nama string
			fmt.Print("Masukkan nama pasien yang dicari (gunakan '_' untuk spasi): ")
			fmt.Scan(&nama)
			idx := cariNamapasien(nama)
			if idx != -1 {
				fmt.Println("Data ditemukan:", datapasien[idx])
			} else {
				fmt.Println("Data tidak ditemukan")
			}
		} else if pilih == 4 {
			var urutan, asc int
			fmt.Println("1. Urutkan berdasarkan Tanggal MCU")
			fmt.Println("2. Urutkan berdasarkan Nama")
			fmt.Print("Pilih: ")
			fmt.Scan(&urutan)
			fmt.Print("1. Ascending, 2. Descending: ")
			fmt.Scan(&asc)
			ascending := asc == 1
			if urutan == 1 {
				sortTanggalMCU(ascending)
			} else if urutan == 2 {
				sortNamaPasien(ascending)
			}
			tampilkandatapasien()
		} else if pilih == 5 {
			var tanggal string
			fmt.Print("Masukkan tanggal (YYYY-MM-DD): ")
			fmt.Scan(&tanggal)
			sortTanggalMCU(true)
			idx := binarySearchTanggal(tanggal)
			if idx != -1 {
				fmt.Println("Data ditemukan:", datapasien[idx])
			} else {
				fmt.Println("Data tidak ditemukan")
			}
		} else if pilih == 0 {
			jalan = false
			fmt.Println("Terima kasih telah menggunakan aplikasi MCU.")
		}
	}
}

func KelolaData() {
	var pilih int
	fmt.Println("\n=== KELOLA DATA ===")
	fmt.Println("1. Tambah Data")
	fmt.Println("2. Edit Data")
	fmt.Println("3. Hapus Data")
	fmt.Println("4. Tampilkan Paket MCU")
	fmt.Println("5. Tampilkan Pasien")
	fmt.Print("Pilih menu: ")
	fmt.Scan(&pilih)

	if pilih == 1 {
		var sub int
		fmt.Println("1. Tambah Paket MCU")
		fmt.Println("2. Tambah Pasien")
		fmt.Print("Pilih: ")
		fmt.Scan(&sub)
		if sub == 1 {
			tambahpaket()
		} else if sub == 2 {
			tambahpasien()
		}
	} else if pilih == 2 {
		var sub int
		fmt.Println("1. Edit Paket MCU")
		fmt.Println("2. Edit Pasien")
		fmt.Print("Pilih: ")
		fmt.Scan(&sub)
		if sub == 1 {
			editpaket()
		} else if sub == 2 {
			var nama string
			fmt.Print("Masukkan nama pasien (gunakan '_' untuk spasi): ")
			fmt.Scan(&nama)
			editpasien(nama)
		}
	} else if pilih == 3 {
		var sub int
		fmt.Println("1. Hapus Paket MCU")
		fmt.Println("2. Hapus Pasien")
		fmt.Print("Pilih: ")
		fmt.Scan(&sub)
		if sub == 1 {
			hapuspaket()
		} else if sub == 2 {
			var nama string
			fmt.Print("Masukkan nama pasien (gunakan '_' untuk spasi): ")
			fmt.Scan(&nama)
			hapuspasien(nama)
		}
	} else if pilih == 4 {
		tampilkanPaket()
	} else if pilih == 5 {
		tampilkandatapasien()
	} else {
		fmt.Println("Pilihan tidak valid")
	}
}

func tambahpaket() {
	if jumlahpaket < MAX {
		fmt.Print("ID Paket: ")
		fmt.Scan(&datapaket[jumlahpaket].IDpaket)
		fmt.Print("Nama Paket (gunakan '_' untuk spasi): ")
		fmt.Scan(&datapaket[jumlahpaket].namaPaket)
		fmt.Print("Harga Paket: ")
		fmt.Scan(&datapaket[jumlahpaket].harga)
		jumlahpaket++
	} else {
		fmt.Println("Data penuh.")
	}
}

func tambahpasien() {
	if jumlahpasien < MAX {
		fmt.Print("Nama Pasien (gunakan '_' untuk spasi): ")
		fmt.Scan(&datapasien[jumlahpasien].namaPasien)
		fmt.Print("Tanggal MCU (YYYY-MM-DD): ")
		fmt.Scan(&datapasien[jumlahpasien].tanggal)
		fmt.Print("ID Paket yang dipilih: ")
		fmt.Scan(&datapasien[jumlahpasien].IDpasien)
		jumlahpasien++
	} else {
		fmt.Println("Data pasien penuh.")
	}
}

func editpaket() {
	var id int
	fmt.Print("Masukkan ID Paket yang ingin diedit: ")
	fmt.Scan(&id)
	for i := 0; i < jumlahpaket; i++ {
		if datapaket[i].IDpaket == id {
			fmt.Print("Nama Paket Baru (gunakan '_' untuk spasi): ")
			fmt.Scan(&datapaket[i].namaPaket)
			fmt.Print("Harga Baru: ")
			fmt.Scan(&datapaket[i].harga)
			return
		}
	}
	fmt.Println("Paket tidak ditemukan")
}

func editpasien(nama string) {
	idx := cariNamapasien(nama)
	if idx != -1 {
		fmt.Print("Tanggal MCU baru (YYYY-MM-DD): ")
		fmt.Scan(&datapasien[idx].tanggal)
		fmt.Print("ID Paket baru: ")
		fmt.Scan(&datapasien[idx].IDpasien)
	} else {
		fmt.Println("Pasien tidak ditemukan")
	}
}

func hapuspaket() {
	var id int
	fmt.Print("Masukkan ID Paket yang akan dihapus: ")
	fmt.Scan(&id)
	for i := 0; i < jumlahpaket; i++ {
		if datapaket[i].IDpaket == id {
			for j := i; j < jumlahpaket-1; j++ {
				datapaket[j] = datapaket[j+1]
			}
			jumlahpaket--
			fmt.Println("Paket dihapus")
			return
		}
	}
	fmt.Println("Paket tidak ditemukan")
}

func hapuspasien(nama string) {
	idx := cariNamapasien(nama)
	if idx != -1 {
		for i := idx; i < jumlahpasien-1; i++ {
			datapasien[i] = datapasien[i+1]
		}
		jumlahpasien--
		fmt.Println("Pasien dihapus")
	} else {
		fmt.Println("Pasien tidak ditemukan")
	}
}

func tampilkanPaket() {
	fmt.Println("\n--- Data Paket MCU ---")
	for i := 0; i < jumlahpaket; i++ {
		fmt.Printf("ID: %d - Nama: %s - Harga: Rp%d\n", datapaket[i].IDpaket, datapaket[i].namaPaket, datapaket[i].harga)
	}
}

func tampilkandatapasien() {
	fmt.Println("\n--- Data Pasien ---")
	for i := 0; i < jumlahpasien; i++ {
		fmt.Printf("Nama: %s, Tanggal MCU: %s, ID Paket: %d\n",
			datapasien[i].namaPasien, datapasien[i].tanggal, datapasien[i].IDpasien)
	}
}

func laporanPemasukan(tanggal, bulan, tahun string) {
	// Gabungkan input terpisah menjadi satu string dengan format YYYY-MM-DD
	tanggalDicari := tahun + "-" + bulan + "-" + tanggal

	total := 0
	dataDitemukan := false

	// Loop melalui semua data pasien
	for i := 0; i < jumlahpasien; i++ {
		// Langsung bandingkan string tanggal yang tersimpan dengan tanggal yang dicari
		if datapasien[i].tanggal == tanggalDicari {
			dataDitemukan = true // Tandai jika ada data pada tanggal ini
			// Cari harga paket berdasarkan ID Paket pasien
			for j := 0; j < jumlahpaket; j++ {
				if datapaket[j].IDpaket == datapasien[i].IDpasien {
					total += datapaket[j].harga // Asumsi ID paket unik, keluar dari loop paket
				}
			}
		}
	}

	// Beri output berdasarkan apakah data ditemukan atau tidak
	if dataDitemukan {
		fmt.Printf("Total pemasukan untuk tanggal %s: Rp%d\n", tanggalDicari, total)
	} else {
		fmt.Printf("Tidak ada data pemasukan untuk tanggal %s.\n", tanggalDicari)
	}
}

// Sequential Search berdasarkan nama
func cariNamapasien(nama string) int {
	for i := 0; i < jumlahpasien; i++ {
		if datapasien[i].namaPasien == nama {
			return i
		}
	}
	return -1
}

// Binary Search yang dimodifikasi untuk menemukan indeks PERTAMA dari data yang cocok.
// Fungsi ini tetap mewajibkan data sudah terurut menaik (ascending).
// ascending : naik
// descending : turun
func binarySearchTanggal(tanggal string) int {
	sortTanggalMCU(true)

	kiri, kanan := 0, jumlahpasien-1
	idx := -1 // Variabel untuk menyimpan indeks terakhir yang ditemukan

	for kiri <= kanan {
		tengah := kiri + (kanan-kiri)/2

		if datapasien[tengah].tanggal == tanggal {
			// Data ditemukan!
			idx = tengah       // Simpan indeks ini.
			kanan = tengah - 1 // Lanjutkan pencarian di sisi KIRI untuk menemukan indeks yang lebih awal.
		} else if datapasien[tengah].tanggal < tanggal {
			// Data yang dicari ada di sisi kanan
			kiri = tengah + 1
		} else {
			// Data yang dicari ada di sisi kiri
			kanan = tengah - 1
		}
	}

	return idx // Kembalikan indeks terakhir (paling kiri) yang ditemukan.
}

// Selection Sort TanggalMCU Asc/Desc
// ascending : naik
// descending : turun
func sortTanggalMCU(ascending bool) {
	for i := 0; i < jumlahpasien-1; i++ {
		minIdx := i
		for j := i + 1; j < jumlahpasien; j++ {
			if ascending {
				if datapasien[j].tanggal < datapasien[minIdx].tanggal {
					minIdx = j
				}
			} else {
				if datapasien[j].tanggal > datapasien[minIdx].tanggal {
					minIdx = j
				}
			}
		}
		// datapasien[i], datapasien[minIdx] = datapasien[minIdx], datapasien[i]

		temp := datapasien[i]
		datapasien[i] = datapasien[minIdx]
		datapasien[minIdx] = temp
	}
}

// Insertion Sort berdasarkan Nama Pasien Asc/Desc
// ascending : naik
// descending : turun
func sortNamaPasien(ascending bool) {
	for i := 1; i < jumlahpasien; i++ {
		temp := datapasien[i]
		j := i - 1
		if ascending {
			for j >= 0 && datapasien[j].namaPasien > temp.namaPasien {
				datapasien[j+1] = datapasien[j]
				j--
			}
		} else {
			for j >= 0 && datapasien[j].namaPasien < temp.namaPasien {
				datapasien[j+1] = datapasien[j]
				j--
			}
		}
		datapasien[j+1] = temp
	}
}
