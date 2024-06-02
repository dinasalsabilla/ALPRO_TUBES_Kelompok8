// Aplikasi Pendaftaran Mahasiswa
// Deskripsi: Aplikasi pendaftaran mahasiswa di perguruan tinggi ini
// mengelola data calon mahasiswa dan data jurusan perguruan tinggi.
// Pengguna aplikasi ini adalah petugas admin aplikasi dan juga calon mahasiswa.
// Spesifikasi:
// a. Pengguna bisa melakukan penambahan, pengubahan (edit), dan penghapusan data mahasiswa dan data jurusan di suatu perguruan tinggi.
// b. Pengguna bisa menampilkan data mahasiswa yang mendaftar pada jurusan tertentu. data mahasiswa yang telah diterima, dan juga ditolak.
// c. Pengguna bisa melakukan penambahan, pengubahan (edit), dan penghapusan nilai tes mahasiswa, yang nantinya akan menentukan mahasiswa diterima atau ditolak.
// d. Pengguna bisa menampilkan data mahasiswa terurut berdasarkan nilai test, jurusan, nama mahasiswa.

package main

import "fmt"

const nMax int = 1000

type mahasiswa struct {
	nama         string
	email        string
	programStudi string
	nilaiTes     int
}

type tabMhs [nMax]mahasiswa

var data tabMhs
var nData int

func main() {
	dataMhs()
	menuPengguna()
}

func dataMhs() {
	data[0] = mahasiswa{"Upin", "upin@gmail.com", "IF", 60}
	data[1] = mahasiswa{"Ismail", "ismail@gmail.com", "SI", 83}
	data[2] = mahasiswa{"Fizi", "fizi@gmail.com", "RPL", 56}
	data[3] = mahasiswa{"Ipin", "ipin@gmail.com", "IF", 89}
	data[4] = mahasiswa{"Meimei", "meimei@gmail.com", "SI", 64}
	data[5] = mahasiswa{"Ehsan", "ehsan@gmail.com", "RPL", 77}
	nData = 6
}

func menuPengguna() {
	var pilihan int

	for pilihan != 3 {
		fmt.Println("----- APLIKASI PENDAFTARAN MAHASISWA -----")
		fmt.Println("Masuk sebagai :						   ")
		fmt.Println("1. Admin aplikasi				    	   ")
		fmt.Println("2. Calon mahasiswa						   ")
		fmt.Println("3. Keluar")
		fmt.Println("------------------------------------------")
		fmt.Print("Pilihan: ")
		fmt.Scan(&pilihan)
		if pilihan == 1 {
			menuAdmin()
		} else if pilihan == 2 {
			menuMhs()
		} else if pilihan == 3 {
			fmt.Println("Terima kasih telah menggunakan aplikasi ini!")
		} else {
			fmt.Println("Pilihan tidak valid.")
		}
	}
}

func menuAdmin() {
	var subPilihan int

	for subPilihan != 13 {
		fmt.Println("---------------- APLIKASI PENDAFTARAN MAHASISWA ----------------")
		fmt.Println("Menu Admin :")
		fmt.Println("1. Tambah nilai tes mahasiswa")
		fmt.Println("2. Edit nilai tes mahasiswa")
		fmt.Println("3. Hapus nilai tes mahasiswa")
		fmt.Println("4. Data mahasiswa yang mendaftar berdasarkan jurusan")
		fmt.Println("5. Data mahasiswa yang diterima berdasarkan jurusan ")
		fmt.Println("6. Data mahasiswa yang ditolak berdasarkan jurusan")
		fmt.Println("7. Ascending berdasarkan nama")
		fmt.Println("8. Ascending berdasarkan jurusan")
		fmt.Println("9. Ascending berdasarkan nilai tes")
		fmt.Println("10. Descending berdasarkan nama")
		fmt.Println("11. Descending berdasarkan jurusan")
		fmt.Println("12. Descending berdasarkan nilai tes")
		fmt.Println("13. Keluar")
		fmt.Println("----------------------------------------------------------------")
		fmt.Print("Pilihan: ")
		fmt.Scan(&subPilihan)
		if subPilihan == 1 {
			tambahNilaiTes(&data, nData)
		} else if subPilihan == 2 {
			editNilaiTes(&data, nData)
		} else if subPilihan == 3 {
			hapusNilaiTes(&data, nData)
		} else if subPilihan == 4 {
			dataMendaftar(data, nData)
		} else if subPilihan == 5 {
			dataDiterima(data, nData)
		} else if subPilihan == 6 {
			dataDitolak(data, nData)
		} else if subPilihan == 7 {
			ascendingNama(data, nData)
		} else if subPilihan == 8 {
			ascendingProgramStudi(data, nData)
		} else if subPilihan == 9 {
			ascendingNilaiTes(data, nData)
		} else if subPilihan == 10 {
			descendingNama(data, nData)
		} else if subPilihan == 11 {
			descendingProgramStudi(data, nData)
		} else if subPilihan == 12 {
			descendingNilaiTes(data, nData)
		} else if subPilihan == 13 {
			subPilihan = 13
		} else {
			fmt.Println("Pilihan tidak valid.")
		}
	}
}

