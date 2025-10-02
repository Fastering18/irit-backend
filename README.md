<div align="center">
<h1>IRiT Backend</h1> 
<img src="docs/irits_thumbnail.png" alt="Register Akun Driver" width="280">  

<strong>Final Project No. 2:</strong> IRiT adalah sebuah sistem REST API berbasis backend service yang dirancang untuk mempermudah proses pemesanan Transportasi di ITS (pengujian dapat dilakukan menggunakan metode Postman). Sistem ini dibangun menggunakan Golang, Gin, GORM, SQLite, Postman. IRiT menyediakan layanan fitur autentikasi, booking, tracking lokasi (dummy/mock), serta riwayat booking dan order.
</div>


## @Kelompok 2 BST 
BST265 -
Muhammad Brahmana Priambudi  
BST238 -
Gabriela Asima Nainggolan  
BST087 -
Rizqi Arya Kuskhilbyano  
BST092 -
Ahmad Farras Favian Al Efasi  
BST206 -
Ahmad Zaki Fauzan Nabil  

<details>
  <summary>Cara Menjalankan Backend IRiT</summary>
  <ul>
    <li>
      <strong>1. Prasyarat: Instalasi Go & MySQL</strong><br>
      Pastikan Anda sudah menginstal software berikut:
      <ul>
        <li><a href="https://go.dev/doc/install">Go (versi 1.18 atau lebih baru)</a></li>
        <li><a href="https://git-scm.com/downloads">Git</a> untuk mengambil kode dari repository</li>
      </ul>
      <br><br>
    </li>
    <li>
      <strong>2. Clone Repository dari GitHub</strong><br>
      Buka terminal atau Command Prompt, navigasi folder pilihan, dan jalankan perintah berikut:
      <pre><code>git clone https://github.com/Fastering18/irit-backend.git
cd irit-backend</code></pre>
      <br><br>
    </li>
    <li>
      <strong>3. Konfigurasi Koneksi Database</strong><br>
      Buka file <code>configs/config.yaml</code> lalu edit <code>dsn</code> untuk lokasi penyimpanan file sqlite
      <pre><code>database: "irit.db"</code></pre>
      <br><br>
    </li>
    <li>
      <strong>4. Install Dependencies</strong><br>
      Unduh package Golang yang diperlukan.
      <pre><code>go mod tidy</code></pre>
      <br><br>
    </li>
    <li>
      <strong>5. Jalankan Aplikasi Backend</strong><br>
      Setelah semua persiapan selesai, jalankan server dengan perintah berikut:
      <pre><code>go run cmd/api/main.go</code></pre>
      Jika berhasil, Anda akan melihat output di terminal yang mirip seperti ini dan siap untuk testing.
      <pre><code>2025/10/03 10:00:00 Koneksi database berhasil.
2025/10/03 10:00:00 Migrasi database berhasil.
2025/10/03 10:00:00 Semua routes telah didaftarkan.
2025/10/03 10:00:00 Server berjalan di :8080</code></pre>
      <br>
    </li>
  </ul>
</details>

<details>
  <summary>Test API dengan Postman</summary>
  <ul>
    <li>
      <strong>1. Register Akun User</strong><br>
      <img src="docs/1.%20register_akun_user.png" alt="Register Akun User" width="600">
      <br><br>
    </li>
    <li>
      <strong>2. Generate Token JWT User</strong><br>
      <img src="docs/2.%20generate_token_jwt_user.png" alt="Generate Token JWT User" width="600">
      <br><br>
    </li>
    <li>
      <strong>3. Akses User Dengan JWT</strong><br>
      <img src="docs/3.%20akses_user_dengan_jwt.png" alt="Akses User Dengan JWT" width="600">
      <br><br>
    </li>
    <li>
      <strong>4. Register Akun Driver</strong><br>
      <img src="docs/4.%20register_akun_driver.png" alt="Register Akun Driver" width="600">
      <br><br>
    </li>
    <li>
      <strong>5. Generate Token JWT Driver</strong><br>
      <img src="docs/5.%20generate_token_jwt_driver.png" alt="Generate Token JWT Driver" width="600">
      <br><br>
    </li>
    <li>
      <strong>6. Akses Driver Dengan JWT</strong><br>
      <img src="docs/6.%20akses_driver_dengan_jwt.png" alt="Akses Driver Dengan JWT" width="600">
      <br><br>
    </li>
    <li>
      <strong>7. Book User ke Driver</strong><br>
      <img src="docs/7.%20book_user2driver.png" alt="Book User ke Driver" width="600">
      <br><br>
    </li>
    <li>
      <strong>8. Driver Cek Daftar Order</strong><br>
      <img src="docs/8.%20driver_cek_daftar_order.png" alt="Driver Cek Daftar Order" width="600">
      <br><br>
    </li>
    <li>
      <strong>9. Driver Accept Order</strong><br>
      <img src="docs/9.%20driver_accept_order.png" alt="Driver Accept Order" width="600">
      <br><br>
    </li>
    <li>
      <strong>10. User Cek Booking History</strong><br>
      <img src="docs/10.%20user_cek_booking_history.png" alt="User Cek Booking History" width="600">
      <br><br>
    </li>
    <li>
      <strong>11. Driver Set Booking Status</strong><br>
      <img src="docs/11.%20driver_set_booking_status.png" alt="Driver Set Booking Status" width="600">
      <br><br>
    </li>
    <li>
      <strong>12. User Cek Jarak (Mock Up)</strong><br>
      <img src="docs/12.%20user_cek_jarak%20(mock%20up).png" alt="User Cek Jarak (Mock Up)" width="600">
      <br><br>
    </li>
  </ul>
</details>

<br />  

---
<p align="center">
  Dibuat sebagai Final Project untuk <strong>BST Advanced 2025</strong>
  <br>
  <small>© 2025 Kelompok 2 BST Advanced 2025. All Rights Reserved.</small>
</p>