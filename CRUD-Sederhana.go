package main
import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)
var dataFashich []string
func create() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Masukkan Data Baru: ")
	inputDataFashich, _ := reader.ReadString('\n')
	inputDataFashich = strings.TrimSpace(inputDataFashich)
	dataFashich = append(dataFashich, inputDataFashich)
	fmt.Println("Data Berhasil Ditambahkan!")
}
func read() {
	fmt.Println("Data Saat Ini: ")
	for i, data := range dataFashich {
		fmt.Printf("%d: %s\n", i, data)
	}
}
func update() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Masukkan Indeks Data yang ingin diupdate: ")
	indexFashichStr, _ := reader.ReadString('\n')
	indexFashichStr = strings.TrimSpace(indexFashichStr)
	indexFashich, err := strconv.Atoi(indexFashichStr)
	if err != nil || indexFashich < 0 || indexFashich >= len(dataFashich) {
		fmt.Println("Indeks Tidak Valid!, Mohon Coba Lagi!")
		return
	}
	fmt.Print("Masukkan Data Baru: ")
	inputDataFashich, _ := reader.ReadString('\n')
	inputDataFashich = strings.TrimSpace(inputDataFashich)
	dataFashich[indexFashich] = inputDataFashich
	fmt.Println("Data Berhasil Di Update!")
}
func delete() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Masukkan Indeks Data Yang Ingin Dihapus: ")
	indexFashichStr, _ := reader.ReadString('\n')
	indexFashichStr = strings.TrimSpace(indexFashichStr)
	indexFashich, err := strconv.Atoi(indexFashichStr)
	if err != nil || indexFashich < 0 || indexFashich >= len(dataFashich) {
		fmt.Println("Indeks Tidak Valid!, Silakan Coba Lagi")
		return
	}
	dataFashich = append(dataFashich[:indexFashich], dataFashich[indexFashich+1:]...)
	fmt.Println("Data Berhasil Dihapus!")
}
func main() {
	reader := bufio.NewReader(os.Stdin)
	keluar := false
	for !keluar {
		fmt.Println("Pilih Operasi:")
		fmt.Println("1. Create")
		fmt.Println("2. Read")
		fmt.Println("3. Update")
		fmt.Println("4. Delete")
		fmt.Println("5. Keluar")
		fmt.Print("Masukkan Pilihan: ")
		pilihanStr, _ := reader.ReadString('\n')
		pilihanStr = strings.TrimSpace(pilihanStr)
		pilihan, err := strconv.Atoi(pilihanStr)
		if err != nil {
			fmt.Println("Pilihan Anda Tidak Valid!, Mohon Pilih Yang Tersedia")
			continue
		}
		switch pilihan {
		case 1:
			create()
		case 2:
			read()
		case 3:
			update()
		case 4:
			delete()
		case 5:
			keluar = true
		default:
			fmt.Println("Pilihan Anda Tidak Valid!, Mohon Pilih Yang Tersedia")
		}
	}
}