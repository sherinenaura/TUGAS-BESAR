package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"sort"
	"strings"
)

type User struct {
	Username     string
	Password     string
	Profile      string
	Friends      []string
	Bio          string
	TanggalLahir string
}
type Status struct {
	Username string
	Content  string
	Comments []string
}

var users []User
var statuses []Status

func main() {
	loadUsersFromFile()
	loadStatusesFromFile()
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Println("\n=== Aplikasi Sosial Media ===")
		fmt.Println("1.  ✏️ Buat Akun")
		fmt.Println("2.  🔐 Login")
		fmt.Println("3.  ❌ Keluar")
		fmt.Print("Pilih menu: ")
		scanner.Scan()
		choice := strings.TrimSpace(scanner.Text())
		switch choice {
		case "1":
			register()
		case "2":
			login()
		case "3":
			fmt.Println("Terima kasih telah menggunakan aplikasi ini 🙏🏻.")
			saveUsersToFile()
			saveStatusesToFile()
			return
		default:
			fmt.Println("Pilihan tidak valid❗. Silakan coba lagi.")
		}
	}
}
func register() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Masukkan username: ")
	scanner.Scan()
	username := strings.TrimSpace(scanner.Text())

	fmt.Print("Masukkan password: ")
	scanner.Scan()
	password := strings.TrimSpace(scanner.Text())
	for _, user := range users {
		if user.Username == username {
			fmt.Println("Username sudah terdaftar 🤩.")
			return
		}
	}
	newUser := User{
		Username: username,
		Password: password,
		Profile:  "Profil belum diatur",
		Friends:  []string{},
	}
	users = append(users, newUser)
	fmt.Println("Akun telah berhasil di buat,silahkan login kembali🤩")
	saveUsersToFile()
}
func login() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Masukkan username: ")
	scanner.Scan()
	username := strings.TrimSpace(scanner.Text())

	fmt.Print("Masukkan password: ")
	scanner.Scan()
	password := strings.TrimSpace(scanner.Text())
	for _, user := range users {
		if user.Username == username && user.Password == password {
			fmt.Printf("\nSelamat datang🙌🏻, %s!\n", username)
			home(username)
			return
		}
	}
	fmt.Println("Username atau password salah 😖.")
}
func profile(username string) {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Println("\n===  👤Profile  ===")
		fmt.Printf("Nama Pengguna: %s\n", username)
		// Menampilkan Bio (jika ada)
		if bio, found := getUserProfile(username); found && bio != "" {
			fmt.Printf(" 📌 Bio: %s\n", bio)
		} else {
			fmt.Println("Bio: Belum ditambahkan")
		}
		// Menampilkan Tanggal Lahir (jika ada)
		if birthdate, found := getUserBirthdate(username); found && birthdate != "" {
			fmt.Printf(" 📅 Tanggal Lahir: %s\n", birthdate)
		} else {
			fmt.Println("Tanggal Lahir: Belum ditambahkan")
		}
		// Menu pilihan
		fmt.Println("\n1. Tambah Bio")
		fmt.Println("2. Edit Bio")
		fmt.Println("3. Tambah Tanggal Lahir")
		fmt.Println("4. Edit Tanggal Lahir")
		fmt.Println("5. ↩️Kembali ke Menu Utama")
		fmt.Print("Pilih menu: ")
		scanner.Scan()
		choice := strings.TrimSpace(scanner.Text())

		switch choice {
		case "1":
			addBio(username)
		case "2":
			editBio(username)
		case "3":
			addBirthdate(username)
		case "4":
			editBirthdate(username)
		case "5":
			return
		default:
			fmt.Println("Pilihan tidak valid❗.")
		}
	}
}
func addBio(username string) {
	for i, user := range users {
		if user.Username == username {
			if user.Bio != "" {
				fmt.Println("Bio sudah ada, Anda dapat mengeditnya.")
				return
			}
			// Menambahkan Bio jika belum ada
			scanner := bufio.NewScanner(os.Stdin)
			fmt.Print("Masukkan Bio Anda: ")
			scanner.Scan()
			bio := strings.TrimSpace(scanner.Text())
			users[i].Bio = bio
			saveUsersToFile()
			fmt.Println("Bio berhasil ditambahkan!✅")
			return
		}
	}
}
func editBio(username string) {
	for i, user := range users {
		if user.Username == username {
			if user.Bio == "" {
				fmt.Println("Bio belum ditambahkan. Gunakan fitur 'Tambah Bio' terlebih dahulu.")
				return
			}
			// Mengedit Bio yang sudah ada
			scanner := bufio.NewScanner(os.Stdin)
			fmt.Print("Masukkan Bio baru: ")
			scanner.Scan()
			bio := strings.TrimSpace(scanner.Text())
			users[i].Bio = bio
			saveUsersToFile()
			fmt.Println("Bio berhasil diperbarui!✅")
			return
		}
	}
}
func addBirthdate(username string) {
	for i, user := range users {
		if user.Username == username {
			if user.TanggalLahir != "" {
				fmt.Println("Tanggal lahir sudah ada, Anda dapat mengeditnya.")
				return
			}
			// Menambahkan Tanggal Lahir jika belum ada
			scanner := bufio.NewScanner(os.Stdin)
			fmt.Print("Masukkan Tanggal Lahir (format: DD/MM/YYYY): ")
			scanner.Scan()
			birthdate := strings.TrimSpace(scanner.Text())
			users[i].TanggalLahir = birthdate
			saveUsersToFile()
			fmt.Println("Tanggal Lahir berhasil ditambahkan!✅")
			return
		}
	}
}
func editBirthdate(username string) {
	for i, user := range users {
		if user.Username == username {
			if user.TanggalLahir == "" {
				fmt.Println("Tanggal lahir belum ditambahkan. Gunakan fitur 'Tambah Tanggal Lahir' terlebih dahulu.")
				return
			}
			// Mengedit Tanggal Lahir yang sudah ada
			scanner := bufio.NewScanner(os.Stdin)
			fmt.Print("Masukkan Tanggal Lahir baru (format: DD/MM/YYYY): ")
			scanner.Scan()
			birthdate := strings.TrimSpace(scanner.Text())
			users[i].TanggalLahir = birthdate
			saveUsersToFile()
			fmt.Println("Tanggal Lahir berhasil diperbarui!✅")
			return
		}
	}
}

