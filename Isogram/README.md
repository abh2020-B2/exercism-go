# Isogram

A Go exercise from [Exercism](https://exercism.org/tracks/go/exercises/isogram) focused on string manipulation and checking for duplicate letters.

## Description

An isogram is a word or phrase in which no letter occurs more than once.

The task is to implement the `IsIsogram()` function, which determines whether a given string is an isogram.

The comparison is case-insensitive, meaning uppercase and lowercase versions of the same letter are considered identical. Spaces and hyphens are ignored.

### Examples

| Word or Phrase | Isogram |
| -------------- | ------- |
| `lumberjacks`  | `true`  |
| `background`   | `true`  |
| `downstream`   | `true`  |
| `six-year-old` | `true`  |
| `isograms`     | `false` |

## Concepts

- Strings
- String Iteration
- Runes
- Maps
- Boolean Values
- Conditional Statements
- Loops
- Case Conversion
- Duplicate Detection

## Function

| Function      | Description                                                    |
| ------------- | -------------------------------------------------------------- |
| `IsIsogram()` | Checks whether a word or phrase contains any repeated letters. |

## Key Learnings

- Iterating over characters in a string.
- Handling uppercase and lowercase letters.
- Ignoring spaces and hyphens.
- Detecting duplicate characters.
- Using maps to keep track of previously encountered letters.

## Language

Go

---

# ایزوگرام (Isogram)

این تمرین از مسیر یادگیری زبان Go در سایت [Exercism](https://exercism.org/tracks/go/exercises/isogram) است و بر پردازش رشته‌ها و تشخیص حروف تکراری تمرکز دارد.

## توضیحات تمرین

ایزوگرام به کلمه یا عبارتی گفته می‌شود که هیچ حرفی در آن بیش از یک بار تکرار نشده باشد.

در این تمرین باید تابع `IsIsogram()` را پیاده‌سازی کنیم تا مشخص کند یک رشته ایزوگرام است یا خیر.

در بررسی حروف، بزرگی و کوچکی حروف تفاوتی ندارد؛ یعنی حروف بزرگ و کوچک یکسان محسوب می‌شوند. همچنین فاصله‌ها و خط تیره‌ها در بررسی تکرار حروف نادیده گرفته می‌شوند.

### مثال‌ها

| کلمه یا عبارت  | نتیجه   |
| -------------- | ------- |
| `lumberjacks`  | `true`  |
| `background`   | `true`  |
| `downstream`   | `true`  |
| `six-year-old` | `true`  |
| `isograms`     | `false` |

## مفاهیم استفاده‌شده

- رشته‌ها (Strings)
- پیمایش رشته‌ها (String Iteration)
- رون‌ها (Runes)
- مپ‌ها (Maps)
- مقادیر بولی (Boolean Values)
- دستورات شرطی
- حلقه‌ها
- تبدیل حروف کوچک و بزرگ
- تشخیص مقادیر تکراری

## تابع

| تابع          | توضیحات                                                            |
| ------------- | ------------------------------------------------------------------ |
| `IsIsogram()` | بررسی می‌کند که آیا یک کلمه یا عبارت دارای حروف تکراری است یا خیر. |

## نکات یادگیری

- پیمایش کاراکترهای یک رشته
- مدیریت حروف بزرگ و کوچک
- نادیده گرفتن فاصله‌ها و خط تیره‌ها
- تشخیص کاراکترهای تکراری
- استفاده از Map برای نگهداری حروفی که قبلاً مشاهده شده‌اند

## زبان برنامه‌نویسی

Go
