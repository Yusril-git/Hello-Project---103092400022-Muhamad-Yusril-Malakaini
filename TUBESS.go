package main

import (
	"fmt"
)

const NMAX int = 100


// Perubahan oleh Muhamad Yusril Malakaini untuk tugas Git

// edit by yusril malakaini

type komik struct{
	id string
	judul string
	penulis string
	pubYear int
}

type member struct{
	id string
	nama string
	Daftar int
}

type peminjaman struct{
	idMember string
	idKomik string
	lamaPinjam string
	kembalikan int
}

var dataKomik [NMAX]komik
var dataMember [NMAX]member
var dataPinjam [NMAX]peminjaman
var jumlahKomik,jumlahMember,jumlahPinjam int
var totalUangDenda int

func main() {

	var menu int
	for menu != 9 {
		fmt.Println("\n-----------------------------")
		fmt.Println("      M E N U   U T A M A      ")
		fmt.Println("-----------------------------")
		fmt.Println("1. Kelola Komik")
		fmt.Println("2. Kelola Member")
		fmt.Println("3. Kelola Peminjaman")
		fmt.Println("4. Tampilkan Semua Data")
		fmt.Println("5. Menu ngubah urutan komik")
		fmt.Println("6. Cari Komik/Member (by ID)")
		fmt.Println("7. Hitung Denda & Pengembalian")
		fmt.Println("8. Hitung Total Uang Masuk")
		fmt.Println("9. Exit")
		fmt.Println("-----------------------------")
		fmt.Print("Pilih menu: ")
		fmt.Scan(&menu)
		
		if menu == 1 {
			menuKelolaKomik()
		} else if menu == 2 {
			menuKelolaMember()
			} else if menu == 3 {
			menuKelolaPeminjaman()
			} else if menu == 4 {
				menuTampilkanData()
		} else if menu == 5 {
			menuUrutan()
			} else if menu == 6 {
				menuCariKomikAtauMember()
		} else if menu == 7 {
			hitungDenda()
		} else if menu == 8 {
			hitungTotalUangMasuk()
			} else if menu == 9 {
				fmt.Println("Keluar dari program. Terima kasih!")
		} else {
			fmt.Println("Pilihan tidak valid.")
		}
	}
}

func menuKelolaMember() {
	var pilihan int
	for pilihan != 5 {
		fmt.Println("\n--- Kelola Member ---")
		fmt.Println("1. Tambah Member")
		fmt.Println("2. Edit Member")
		fmt.Println("3. Hapus Member")
		fmt.Println("4. Tampilkan semua member")
		fmt.Println("5. Kembali ke Menu Utama")
		fmt.Print("Pilihan: ")
		fmt.Scan(&pilihan)
		if pilihan == 1 {
			tambahMember()
		} else if pilihan == 2 {
			editMember()
		} else if pilihan == 3 {
			hapusMember()
		} else if pilihan == 4 {
			tampilkanSemuaMember()
		}
	}
}

func menuKelolaKomik() {
	var pilihan int
	for pilihan != 5 {
		fmt.Println("\n--- Kelola Komik ---")
		fmt.Println("1. Tambah Komik")
		fmt.Println("2. Edit Komik")
		fmt.Println("3. Hapus Komik")
		fmt.Println("4. Tampilkan semua komik")
		fmt.Println("5. Kembali ke Menu Utama")
		fmt.Print("Pilihan: ")
		fmt.Scan(&pilihan)
		if pilihan == 1 {
			tambahKomik()
		} else if pilihan == 2 {
			editKomik()
		} else if pilihan == 3 {
			hapusKomik()
		} else if pilihan == 4 {
			tampilkanSemuaKomik()
		}
	}
}

func seqSearchkomik(id string) int{
	for i:=0;i<jumlahKomik;i++{
		if dataKomik[i].id == id {
			return i
		}
	}
	return -1
}

func seqSearchmember(id string)int{
	for i:=0;i<jumlahMember;i++{
		if dataMember[i].id == id{
			return i
		}
	}
	return -1
}

func cariPeminjaman(idM,idK string) int {
	for i:=0;i<jumlahPinjam;i++{
		if dataPinjam[i].idMember == idM && dataPinjam[i].idKomik == idK {
			return i
		}
	}
	return -1
}

