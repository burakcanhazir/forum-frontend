# FORUM-PROJECT
Programla dilleri ve benzer teknolojiler için yardımlaşma adını yapılmış bir web sitesi

## TANITIM

Bu proje, programlama dilleri, algoritmalar ve yazılım teknolojileri hakkında kullanıcıların problem ve çözümlerini birbirlerine aktardığı bir forum platformudur. 
Projede, RESTful API'ler kullanılarak kullanıcılar arasında veri paylaşımı yapılmaktadır. Backend kısmında Go dili kullanılmıştır, frontend kısmı ise HTML, CSS ve JavaScript ile geliştirilmiştir.
Projeye Docker ile erişim sağlayarak, tüm servisleri kolayca çalıştırabilirsiniz.

Proje üç ana bileşene ayrılmıştır:
- **backend**: Go ile yazılmış API'ler
- **frontend**: HTML, CSS, JS ile yapılmış arayüz
- **deploy**: Docker ile dağıtım işlemleri

## KURULUM

Projeyi lokalinizde çalıştırabilmek için aşağıdaki adımları takip edin:

1. **Docker'ı kurun**: [Docker Kurulum Videosu](https://www.youtube.com/watch?v=iqqDU2crIEQ)
2. **Tüm repoları lokalinizde indirin**:
   - [Forum Frontend](https://github.com/burakcanhazir/forum-deploy) - Frontend kodu
   - [Forum Backend](https://github.com/burakcanhazir/burak-forumend) - Backend kodu
   - [Forum Deployment](https://github.com/burakcanhazir/burakforum) - Dağıtım ve Docker yapılandırması

---

### **. Kullanım Bölümü**
Kullanım kısmı çok net, fakat bir kaç yazım hatası ve daha fazla açıklama ekleyerek kısımları daha anlaşılır hale getirebilirsiniz.

```markdown
## KULLANIM

1. **Backend yapılandırması**: 
   Backend reposunda bulunan `.env.example` dosyasının içeriğini kopyalayın.
   
2. **.env Dosyası Oluşturun**:
   `forum-backend` klasöründe yeni bir `.env` dosyası oluşturun ve içeriği kopyaladığınız verilerle yapılandırın.

3. **Docker ile Başlatma**:
   `docker-deployment` klasörüne gidin ve terminalinizde şu komutu çalıştırın:
   ```bash
   docker-compose up --build

## Lisans

Bu proje MIT Lisansı ile lisanslanmıştır - [Lisans dosyasını](LICENSE) inceleyin. 

## İletişim

Proje ile ilgili herhangi bir sorunuz veya öneriniz varsa, [hazirburakcan@gmail.com] adresi ile iletişime geçebilirsiniz.


## Yazarlar

- [Burakcan Hazir](https://github.com/burakcanhazir)