func tambahNilaiTes(A *tabMhs, n int) { // dapat menambah nilaiTes mahasiswa berdasarkan pencarian nama dan nantinya menjadi penentu apakah diterima atau ditolak
	var namaMhs string
	var nilaiTesMhs int

	fmt.Print("Masukkan nama mahasiswa: ")
	fmt.Scan(&namaMhs)

	idx := sequentialSearchNama(*A, n, namaMhs)
	if idx == -1 {
		fmt.Println("Mahasiswa dengan nama", namaMhs, "tidak ditemukan!")
	} else {
		fmt.Println("Mahasiswa dengan nama", namaMhs, "ditemukan. Silakan masukkan nilai tes!")
		fmt.Print("Masukkan nilai tes: ")
		fmt.Scan(&nilaiTesMhs)
		A[idx].nilaiTes = nilaiTesMhs
		fmt.Println("Nilai tes mahasiswa berhasil ditambahkan.")
	}
}

func editNilaiTes(A *tabMhs, n int) { // dapat menghapus nilaiTes mahasiswa berdasarkan pencarian nama dan nantinya menjadi penentu apakah diterima atau ditolak
	var namaMhs string

	fmt.Print("Masukkan nama mahasiswa yang akan diedit: ")
	fmt.Scan(&namaMhs)

	idx := sequentialSearchNama(*A, n, namaMhs)
	if idx == -1 {
		fmt.Println("Mahasiswa dengan nama", namaMhs, "tidak ditemukan!")
	} else {
		if A[idx].nilaiTes == 0 {
			fmt.Println("Mohon maaf, nilai tes untuk mahasiswa ini belum diinput.")
		} else {
			fmt.Println("Nilai tes mahasiswa dengan nama", namaMhs, "ditemukan. Silakan masukkan perubahan!")
			fmt.Print("Nilai tes: ")
			fmt.Scan(&A[idx].nilaiTes)
			fmt.Println("Nilai tes mahasiswa berhasil diubah!")
		}
	}
}

func hapusNilaiTes(A *tabMhs, n int) { // dapat menghapus nilaiTes mahasiswa berdasarkan pencarian nama menggunakan binary search dan nantinya menjadi penentu apakah diterima atau ditolak
	var namaMhs string
	var temp mahasiswa

	fmt.Print("Masukkan nama mahasiswa yang akan dihapus: ")
	fmt.Scan(&namaMhs)

	// selection sort
	for i := 0; i < n-1; i++ {
		idx := i
		for j := i + 1; j < n; j++ {
			if (*A)[j].nama < (*A)[idx].nama {
				idx = j
			}
		}
		temp = (*A)[idx]
		(*A)[idx] = (*A)[i]
		(*A)[i] = temp
	}

	idx := binarySearch(*A, n, namaMhs)
	if idx == -1 {
		fmt.Println("Nilai tes mahasiswa dengan nama", namaMhs, "tidak ditemukan!")
	} else {
		fmt.Println("Nilai tes mahasiswa dengan nama", namaMhs, "ditemukan dan akan dihapus.")
		(*A)[idx].nilaiTes = 0
		fmt.Println("Nilai tes mahasiswa dengan nama", namaMhs, "berhasil dihapus!")
	}
}