func tambahKomik(){
	var n int
	fmt.Println("masukan berapa member yang ingin di tambahkan? ")
	fmt.Scan(&n)

	if jumlahKomik+n > NMAX{
		fmt.Printf("Kapasitas tidak cukup. Anda hanya bisa menambahkan %d komik lagi.\n", NMAX-jumlahKomik)
		return
	}

	for i:=0;i<n;i++{
		fmt.Printf("\n -----masukan data komik ke-%d-----\n", i+1)
		var k komik
		fmt.Println("masukan ID komik baru: ")
		fmt.Scan(&k.id)
		if seqSearchkomik(k.id) != -1{
			fmt.Println("ID Komik sudah ada. Data ini dilewati.")
		} else {
			fmt.Print("Masukkan Judul: ")
			fmt.Scan(&k.judul)
			fmt.Print("Masukkan Penulis: ")
			fmt.Scan(&k.penulis)
			fmt.Print("Masukkan Tahun Terbit: ")
			fmt.Scan(&k.pubYear)
			dataKomik[jumlahKomik] = k
			jumlahKomik++
			fmt.Println("Komik baru berhasil ditambahkan.")
		}
	}
}

func editKomik(){
	var id string
	fmt.Print("Masukan ID komik yang mau diedit: ")
	fmt.Scan(&id)
	index := seqSearchkomik(id)
	if index == -1{
		fmt.Println("ID komik tidak ditemukan" )
		return
	}
	fmt.Printf("Mengedit Komik: %s - %s\n", dataKomik[index].id, dataKomik[index].judul)
	fmt.Print("Masukkan Judul baru: ")
	fmt.Scan(&dataKomik[index].judul)
	fmt.Print("Masukkan Penulis baru: ")
	fmt.Scan(&dataKomik[index].penulis)
	fmt.Print("Masukkan Tahun Terbit baru: ")
	fmt.Scan(&dataKomik[index].pubYear)
	fmt.Println("Data komik berhasil diperbarui.")
}

func hapusKomik(){
	var id string
	fmt.Print("Masukan ID Komik Yang mau dihapus: ")
	fmt.Scan(&id)
	index := seqSearchkomik(id)
	if index == -1{
		fmt.Println("ID komik tidak ditemukan: ")
	}
	for i:=0;i<jumlahPinjam;i++{
		if dataPinjam[i].idKomik == id && dataPinjam[i].kembalikan == 0{
			fmt.Println("Gagal: Komik sedang dipinjam oleh member", dataPinjam[i].idMember, "dan tidak dapat dihapus.")
			return
		}
	}
	for i:=index;i < jumlahKomik-1;i++{
		dataKomik[i] = dataKomik[i+1]
	}
	jumlahKomik--
	fmt.Println("Komik berhasil dihapus")
}

//member

func tambahMember(){
	var n int
	fmt.Print("Berapa banyak member yang ingin ditambahkan: ")
	fmt.Scan(&n)

	if jumlahMember+n > NMAX{
		fmt.Printf("Kapasitas tidak cukup. Anda hanya bisa menambahkan %d member lagi.\n", NMAX-jumlahMember)
	}

	for i := 0; i < n; i++ {
		fmt.Printf("\n--- Memasukkan data member ke-%d ---\n", i+1)
		var m member
		fmt.Print("Masukkan ID Member baru: ")
		fmt.Scan(&m.id)
		if seqSearchmember(m.id) != -1 {
			fmt.Println("ID Member sudah ada. Data ini dilewati.")
		} else {
			fmt.Print("Masukkan Nama: ")
			fmt.Scan(&m.nama)
			fmt.Print("Masukkan Biaya Pendaftaran: ")
			fmt.Scan(&m.Daftar)
			dataMember[jumlahMember] = m
			jumlahMember++
			fmt.Println("Member baru berhasil ditambahkan.")
		}
	}
}

func editMember(){
	var id string
	fmt.Print("masukan ID yang mau diedit: ")
	fmt.Scan(&id)
	index := seqSearchmember(id)
	if index == -1 {
		fmt.Println("ID member tidak ditemukan")
		return
	}
	fmt.Printf("Mengedit Member: %s - %s\n", dataMember[index].id, dataMember[index].nama)
	fmt.Print("Masukkan Nama baru: ")
	fmt.Scan(&dataMember[index].nama)
	fmt.Print("Masukkan Biaya Pendaftaran baru: ")
	fmt.Scan(&dataMember[index].Daftar)
	fmt.Println("Data member berhasil diperbarui.")
}

