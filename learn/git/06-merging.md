# Senaryo 6: Merging

## Öğrenilenler
- git merge: bir branch'teki degisiklikleri baska bir branch'e entegre eder
- Fast-forward merge: eger hedef branch (main) merge edilen branch'ten sonra
  hic degismemisse, Git sadece pointer'i ileri tasir, ayri bir merge commit
  olusturmaz
- Fast-forward olmayan durumda (main'de de yeni commit'ler varsa) Git otomatik
  bir "merge commit" olusturur, iki gelisim hattini birlestirir