func dataMendaftar(A tabMhs, n int) { // dapat menampilkan semua data mahasiswa yang mendaftar pada suatu programStudi
	var proDi string

	fmt.Print("Masukkan program studi yang akan ditampilkan data pendaftarnya: ")
	fmt.Scan(&proDi)

	fmt.Println("Data mahasiswa yang mendaftar pada program studi", proDi, ":")

	// sequential search
	found := false
	for i := 0; i < n; i++ {
		if A[i].programStudi == proDi {
			fmt.Println("-------------------------------")
			fmt.Println("Nama: ", A[i].nama)
			fmt.Println("Email: ", A[i].email)
			fmt.Println("Nilai Tes: ", A[i].nilaiTes)
			if A[i].nilaiTes == 0 {
				fmt.Println("Status Seleksi: Sedang diproses")
			} else if A[i].nilaiTes >= 65 {
				fmt.Println("Status Seleksi: Diterima")
			} else {
				fmt.Println("Status Seleksi: Ditolak")
			}
			found = true
		}
	}
	if !found {
		fmt.Println("Tidak ada mahasiswa yang mendaftar pada program studi", proDi)
	}
}

func dataDiterima(A tabMhs, n int) { // dapat menampilkan semua data mahasiswa yang diterima pada suatu programStudi
	var proDi string

	fmt.Print("Masukkan program studi yang akan ditampilkan data mahasiswanya yang diterima: ")
	fmt.Scan(&proDi)

	fmt.Println("Data mahasiswa yang diterima pada program studi", proDi, ":")

	// sequential search
	found := false
	for i := 0; i < n; i++ {
		if A[i].programStudi == proDi && A[i].nilaiTes >= 65 {
			fmt.Println("-------------------------------")
			fmt.Println("Nama: ", A[i].nama)
			fmt.Println("Email: ", A[i].email)
			fmt.Println("Nilai Tes: ", A[i].nilaiTes)
			fmt.Println("Status Seleksi: Diterima")
			found = true
		}
	}
	if !found {
		fmt.Println("Tidak ada mahasiswa yang diterima pada program studi", proDi)
	}
}

func dataDitolak(A tabMhs, n int) { // dapat menampilkan semua data mahasiswa yang ditolak pada suatu programStudi
	var proDi string

	fmt.Print("Masukkan program studi yang akan ditampilkan data mahasiswanya yang ditolak: ")
	fmt.Scan(&proDi)

	fmt.Println("Data mahasiswa yang ditolak pada program studi", proDi, ":")

	// sequential search
	found := false
	for i := 0; i < n; i++ {
		if A[i].programStudi == proDi && A[i].nilaiTes > 0 && A[i].nilaiTes < 65 {
			fmt.Println("-------------------------------")
			fmt.Println("Nama: ", A[i].nama)
			fmt.Println("Email: ", A[i].email)
			fmt.Println("Nilai Tes: ", A[i].nilaiTes)
			fmt.Println("Status Seleksi: Ditolak")
			found = true
		}
	}
	if !found {
		fmt.Println("Tidak ada mahasiswa yang ditolak pada program studi", proDi)
	}
}

func ascendingNama(A tabMhs, n int) { // dapat mengurutkan semua data mahasiswa ascending berdasarkan nama
	var temp mahasiswa

	// selection sort
	for i := 0; i < n-1; i++ {
		idx := i
		for j := i + 1; j < n; j++ {
			if A[j].nama < A[idx].nama {
				idx = j
			}
		}
		temp = A[idx]
		A[idx] = A[i]
		A[i] = temp
	}

	fmt.Println("Data mahasiswa setelah diurutkan ascending berdasarkan nama:")
	for i := 0; i < n; i++ {
		fmt.Println("-------------------------------")
		fmt.Println("Nama: ", A[i].nama)
		fmt.Println("Email: ", A[i].email)
		fmt.Println("Program Studi: ", A[i].programStudi)
		fmt.Println("Nilai Tes: ", A[i].nilaiTes)
		if A[i].nilaiTes == 0 {
			fmt.Println("Status Seleksi: Sedang diproses")
		} else if A[i].nilaiTes >= 65 {
			fmt.Println("Status Seleksi: Diterima")
		} else {
			fmt.Println("Status Seleksi: Ditolak")
		}
	}
}