func hapusMember(){
	var id string
	fmt.Print("Masukan id yang ingin dihapus: ")
	fmt.Scan(&id)
	index := seqSearchmember(id)
	if index == -1{
		fmt.Println("ID member tidak ditemukan")
		return
	}
	for i:=0;i<jumlahMember;i++{
		if dataPinjam[i].idKomik == id && dataPinjam[i].kembalikan == 0 {
			fmt.Println("Gagal: Member masih memiliki pinjaman aktif dan tidak dapat dihapus.")
			return
		}
	}
	for i:= index; i<jumlahMember-1;i++{
		dataMember[i] = dataMember[i+1]
	}
	jumlahMember--
	fmt.Println("Data member berhasil dihapus")
}

func tambahPeminjaman(){
	var n int 
	fmt.Print("Berapa banyak transaksi peminjaman yang ingin ditambahkan? ")
	fmt.Scan(&n)

	if jumlahPinjam+n > NMAX {
		fmt.Printf("Kapasitas tidak cukup. Anda hanya bisa menambahkan %d peminjaman lagi.\n", NMAX-jumlahPinjam)
		return
	}

	for i:=0;i<n;i++{
		fmt.Printf("\n--- Memasukkan data peminjaman ke-%d ---\n", i+1)
		var p peminjaman
		fmt.Print("Masukkan ID Member: ")
		fmt.Scan(&p.idMember)
		if seqSearchmember(p.idMember) == -1{
			fmt.Println("Member tidak ditemukan. Transaksi ini dilewati.")
		} else {
			fmt.Print("Masukkan ID Komik: ")
			fmt.Scan(&p.idKomik)
			if seqSearchkomik(p.idKomik) == -1 {
				fmt.Println("Komik tidak ditemukan. Transaksi ini dilewati.")
			} else {
				fmt.Print("Masukkan Lama Pinjam (hari): ")
				fmt.Scan(&p.lamaPinjam)
				p.kembalikan = 0
				dataPinjam[jumlahPinjam] = p
				jumlahPinjam++
				fmt.Println("Peminjaman berhasil dicatat.")
			}
		}
	}
}

func editPeminjaman(){
	var idK,idM string
	fmt.Print("Masukkan ID Member dari peminjaman yang akan diedit: ")
	fmt.Scan(&idM)
	fmt.Print("Masukkan ID Komik dari peminjaman yang akan diedit: ")
	fmt.Scan(&idK)
	index := cariPeminjaman(idM,idK)
	if index == -1{
		fmt.Println("Data peminjaman tidak ditemukan.")
		return
	}
	fmt.Println("Mengedit peminjaman oleh", idM, "untuk komik", idK)
	fmt.Print("Masukkan Lama Pinjam (hari) baru: ")
	fmt.Scan(&dataPinjam[index].lamaPinjam)
	fmt.Print("Apakah sudah dikembalikan? (1=Ya, 0=Tidak): ")
	var status int
	fmt.Scan(&status)
	if status == 1{
		dataPinjam[index].kembalikan = 1
	} else {
		dataPinjam[index].kembalikan = 0
	}
	fmt.Println("Data peminjaman berhasil diperbarui.")
}

func hapusPeminjaman(){
	var idM, idK string
	fmt.Print("Masukkan ID Member dari peminjaman yang akan dihapus: ")
	fmt.Scan(&idM)
	fmt.Print("Masukkan ID Komik dari peminjaman yang akan dihapus: ")
	fmt.Scan(&idK)
	index := cariPeminjaman(idK,idM)
	if index == -1{
		fmt.Println("Data peminjaman tidak ditemukan.")
		return
	}
	for i := index; i<jumlahPinjam-1;i++{
		dataPinjam[i] = dataPinjam[i+1]
	}
	jumlahPinjam--
	fmt.Println("Data peminjaman berhasil dihapus.")
}

//perhitungan
func hitungDenda(){
	var idM, idK string
	var hariKeterlambatan int
	fmt.Print("Masukkan ID Member: ")
	fmt.Scan(&idM)
	fmt.Print("Masukkan ID Komik: ")
	fmt.Scan(&idK)
	index := cariPeminjaman(idM, idK)
	if index == -1 {
		fmt.Println("Data peminjaman tidak ditemukan.")
		return
	}
	fmt.Print("Masukkan jumlah hari keterlambatan: ")
	fmt.Scan(&hariKeterlambatan)
	if hariKeterlambatan > 0{
		denda := hariKeterlambatan * 1000
		totalUangDenda += denda
		fmt.Printf("Denda yang harus dibayar: Rp%d\n", denda)
	} else {
		fmt.Println("Tidak ada keterlambatan, tidak ada denda.")
	}
	dataPinjam[index].kembalikan = 1
	fmt.Println("Status komik telah diubah menjadi 'sudah kembali'.")
}

