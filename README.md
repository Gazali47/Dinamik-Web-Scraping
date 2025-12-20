# Dinamik Web Scraping ve Sayfa Analiz Aracı

[cite_start]Bu proje, **Siber Tehdit İstihbaratı (CTI)** faaliyetlerinde kullanılmak üzere tasarlanmış, Go dili ile geliştirilen dinamik bir web veri toplama aracıdır[cite: 16]. [cite_start]Araç, modern web uygulamalarındaki JavaScript tabanlı içerikleri Headless Browser yaklaşımıyla analiz eder[cite: 25].

## Özellikler

* [cite_start]**Dinamik İçerik Analizi:** Chromedp kullanarak tarayıcı seviyesinde render edilmiş DOM yapısına erişim sağlar[cite: 29].
* [cite_start]**Görsel Kanıt Üretimi:** Sayfanın tarama anındaki durumunu tam sayfa ekran görüntüsü (.png) olarak kaydeder[cite: 36, 71].
* [cite_start]**Link Haritalama:** Sayfa üzerindeki tüm harici bağlantıları (linkleri) ayıklayarak liste oluşturur[cite: 22, 103].
* [cite_start]**HTTP Durum Tespiti:** Hedef URL'nin HTTP yanıt kodunu tespit eder ve loglar[cite: 34, 49].
* [cite_start]**Otomatik Dosya Sistemi:** Çıktıları hedef site adına göre dinamik olarak isimlendirip depolar[cite: 61].

## Kullanılan Teknolojiler

* [cite_start]**Dil:** Go (Golang) [cite: 27]
* [cite_start]**Motor:** [chromedp](https://github.com/chromedp/chromedp) (Chrome DevTools Protocol) [cite: 28]
* **Kütüphane:** `golang.org/x/net/html` (Link ayrıştırma için)

## Kurulum ve Çalıştırma

1. Projeyi klonlayın:
   ```bash
   git clone [https://github.com/Gazali47/Dinamik-Web-Scraping.git](https://github.com/Gazali47/Dinamik-Web-Scraping.git)
   cd Dinamik-Web-Scraping
   Bağımlılıkları yükleyin:

go mod init web-scraper
go get [github.com/chromedp/chromedp](https://github.com/chromedp/chromedp)
go get golang.org/x/net/html
Programı çalıştırın:
go run main.go <HEDEF_URL>