func ascendingProgramStudi(A tabMhs, n int) { // dapat mengurutkan semua data mahasiswa ascending berdasarkan programStudi
	var temp mahasiswa

	// selection sort
	for i := 0; i < n-1; i++ {
		idx := i
		for j := i + 1; j < n; j++ {
			if A[j].programStudi < A[idx].programStudi {
				idx = j
			}
		}
		temp = A[idx]
		A[idx] = A[i]
		A[i] = temp
	}

	fmt.Println("Data mahasiswa setelah diurutkan ascending berdasarkan program studi:")
	for i := 0; i < n; i++ {
		fmt.Println("-------------------------------")
		fmt.Println("Nama: ", A[i].nama)
		fmt.Println("Email: ", A[i].email)
		fmt.Println("Program Studi: ", A[i].programStudi)
		fmt.Println("Nilai Tes: ", A[i].nilaiTes)
		if A[i].nilaiTes == 0 {
			fmt.Println("Status Seleksi: Sedang diproses")
		} else if A[i].nilaiTes >= 65 {
			fmt.Println("Status Seleksi: Diterima")
		} else {
			fmt.Println("Status Seleksi: Ditolak")
		}
	}
}

func ascendingNilaiTes(A tabMhs, n int) { // dapat mengurutkan semua data mahasiswa ascending berdasarkan nilaiTes
	var temp mahasiswa

	// selection sort
	for i := 0; i < n-1; i++ {
		idx := i
		for j := i + 1; j < n; j++ {
			if A[j].nilaiTes < A[idx].nilaiTes {
				idx = j
			}
		}
		temp = A[idx]
		A[idx] = A[i]
		A[i] = temp
	}

	fmt.Println("Data mahasiswa setelah diurutkan ascending berdasarkan nilai tes:")
	for i := 0; i < n; i++ {
		fmt.Println("-------------------------------")
		fmt.Println("Nama: ", A[i].nama)
		fmt.Println("Email: ", A[i].email)
		fmt.Println("Program Studi: ", A[i].programStudi)
		fmt.Println("Nilai Tes: ", A[i].nilaiTes)
		if A[i].nilaiTes == 0 {
			fmt.Println("Status Seleksi: Sedang diproses")
		} else if A[i].nilaiTes >= 65 {
			fmt.Println("Status Seleksi: Diterima")
		} else {
			fmt.Println("Status Seleksi: Ditolak")
		}
	}
}

func descendingNama(A tabMhs, n int) { // dapat mengurutkan semua data mahasiswa descending berdasarkan nama
	var temp mahasiswa

	// insertion sort
	for i := 1; i < n; i++ {
		temp = A[i]
		j := i - 1
		for j >= 0 && temp.nama > A[j].nama {
			A[j+1] = A[j]
			j--
		}
		A[j+1] = temp
	}

	fmt.Println("Data mahasiswa setelah diurutkan descending berdasarkan nama:")
	for i := 0; i < n; i++ {
		fmt.Println("-------------------------------")
		fmt.Println("Nama: ", A[i].nama)
		fmt.Println("Email: ", A[i].email)
		fmt.Println("Program Studi: ", A[i].programStudi)
		fmt.Println("Nilai Tes: ", A[i].nilaiTes)
		if A[i].nilaiTes == 0 {
			fmt.Println("Status Seleksi: Sedang diproses")
		} else if A[i].nilaiTes >= 65 {
			fmt.Println("Status Seleksi: Diterima")
		} else {
			fmt.Println("Status Seleksi: Ditolak")
		}
	}
}

func descendingProgramStudi(A tabMhs, n int) { // dapat mengurutkan semua data mahasiswa descending berdasarkan programStudi
	var temp mahasiswa

	// insertion sort
	for i := 1; i < n; i++ {
		temp = A[i]
		j := i - 1
		for j >= 0 && temp.programStudi > A[j].programStudi {
			A[j+1] = A[j]
			j--
		}
		A[j+1] = temp
	}

	fmt.Println("Data mahasiswa setelah diurutkan descending berdasarkan program studi:")
	for i := 0; i < n; i++ {
		fmt.Println("-------------------------------")
		fmt.Println("Nama: ", A[i].nama)
		fmt.Println("Email: ", A[i].email)
		fmt.Println("Program Studi: ", A[i].programStudi)
		fmt.Println("Nilai Tes: ", A[i].nilaiTes)
		if A[i].nilaiTes == 0 {
			fmt.Println("Status Seleksi: Sedang diproses")
		} else if A[i].nilaiTes >= 65 {
			fmt.Println("Status Seleksi: Diterima")
		} else {
			fmt.Println("Status Seleksi: Ditolak")
		}
	}
}