func hitungTotalUangMasuk(){
	var totalUangPendaftaran int
	for i:=0;i<jumlahMember;i++{
		totalUangPendaftaran += dataMember[i].Daftar
	}
	total := totalUangPendaftaran + totalUangDenda
	fmt.Println("\n--- Rincian Uang Masuk ---")
	fmt.Printf("Total dari Pendaftaran Member: Rp%d\n", totalUangPendaftaran)
	fmt.Printf("Total dari Denda:              Rp%d\n", totalUangDenda)
	fmt.Printf("Total Uang Masuk Keseluruhan:  Rp%d\n", total)
}

//cetak dan sorting
func tampilkanSemuaKomik() {
    if jumlahKomik == 0 {
        fmt.Println("Belum ada data komik.")
        return
    }
    fmt.Println("\n--- Daftar Semua Komik ---")
    fmt.Println("====================================================================")
    fmt.Printf("| %-8s | %-20s | %-20s | %-4s |\n", "ID", "Judul", "Penulis", "Tahun")
    fmt.Println("====================================================================")
    
    for i := 0; i < jumlahKomik; i++ {
        fmt.Printf("| %-8s | %-20s | %-20s | %-5d |\n", 
            dataKomik[i].id, 
            dataKomik[i].judul, 
            dataKomik[i].penulis, 
            dataKomik[i].pubYear)
    }
    fmt.Println("====================================================================")
    fmt.Printf("Total Komik: %d\n", jumlahKomik)
}

func tampilkanSemuaMember() {
    if jumlahMember == 0 {
        fmt.Println("Belum ada data member.")
        return
    }
    fmt.Println("\n--- Daftar Semua Member ---")
    fmt.Println("=====================================================")
    fmt.Printf("| %-8s | %-20s | %-15s |\n", "ID", "Nama", "Biaya Daftar")
    fmt.Println("=====================================================")
    
    for i := 0; i < jumlahMember; i++ {
        fmt.Printf("| %-8s | %-20s | Rp%-12d |\n", 
            dataMember[i].id, 
            dataMember[i].nama, 
            dataMember[i].Daftar)
    }
    fmt.Println("=====================================================")
    fmt.Printf("Total Member: %d\n", jumlahMember)
}

func tampilkanSemuaPeminjaman() {
    if jumlahPinjam == 0 {
        fmt.Println("Belum ada data peminjaman.")
        return
    }
    fmt.Println("\n--- Daftar Semua Peminjaman ---")
    fmt.Println("====================================================================================")
    fmt.Printf("| %-10s | %-8s | %-15s | %-12s | %-15s |\n", "Member ID", "Komik ID", "Lama Pinjam", "Status", "Denda")
    fmt.Println("====================================================================================")
    
    for i := 0; i < jumlahPinjam; i++ {
        status := "Belum Kembali"
        denda := "-"
        if dataPinjam[i].kembalikan == 1 {
            status = "Sudah Kembali"
        } else {
            denda = "Rp0"
        }
        
        fmt.Printf("| %-10s | %-8s | %-15s | %-12s | %-15s |\n", 
            dataPinjam[i].idMember, 
            dataPinjam[i].idKomik, 
            dataPinjam[i].lamaPinjam, 
            status, 
            denda)
    }
    fmt.Println("====================================================================================")
    fmt.Printf("Total Peminjaman: %d\n", jumlahPinjam)
}

func menuUrutan() {
	var pilihan int
	for pilihan != 4 {
		fmt.Println("\n--- Kelola urutan ---")
		fmt.Println("1. urutan berdasarkan ID")
		fmt.Println("2. Urutan berdasarkan tahun/pubyear")
		fmt.Print("Pilihan: ")
		fmt.Scan(&pilihan)
		if pilihan == 1 {
			urutkanKomikByID()
		} else if pilihan == 2 {
			urutkanKomikByTahun()
		}
	}
}

