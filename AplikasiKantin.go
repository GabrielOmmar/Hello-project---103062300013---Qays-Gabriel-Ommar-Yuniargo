package main

import (
	"fmt"
	"os"
	"sort"
)

// Struktur untuk menyimpan informasi tenant
type Tenant struct {
	Nama            string
	TotalTransaksi  float64
	JumlahTransaksi int
}

// Struktur untuk menyimpan transaksi
type Transaksi struct {
	NamaTenant string
	Jumlah     float64
}

// Data tenant dan transaksi disimpan dalam slice
var tenants []Tenant
var transaksi []Transaksi

// Prosedur untuk menambahkan tenant
func tambahTenant(nama string) {
	tenant := Tenant{Nama: nama}
	tenants = append(tenants, tenant)
}

// Prosedur untuk mengubah data tenant
func ubahTenant(namaLama string, namaBaru string) {
	for i := range tenants {
		if tenants[i].Nama == namaLama {
			tenants[i].Nama = namaBaru
			break
		}
	}
}

// Prosedur untuk menghapus tenant
func hapusTenant(nama string) {
	for i := range tenants {
		if tenants[i].Nama == nama {
			tenants = append(tenants[:i], tenants[i+1:]...)
			break
		}
	}
}

// Prosedur untuk mencatat transaksi
func tambahTransaksi(namaTenant string, jumlah float64) {
	transaksiBaru := Transaksi{NamaTenant: namaTenant, Jumlah: jumlah}
	transaksi = append(transaksi, transaksiBaru)
	for i := range tenants {
		if tenants[i].Nama == namaTenant {
			tenants[i].TotalTransaksi += jumlah
			tenants[i].JumlahTransaksi++
			break
		}
	}
}

// Fungsi untuk menghitung jumlah uang yang diperoleh tenant dan admin kantin
func hitungPendapatan() ([]float64, float64) {
	pendapatanTenant := make([]float64, len(tenants))
	var pendapatanAdmin float64
	for _, t := range transaksi {
		bagianTenant := t.Jumlah * 0.75
		bagianAdmin := t.Jumlah * 0.25
		for i, tenant := range tenants {
			if tenant.Nama == t.NamaTenant {
				pendapatanTenant[i] += bagianTenant
				break
			}
		}
		pendapatanAdmin += bagianAdmin
	}
	return pendapatanTenant, pendapatanAdmin
}

// Prosedur untuk menampilkan daftar tenant secara terurut berdasarkan banyak transaksi
func daftarTenantBerdasarkanTransaksi() {
	sort.SliceStable(tenants, func(i, j int) bool {
		return tenants[i].JumlahTransaksi > tenants[j].JumlahTransaksi
	})

	file, _ := os.Create("daftar_tenant.txt")
	defer file.Close()

	fmt.Println("Daftar Tenant berdasarkan banyak transaksi:")
	for _, tenant := range tenants {
		output := fmt.Sprintf("Nama: %s, Jumlah Transaksi: %d, Total Uang: %.2f\n", tenant.Nama, tenant.JumlahTransaksi, tenant.TotalTransaksi)
		file.WriteString(output)
		fmt.Print(output)
	}
	fmt.Println("Daftar tenant berhasil ditulis ke daftar_tenant.txt")
}

// Prosedur untuk menampilkan pendapatan tenant dan admin ke dalam file
func tampilkanPendapatanKeFile() {
	pendapatanTenant, pendapatanAdmin := hitungPendapatan()
	file, _ := os.Create("pendapatan.txt")
	defer file.Close()

	fmt.Println("Pendapatan Tenant:")
	for i, tenant := range tenants {
		output := fmt.Sprintf("Tenant Nama: %s, Pendapatan: %.2f\n", tenant.Nama, pendapatanTenant[i])
		file.WriteString(output)
		fmt.Print(output)
	}
	outputAdmin := fmt.Sprintf("Pendapatan Admin: %.2f\n", pendapatanAdmin)
	file.WriteString(outputAdmin)
	fmt.Print(outputAdmin)

	fmt.Println("Pendapatan berhasil ditulis ke pendapatan.txt")
}

func main() {
	var pilihan int
	for {
		fmt.Println(`
===================================
|         Menu                    |
| 1. Tambah Tenant                |
| 2. Ubah Tenant                  |
| 3. Hapus Tenant                 |
| 4. Tambah Transaksi             |
| 5. Tampilkan Pendapatan         |
| 6. Tampilkan Daftar Tenant      |
|    Berdasarkan Banyak Transaksi |
| 7. Keluar                       |
===================================
		`)
		fmt.Print("Pilih opsi: ")
		fmt.Scan(&pilihan)

		if pilihan == 1 {
			var nama string
			fmt.Print("Masukkan Nama Tenant: ")
			fmt.Scan(&nama)
			tambahTenant(nama)
			fmt.Println("Tenant berhasil ditambahkan.")
		} else if pilihan == 2 {
			var namaLama, namaBaru string
			fmt.Print("Masukkan Nama Tenant yang ingin diubah: ")
			fmt.Scan(&namaLama)
			fmt.Print("Masukkan Nama Baru: ")
			fmt.Scan(&namaBaru)
			ubahTenant(namaLama, namaBaru)
			fmt.Println("Tenant berhasil diubah.")
		} else if pilihan == 3 {
			var nama string
			fmt.Print("Masukkan Nama Tenant yang ingin dihapus: ")
			fmt.Scan(&nama)
			hapusTenant(nama)
			fmt.Println("Tenant berhasil dihapus.")
		} else if pilihan == 4 {
			var namaTenant string
			var jumlah float64
			fmt.Print("Masukkan Nama Tenant: ")
			fmt.Scan(&namaTenant)
			fmt.Print("Masukkan Jumlah Transaksi: ")
			fmt.Scan(&jumlah)
			tambahTransaksi(namaTenant, jumlah)
			fmt.Println("Transaksi berhasil dicatat.")
		} else if pilihan == 5 {
			tampilkanPendapatanKeFile()
		} else if pilihan == 6 {
			daftarTenantBerdasarkanTransaksi()
		} else if pilihan == 7 {
			fmt.Println("Keluar dari program.")
			return
		} else {
			fmt.Println("Opsi tidak valid.")
		}
	}
}