func descendingNilaiTes(A tabMhs, n int) { // dapat mengurutkan semua data mahasiswa descending berdasarkan nilaiTes
	var temp mahasiswa

	// insertion sort
	for i := 1; i < n; i++ {
		temp = A[i]
		j := i - 1
		for j >= 0 && temp.nilaiTes > A[j].nilaiTes {
			A[j+1] = A[j]
			j--
		}
		A[j+1] = temp
	}

	fmt.Println("Data mahasiswa setelah diurutkan descending berdasarkan nilai tes:")
	for i := 0; i < n; i++ {
		fmt.Println("-------------------------------")
		fmt.Println("Nama: ", A[i].nama)
		fmt.Println("Email: ", A[i].email)
		fmt.Println("Program Studi: ", A[i].programStudi)
		fmt.Println("Nilai Tes: ", A[i].nilaiTes)
		if A[i].nilaiTes == 0 {
			fmt.Println("Status Seleksi: Sedang diproses")
		} else if A[i].nilaiTes >= 65 {
			fmt.Println("Status Seleksi: Diterima")
		} else {
			fmt.Println("Status Seleksi: Ditolak")
		}
	}
}

func menuMhs() {
	var subPilihan int

	for subPilihan != 6 {
		fmt.Println("---------------- APLIKASI PENDAFTARAN MAHASISWA ----------------")
		fmt.Println("Menu Calon Mahasiswa :")
		fmt.Println("1. Tambah data mahasiswa")
		fmt.Println("2. Edit data mahasiswa")
		fmt.Println("3. Hapus data mahasiswa")
		fmt.Println("4. Tampil data mahasiswa")
		fmt.Println("5. Status seleksi mahasiswa")
		fmt.Println("6. Keluar")
		fmt.Println("----------------------------------------------------------------")
		fmt.Print("Pilihan: ")
		fmt.Scan(&subPilihan)
		if subPilihan == 1 {
			tambahData(&data, &nData)
		} else if subPilihan == 2 {
			editData(&data, nData)
		} else if subPilihan == 3 {
			hapusData(&data, &nData)
		} else if subPilihan == 4 {
			tampilData(data, nData)
		} else if subPilihan == 5 {
			statusSeleksi(data, nData)
		} else if subPilihan == 6 {
			subPilihan = 6
		} else {
			fmt.Println("Pilihan tidak valid.")
		}
	}
}

func tambahData(A *tabMhs, n *int) { // dapat menambah data mahasiswa (nama, email, programStudi)

	if *n >= nMax {
		fmt.Println("Kapasitas maksimum tercapai!")
	} else {
		fmt.Println("Masukkan data mahasiswa!")
		fmt.Print("Nama: ")
		fmt.Scan(&A[*n].nama)
		fmt.Print("Email: ")
		fmt.Scan(&A[*n].email)
		fmt.Print("Program studi yang dipilih: ")
		fmt.Scan(&A[*n].programStudi)

		fmt.Println("Data mahasiswa dengan nama", A[*n].nama, "berhasil ditambahkan!")
		*n++
	}
}

func editData(A *tabMhs, n int) { // dapat mengedit data mahasiswa (nama, email, programStudi) berdasarkan pencarian nama
	var namaMhs string

	fmt.Print("Masukkan nama mahasiswa yang akan diedit: ")
	fmt.Scan(&namaMhs)

	idx := sequentialSearchNama(*A, n, namaMhs)
	if idx == -1 {
		fmt.Println("Data mahasiswa dengan nama", namaMhs, "tidak ditemukan!")
	} else {
		fmt.Println("Data mahasiswa dengan nama", namaMhs, "ditemukan. Silakan masukkan perubahan!")
		fmt.Print("Nama: ")
		fmt.Scan(&A[idx].nama)
		fmt.Print("Email: ")
		fmt.Scan(&A[idx].email)
		fmt.Print("Program studi yang dipilih: ")
		fmt.Scan(&A[idx].programStudi)
		fmt.Println("Data mahasiswa berhasil diubah!")
	}
}