// Helper function untuk mendapatkan Bio pengguna
func getUserProfile(username string) (string, bool) {
	for _, user := range users {
		if user.Username == username {
			return user.Bio, true
		}
	}
	return "", false
}

// Helper function untuk mendapatkan Tanggal Lahir pengguna
func getUserBirthdate(username string) (string, bool) {
	for _, user := range users {
		if user.Username == username {
			return user.TanggalLahir, true
		}
	}
	return "", false
}
func home(username string) {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Println("\n===  🎯 Menu Home  ===")
		fmt.Println("1.  ✍🏻️ Status")
		fmt.Println("2.  🤝 Kelola Teman")
		fmt.Println("3.  👤 Profile")
		fmt.Println("4.  🔎 Cari Pengguna")
		fmt.Println("5.  ❌ Keluar")
		fmt.Print("Pilih menu: ")
		scanner.Scan()
		choice := strings.TrimSpace(scanner.Text())

		switch choice {
		case "1":
			statusMenu(username)
		case "2":
			manageFriends(username)
		case "3":
			profile(username)
		case "4":
			searchUsers(username)
		case "5":
			return
		default:
			fmt.Println("Pilihan tidak valid❗. Silakan coba lagi.")
		}
	}
}

func statusMenu(username string) {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Println("\n=== Menu Status ===")
		fmt.Println("1.  ✍🏻 Tambahkan Status")
		fmt.Println("2.  🔎 Lihat Semua Status")
		fmt.Println("3.  💬 Tambah Komentar")
		fmt.Println("4.  🗑️ Hapus Status")
		fmt.Println("5.  ↩️ Kembali ke Menu Home")
		fmt.Print("Pilih menu: ")
		scanner.Scan()
		choice := strings.TrimSpace(scanner.Text())

		switch choice {
		case "1":
			addStatus(username)
		case "2":
			viewAllStatuses()
		case "3":
			addComment()
		case "4":
			deleteStatus(username)
		case "5":
			return
		default:
			fmt.Println("Pilihan tidak valid ❗. Silakan coba lagi.")
		}
	}
}

func addStatus(username string) {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Masukkan status: ")
	scanner.Scan()
	content := strings.TrimSpace(scanner.Text())
	statuses = append(statuses, Status{Username: username, Content: content})
	saveStatusesToFile()
	fmt.Println("Status berhasil ditambahkan! ✅")
}
func addComment() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Masukkan nomor status yang ingin Anda komentari: ")
	scanner.Scan()
	var statusIndex int
	fmt.Sscanf(scanner.Text(), "%d", &statusIndex)
	if statusIndex < 1 || statusIndex > len(statuses) {
		fmt.Println("Nomor status tidak valid.")
		return
	}

	fmt.Print("Masukkan komentar: ")
	scanner.Scan()
	comment := strings.TrimSpace(scanner.Text())

	// Menambahkan komentar ke status
	statuses[statusIndex-1].Comments = append(statuses[statusIndex-1].Comments, comment)
	saveStatusesToFile()
	fmt.Println("Komentar berhasil ditambahkan! ✅")
}

