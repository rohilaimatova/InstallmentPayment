# 📦 Installment Payment

**Installment Payment** — это консольное приложение на Go, предназначенное для расчёта суммы рассрочки при покупке товаров в интернет-магазине. Программа запрашивает у пользователя данные, производит расчёт общей суммы с учётом процентов и эмулирует отправку СМС с информацией о покупке.

---

## 🧾 Пример задачи

Финансовое учреждение заключило договор с интернет-магазином. В зависимости от категории товара и срока рассрочки к сумме применяются надбавки:

| Категория   | Поддерживаемые сроки (мес) | Надбавка за каждые +3 мес (от 3 мес) |
|-------------|-----------------------------|---------------------------------------|
| Смартфон    | 3, 6, 9                     | +3%                                   |
| Компьютер   | 3, 6, 9, 12                 | +4%                                   |
| Телевизор   | 3, 6, 9, 12, 15, 18         | +5%                                   |

Например:
- Смартфон на 3 мес за 1000 сомони → **1000 сомони**
- Смартфон на 9 мес за 1000 сомони → **1060 сомони** (6% надбавка = 2 шага * 3%)

---