func hapusData(A *tabMhs, n *int) { // dapat menghapus data mahasiswa (nama, email, programStudi) berdasarkan pencarian nama menggunakan binary search
	var namaMhs string
	var temp mahasiswa

	fmt.Print("Masukkan nama mahasiswa yang akan dihapus: ")
	fmt.Scan(&namaMhs)

	// selection sort
	for i := 0; i < *n-1; i++ {
		idx := i
		for j := i + 1; j < *n; j++ {
			if A[j].nama < A[idx].nama {
				idx = j
			}
		}
		temp = A[idx]
		A[idx] = A[i]
		A[i] = temp
	}

	idx := binarySearch(*A, *n, namaMhs)
	if idx == -1 {
		fmt.Println("Data mahasiswa dengan nama", namaMhs, "tidak ditemukan!")
	} else {
		fmt.Println("Data mahasiswa dengan nama", namaMhs, "ditemukan dan akan dihapus.")
		for i := idx; i < *n-1; i++ {
			(*A)[i] = (*A)[i+1]
		}
		*n--
		fmt.Println("Data mahasiswa dengan nama", namaMhs, "berhasil dihapus!")
	}
}

func tampilData(A tabMhs, n int) { // dapat menampilkan data mahasiswa (nama, email, programStudi) berdasarkan pencarian nama
	var namaMhs string

	fmt.Print("Masukkan nama mahasiswa yang akan ditampilkan datanya: ")
	fmt.Scan(&namaMhs)

	idx := sequentialSearchNama(A, n, namaMhs)
	if idx == -1 {
		fmt.Println("Data mahasiswa dengan nama", namaMhs, "tidak ditemukan!")
	} else {
		fmt.Println("Data mahasiswa dengan nama", namaMhs, "ditemukan!")
		fmt.Println("Nama: ", A[idx].nama)
		fmt.Println("Email: ", A[idx].email)
		fmt.Println("Program Studi: ", A[idx].programStudi)
	}
}

func statusSeleksi(A tabMhs, n int) { // dapat menampilkan sedang diproses bila admin belum melakukan tambahNilaiTes() atau diterima bila admin telah melakukan tambahNilaiTes() dan nilai lebih dari sama dengan 65 atau ditolak bila admin telah melakukan tambahNilaiTes() dan nilai kurang dari 65
	var namaMhs string

	fmt.Print("Masukkan nama mahasiswa yang akan ditampilkan status seleksinya: ")
	fmt.Scan(&namaMhs)

	idx := sequentialSearchNama(A, n, namaMhs)
	if idx == -1 {
		fmt.Println("Data mahasiswa dengan nama", namaMhs, "tidak ditemukan!")
	} else {
		fmt.Println("Data mahasiswa dengan nama", namaMhs, "ditemukan!")
		fmt.Println("Nama: ", A[idx].nama)
		fmt.Println("Email: ", A[idx].email)
		fmt.Println("Program Studi: ", A[idx].programStudi)
		if A[idx].nilaiTes == 0 {
			fmt.Println("Status Seleksi: Sedang diproses")
			fmt.Println("Silahkan dicek secara berkala.")
		} else if A[idx].nilaiTes >= 65 {
			fmt.Println("Nilai tes:", A[idx].nilaiTes)
			fmt.Println("Status Seleksi: Diterima")
			fmt.Println("Selamat! Manfaatkan setiap kesempatan untuk belajar dan berkembang.")
		} else {
			fmt.Println("Nilai tes:", A[idx].nilaiTes)
			fmt.Println("Status seleksi: Ditolak")
			fmt.Println("Mohon maaf! Jangan putus asa dan tetap semangat!")
		}
	}
}

func sequentialSearchNama(A tabMhs, n int, namaMhs string) int {
	i := 0
	idx := -1
	for i < n && idx == -1 {
		if A[i].nama == namaMhs {
			idx = i
		}
		i++
	}
	return idx
}

func binarySearch(A tabMhs, n int, namaMhs string) int {
	left := 0
	right := n - 1
	idx := -1
	for left <= right && idx == -1 {
		med := (left + right) / 2
		if A[med].nama < namaMhs {
			left = med + 1
		} else if A[med].nama > namaMhs {
			right = med - 1
		} else {
			idx = med
		}
	}
	return idx
}