func binarySearchKomik(id string) int {
	kiri := 0
	kanan := jumlahKomik - 1

	for kiri <= kanan {
		mid := (kiri + kanan) / 2
		if dataKomik[mid].id == id {
			return mid
		} else if dataKomik[mid].id < id {
			kiri = mid + 1
		} else {
			kanan = mid - 1
		}
	}
	return -1 // tidak ditemukan
}


func urutkanKomikByTahun() {
	for i := 1; i < jumlahKomik; i++ {
		key := dataKomik[i]
		j := i - 1
		for j >= 0 && dataKomik[j].pubYear > key.pubYear {
			dataKomik[j+1] = dataKomik[j]
			j = j - 1
		}
		dataKomik[j+1] = key
	}
	fmt.Println("Data komik telah berhasil diurutkan berdasarkan Tahun.")
}

func urutkanKomikByID() {
	//Insertion Sort
	for i := 1; i < jumlahKomik; i++ {
		key := dataKomik[i]
		j := i - 1
		for j >= 0 && dataKomik[j].id > key.id {
			dataKomik[j+1] = dataKomik[j]
			j = j - 1
		}
		dataKomik[j+1] = key
	}
	fmt.Println("Data komik telah berhasil diurutkan berdasarkan ID.")
}




func menuKelolaPeminjaman() {
	var pilihan int
	for pilihan != 5 {
		fmt.Println("\n--- Kelola Peminjaman ---")
		fmt.Println("1. Tambah Peminjaman")
		fmt.Println("2. Edit Peminjaman")
		fmt.Println("3. Hapus Peminjaman")
		fmt.Println("3. Tampilkan Semua Peminjaman")
		fmt.Println("4. Kembali ke Menu Utama")
		fmt.Print("Pilihan: ")
		fmt.Scan(&pilihan)
		if pilihan == 1 {
			tambahPeminjaman()
		} else if pilihan == 2 {
			editPeminjaman()
		} else if pilihan == 3 {
			hapusPeminjaman()
		} else if pilihan == 4 {
			tampilkanSemuaPeminjaman()
		} 
	}
}

func menuCariKomikAtauMember() {
	var pilihan int
	fmt.Println("\nPilih yang ingin dicari:")
	fmt.Println("1. Komik")
	fmt.Println("2. Member")
	fmt.Print("Pilihan: ")
	fmt.Scan(&pilihan)
	if pilihan == 1 {
		var id string
		fmt.Print("Masukkan ID Komik: ")
		fmt.Scan(&id)
		index := seqSearchkomik(id)
		if index != -1 {
			fmt.Println("--- Komik Ditemukan ---")
			fmt.Printf("ID: %s, Judul: %s, Penulis: %s, Tahun: %d\n", dataKomik[index].id, dataKomik[index].judul, dataKomik[index].penulis, dataKomik[index].pubYear)
		} else {
			fmt.Println("Komik dengan ID tersebut tidak ditemukan.")
		}
	} else if pilihan == 2 {
		var id string
		fmt.Print("Masukkan ID Member: ")
		fmt.Scan(&id)
		index := seqSearchmember(id)
		if index != -1 {
			fmt.Println("--- Member Ditemukan ---")
			fmt.Printf("ID: %s, Nama: %s, Bayar Daftar: Rp%d\n", dataMember[index].id, dataMember[index].nama, dataMember[index].Daftar)
		} else {
			fmt.Println("Member dengan ID tersebut tidak ditemukan.")
		}
	} else {
		fmt.Println("Pilihan tidak valid.")
	}
}

func menuTampilkanData() {
	var pilihan int
	for pilihan != 5 {
		fmt.Println("\n--- Tampilkan Data ---")
		fmt.Println("1. Tampilkan Semua Komik")
		fmt.Println("2. Tampilkan Semua Member")
		fmt.Println("3. Tampilkan Semua Peminjaman")
		fmt.Println("4. Tampilkan Semua data pinjaman")
		fmt.Println("5. Kembali ke Menu Utama")
		fmt.Print("Pilihan: ")
		fmt.Scan(&pilihan)
		if pilihan == 1 {
			tampilkanSemuaKomik()
		} else if pilihan == 2 {
			tampilkanSemuaMember()
		} else if pilihan == 3 {
			tampilkanSemuaPeminjaman()
		} else if pilihan == 4{
			tampilkanSemuaKomik()
			tampilkanSemuaMember()
			tampilkanSemuaPeminjaman()
		}
	}
}

