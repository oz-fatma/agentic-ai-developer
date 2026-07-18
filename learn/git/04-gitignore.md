# Senaryo 4: Ignoring Files

## Öğrenilenler
- .gitignore: Git'in hangi dosya/klasörleri takip etmeyecegini belirten dosya
- Repo'nun kök dizininde olmalı
- Wildcard (*) kullanarak uzantı bazlı kurallar yazılabilir (örn. *.log)
- Klasör sonuna / eklenerek tüm klasör yok sayılabilir (örn. node_modules/)
- Zaten commit edilmiş bir dosya .gitignore'a eklense bile takip edilmeye devam eder
  (önce git rm --cached ile çıkarılması gerekir)git add .
git commit -m "Scenario 4: gitignore eklendi"
git push origin main