// Modifikasi fungsi viewAllStatuses untuk menampilkan komentar
func viewAllStatuses() {
	if len(statuses) == 0 {
		fmt.Println("Belum ada status ❓.")
		return
	}
	for i, status := range statuses {
		fmt.Printf("%d. %s: %s\n", i+1, status.Username, status.Content)
		if len(status.Comments) > 0 {
			fmt.Println("   Komentar:")
			for _, comment := range status.Comments {
				fmt.Printf("   - %s\n", comment)
			}
		}
	}
}
func deleteStatus(username string) {
	scanner := bufio.NewScanner(os.Stdin)
	viewAllStatuses()

	fmt.Print("\nMasukkan nomor status yang ingin dihapus: ")
	scanner.Scan()
	var statusIndex int
	fmt.Sscanf(scanner.Text(), "%d", &statusIndex)

	if statusIndex < 1 || statusIndex > len(statuses) {
		fmt.Println("Nomor status tidak valid.")
		return
	}

	if statuses[statusIndex-1].Username != username {
		fmt.Println("Anda hanya dapat menghapus status Anda sendiri.")
		return
	}

	statuses = append(statuses[:statusIndex-1], statuses[statusIndex:]...)
	saveStatusesToFile()
	fmt.Println("Status berhasil dihapus! ✅")
}
func manageFriends(username string) {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Println("\n=== Kelola Teman ===")
		fmt.Println("1.  💕 Tambah Teman")
		fmt.Println("2.  🚮 Hapus Teman")
		fmt.Println("3.  👬 Lihat Daftar Teman")
		fmt.Println("4.  📊 Urutkan Teman")
		fmt.Println("5.  ↩️ Kembali ke Menu Utama")
		fmt.Print("Pilih menu: ")
		scanner.Scan()
		choice := strings.TrimSpace(scanner.Text())

		switch choice {
		case "1":
			fmt.Print("Masukkan username teman yang ingin ditambahkan: ")
			scanner.Scan()
			friend := strings.TrimSpace(scanner.Text())
			for i, user := range users {
				if user.Username == username {
					for _, f := range users[i].Friends {
						if f == friend {
							fmt.Println("Pengguna sudah menjadi teman Anda.✅")
							return
						}
					}
					users[i].Friends = append(users[i].Friends, friend)
					saveUsersToFile()
					fmt.Println("Teman berhasil ditambahkan!✅")
					return
				}
			}
		case "2":
			fmt.Print("Masukkan username teman yang ingin dihapus: ")
			scanner.Scan()
			friend := strings.TrimSpace(scanner.Text())
			for i, user := range users {
				if user.Username == username {
					for j, f := range users[i].Friends {
						if f == friend {
							users[i].Friends = append(users[i].Friends[:j], users[i].Friends[j+1:]...)
							saveUsersToFile()
							fmt.Println("Teman berhasil dihapus!✅")
							return
						}
					}
					fmt.Println("Teman tidak ditemukan❗.")
					return
				}
			}
		case "3":
			for _, user := range users {
				if user.Username == username {
					if len(user.Friends) == 0 {
						fmt.Println("Anda belum memiliki teman.❓")
						return
					}
					fmt.Println("Daftar Teman:")
					for _, f := range user.Friends {
						fmt.Println("-", f)
					}
					return
				}
			}
		case "4":
			sortFriends(username)
		case "5":
			return
		default:
			fmt.Println("Pilihan tidak valid❗.")
		}
	}
}
func sortFriends(username string) {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("\n===  📊Urutkan Teman ===")
	fmt.Println("1. Urutkan berdasarkan Abjad")
	fmt.Println("2. Urutkan berdasarkan Paling Awal Ditambahkan")
	fmt.Println("3. Urutkan berdasarkan Paling Akhir Ditambahkan")
	fmt.Println("4. Kembali ke Menu Kelola Teman")
	fmt.Print("Pilih menu: ")

	scanner.Scan()
	choice := strings.TrimSpace(scanner.Text())

	switch choice {
	case "1":
		sortFriendsByName(username)
	case "2":
		sortFriendsByFirstAdded(username)
	case "3":
		sortFriendsByLastAdded(username)
	case "4":
		return
	default:
		fmt.Println("Pilihan tidak valid❗.")
	}
}
func addFriend(username string) {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Masukkan username teman yang ingin ditambahkan: ")
	scanner.Scan()
	friend := strings.TrimSpace(scanner.Text())
	for i, user := range users {
		if user.Username == username {
			for _, f := range users[i].Friends {
				if f == friend {
					fmt.Println("Pengguna sudah menjadi teman Anda.")
					return
				}
			}
			users[i].Friends = append(users[i].Friends, friend)
			saveUsersToFile()
			fmt.Println("Teman berhasil ditambahkan!✅")
			return
		}
	}
}
func removeFriend(username string) {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Masukkan username teman yang ingin dihapus: ")
	scanner.Scan()
	friend := strings.TrimSpace(scanner.Text())
	for i, user := range users {
		if user.Username == username {
			for j, f := range users[i].Friends {
				if f == friend {
					users[i].Friends = append(users[i].Friends[:j], users[i].Friends[j+1:]...)
					saveUsersToFile()
					fmt.Println("Teman berhasil dihapus!✅")
					return
				}
			}
			fmt.Println("Teman tidak ditemukan❗.")
			return
		}
	}
}
func listFriends(username string) {
	for _, user := range users {
		if user.Username == username {
			if len(user.Friends) == 0 {
				fmt.Println("Anda belum memiliki teman❓.")
				return
			}
			fmt.Println("Daftar Teman:")
			for _, f := range user.Friends {
				fmt.Println("-", f)
			}
			return
		}
	}
}

