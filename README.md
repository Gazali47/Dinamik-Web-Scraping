# Dinamik Web Scraping ve Sayfa Analiz Aracı

Bu proje, **Siber Tehdit İstihbaratı (CTI)** faaliyetlerinde kullanılmak üzere tasarlanmış, Go dili ile geliştirilen dinamik bir web veri toplama aracıdır. Araç, modern web uygulamalarındaki JavaScript tabanlı içerikleri Headless Browser yaklaşımıyla analiz eder.

## Özellikler

* **Dinamik İçerik Analizi:** Chromedp kullanarak tarayıcı seviyesinde render edilmiş DOM yapısına erişim sağlar.
* **Görsel Kanıt Üretimi:** Sayfanın tarama anındaki durumunu tam sayfa ekran görüntüsü (.png) olarak kaydeder.
* **Link Haritalama:** Sayfa üzerindeki tüm harici bağlantıları (linkleri) ayıklayarak liste oluşturur.
* **HTTP Durum Tespiti:** Hedef URL'nin HTTP yanıt kodunu tespit eder ve loglar.
* **Otomatik Dosya Sistemi:** Çıktıları hedef site adına göre dinamik olarak isimlendirip depolar.

## Kullanılan Teknolojiler

* **Dil:** Go (Golang) 
* **Motor:** [chromedp](https://github.com/chromedp/chromedp) (Chrome DevTools Protocol) 
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
