# Belajar Dengan Jenius Golang

## Penulis : Gun Gun Febrianza

## *Data Hierarchy*

Data yang akan diproses oleh komputer memiliki bentuk data hirarki, mulai dari yang paling kecil dan tidak dapat dibagi yaitu bits sampai ke dalam bentuk yang lebih kompleks

### *Bit*

Bit adalah kependekan dari **Binary Digit**, unit terkecil sebuah informasi dalam mesin komputer. Satu buah *bit* dapat menampung dua nilai diantaranya adalah **0** atau **1**. 

Jika terdapat 8 *bit* maka kita dapat menyebutnya sebagai **1 byte**.

### *Byte*

*Byte* adalah kependekan dari **Binary Term**, sebuah unit penyimpanan yang sudah memiliki kapabilitas paling sederhana untuk menyimpan sebuah karakter tunggal. 

1 *byte* = sekumpulan *bit* (terdapat delapan bit). Contoh : 0 1 0 1 1 0 1 0

1 *byte* dapat menyimpan karakter contoh : 'A' atau 'x' atau '$'

### *Bytes*

*Byte* adalah unit yang dipat digunakan untuk menyimpan informasi, seluruh penyimpanan diukur menggunakan *bytes*.

| **Number of Bytes**  | **Unit**         | **Representation**      |
| -------------------- | ---------------- | ----------------------- |
| 1                    | *Byte*           | *One Character*         |
| 1024                 | *KiloByte (Kb)*  | *Small Text In notepad* |
| 1,048,576            | *MegaByte (Mb)*  | *Ebook*                 |
| 1,073,741,824        | *GigaByte (Gb)*  | *Movie*                 |
| 1,099,511,627,776    | *TeraByte (Tb)*  | *Archive*               |
| *Approximately* 1015 | *PetaByte (Pb)*  | *Big Data*              |
| *Approximately* 1018 | *ExaByte (Eb)*   | *Big Data*              |
| *Approximately* 1021 | *ZettaByte (Zb)* | *Big Data*              |

### *Character*

Data dalam bentuk *bits* tidak mudah untuk dikelola sehingga perlu bentuk lain yang dapat digunakan manusia dan mempermudah proses pengelolaan informasi. Untuk mewujudkan hal tersebut data dalam *bits* harus bisa direpresentasikan dalam bentuk *character*. Seperti *decimal* *digit* (0-9), *letter* (A-Z dan a-z), dan *special symbol* (!@#$%^&*()-=_+). *Digit, letter* dan *symbol* disebut dengan *characters*. 

**Characters Set** adalah sekumpulan *characters* yang digunakan untuk menulis program dan merepresentasikan sebuah informasi. Dikarenakan komputer memproses informasi dalam bentuk 1 dan 0, maka sebuah *character* dapat direpresentasikan menggunakan 1 dan 0.

*Character* adalah unit terkecil dalam sistem teks dan memiliki makna. 

Sekumpulan *character* dapat membentuk *string* yang selanjutnya dapat digunakan untuk memvisualisasikan suatu bahasa verbal secara digital. Contoh *character* adalah abjad, angka dan simbol lainya. 

<img src="../../../assets/Grapheme.png" style="zoom:75%;" />

#### *ASCII*

Komputer merepresentasikan sebuah data dengan *number*, di awal pengembangan komputer tepatnya sekitar tahun 1940. Penggunaan teks dalam komputer untuk disimpan dan dimanipulasi dapat dilakukan, dengan cara merepresentasikan abjad dalam alfabet menggunakan *number*. Sebagai contoh angka 65 merepresentasikan huruf A dan angka 66 merepresentasikan huruf B hingga seterusnya.

Pada tahun 1950 saat komputer sudah semakin banyak digunakan untuk berkomunikasi, standar untuk merepresentasikan *text* agar dapat difahami oleh berbagai model dan *brand* komputer diusung. 

*ASCII* (*American Standard Code for Information Interchange*) adalah karya yang diusung, pertama dipublikasikan pada tahun 1963. Saat pertama kali dipublikasikan *ASCII* masih digunakan untuk *teleprinter technology*. *ASCII* terus direvisi hingga akhirnya *7-bit ASCII Table* diadopsi oleh *American National Standards Institute* (*ANSI*).

Dengan *7-bit* maka terdapat 128 *unique binary pattern* yang dapat digunakan untuk merepresentasikan suatu karakter. Kita dapat merepresentasikan **alphanumeric** (abjad a-z, A-Z, angka 0-9, dan *special character* seperti “!@#$%^&*”). 

Pada gambar di bawah ini huruf kapital G memiliki representasi dalam bentuk biner 100 0111 (7 *binary digit*) dan huruf kapital F memiliki representasi dalam bentuk *binary pattern* 100 0110 : 

![](../../../assets/ASCIICode.png)

Pada huruf kapital G angka 107 adalah representasi dalam *octal numeral system*, angka 71 adalah representasi dalam *decimal numeral system* dan angka 46 adalah representasi dalam *hexadecimal*. Representasi tidak hanya dalam bentuk *binary*. Untuk *table ASCII* lebih lengkapnya anda dapat melihat di wikipedia.