// Fungsi untuk mengurutkan teman berdasarkan abjad
func sortFriendsByName(username string) {
	for _, user := range users {
		if user.Username == username {
			sort.Strings(user.Friends)
			saveUsersToFile()
			fmt.Println("Teman berhasil diurutkan berdasarkan abjad!✅")
			return
		}
	}
}
func sortFriendsByFirstAdded(username string) {
	fmt.Println("Teman sudah diurutkan berdasarkan urutan penambahan pertama.")
}
func sortFriendsByLastAdded(username string) {
	for _, user := range users {
		if user.Username == username {
			// Membalikkan urutan teman
			for j, k := 0, len(user.Friends)-1; j < k; j, k = j+1, k-1 {
				user.Friends[j], user.Friends[k] = user.Friends[k], user.Friends[j]
			}
			saveUsersToFile()
			fmt.Println("Teman berhasil diurutkan berdasarkan yang terakhir ditambahkan!✅")
			return
		}
	}
}
func editProfile(username string) {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Masukkan profil baru: ")
	scanner.Scan()
	newProfile := strings.TrimSpace(scanner.Text())
	for i, user := range users {
		if user.Username == username {
			users[i].Profile = newProfile
			fmt.Println("Profil berhasil diperbarui!")
			saveUsersToFile()
			return
		}
	}
}
func searchUsers(username string) {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Masukkan username yang ingin dicari: ")
	scanner.Scan()
	searchQuery := strings.TrimSpace(scanner.Text())
	fmt.Println("\nHasil Pencarian:")
	found := false
	for _, user := range users {
		if strings.Contains(user.Username, searchQuery) {
			found = true
			fmt.Printf("- Username: %s, Profil: %s\n", user.Username, user.Profile)
		}
	}
	if !found {
		fmt.Println("Tidak ada pengguna yang ditemukan.")
	}
}
func saveUsersToFile() {
	file, err := os.Create("users.json")
	if err != nil {
		fmt.Println("Gagal menyimpan data pengguna:", err)
		return
	}
	defer file.Close()
	encoder := json.NewEncoder(file)
	err = encoder.Encode(users)
	if err != nil {
		fmt.Println("Gagal encode data pengguna:", err)
	}
}
func loadUsersFromFile() {
	file, err := os.Open("users.json")
	if err != nil {
		if os.IsNotExist(err) {
			return
		}
		fmt.Println("Gagal membaca file pengguna:", err)
		return
	}
	defer file.Close()
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&users)
	if err != nil {
		fmt.Println("Gagal decode data pengguna:", err)
	}
}
func saveStatusesToFile() {
	file, err := os.Create("statuses.json")
	if err != nil {
		fmt.Println("Gagal menyimpan status:", err)
		return
	}
	defer file.Close()
	encoder := json.NewEncoder(file)
	err = encoder.Encode(statuses)
	if err != nil {
		fmt.Println("Gagal encode status:", err)
	}
}
func loadStatusesFromFile() {
	file, err := os.Open("statuses.json")
	if err != nil {
		if os.IsNotExist(err) {
			return
		}
		fmt.Println("Gagal membaca file status:", err)
		return
	}
	defer file.Close()
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&statuses)
	if err != nil {
		fmt.Println("Gagal decode status:", err)
	}
}
