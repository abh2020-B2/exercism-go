# Cars Assemble

**Exercise:** [Exercism - Cars Assemble](https://exercism.org/tracks/go/exercises/cars-assemble)

## Description

This exercise focuses on numerical operations, type conversions, and arithmetic operators in Go.

The goal is to analyze a car assembly line by calculating the number of working cars produced per hour and per minute, as well as the total production cost.

## Concepts

- Numbers (`int`, `float64`, `uint`)
- Arithmetic Operators
- Type Conversion
- Integer Division
- Modulo Operator (`%`)
- Functions
- Return Values

## Solution

The solution implements three functions:

- `CalculateWorkingCarsPerHour()`: Calculates the number of successfully produced cars per hour based on the production rate and success rate.
- `CalculateWorkingCarsPerMinute()`: Calculates the number of successfully produced cars per minute.
- `CalculateCost()`: Calculates the total production cost using bulk pricing for groups of ten cars and individual pricing for the remaining cars.

## Key Learnings

- Working with different numeric types in Go.
- Converting between numeric types.
- Using arithmetic operators and integer division.
- Applying the modulo operator to calculate remainders.
- Understanding type compatibility in arithmetic operations.

## Language

Go

---

# مونتاژ خودروها (Cars Assemble)

## توضیحات تمرین

این تمرین بر عملیات عددی، تبدیل نوع داده‌ها و عملگرهای حسابی در زبان برنامه‌نویسی Go تمرکز دارد.

هدف تمرین، تحلیل عملکرد یک خط مونتاژ خودرو از طریق محاسبه تعداد خودروهای سالم تولیدشده در هر ساعت و هر دقیقه و همچنین هزینه کل تولید است.

## مفاهیم استفاده‌شده

- **Numbers (انواع داده عددی):** آشنایی با انواع `int`، `float64` و `uint`.
- **Arithmetic Operators (عملگرهای حسابی):** استفاده از عملگرهای جمع، تفریق، ضرب و تقسیم.
- **Type Conversion (تبدیل نوع داده):** تبدیل مقادیر بین انواع عددی مختلف.
- **Integer Division (تقسیم صحیح):** انجام تقسیم اعداد صحیح بدون در نظر گرفتن بخش اعشاری.
- **Modulo Operator (عملگر باقیمانده):** استفاده از عملگر `%` برای محاسبه باقیمانده تقسیم.
- **Functions (توابع):** تعریف و پیاده‌سازی توابع محاسباتی.
- **Return Values (مقادیر بازگشتی):** بازگرداندن نتایج محاسبات با انواع داده مشخص.

## روش حل

برای حل تمرین، سه تابع اصلی پیاده‌سازی می‌شود:

1. `CalculateWorkingCarsPerHour()`: محاسبه تعداد خودروهای سالم تولیدشده در هر ساعت با توجه به سرعت تولید و درصد موفقیت خط مونتاژ.
2. `CalculateWorkingCarsPerMinute()`: محاسبه تعداد خودروهای سالم تولیدشده در هر دقیقه.
3. `CalculateCost()`: محاسبه هزینه کل تولید خودرو با در نظر گرفتن تخفیف تولید گروهی؛ هر ۱۰ خودرو با هزینه ۹۵ هزار دلار و هر خودروی باقی‌مانده با هزینه ۱۰ هزار دلار تولید می‌شود.

## نکات یادگیری

- کار با انواع داده عددی مختلف در Go.
- تبدیل نوع داده‌ها برای انجام محاسبات.
- تفاوت تقسیم صحیح و تقسیم اعشاری.
- کاربرد عملگر باقیمانده `%`.
- رعایت سازگاری نوع داده‌ها در عملیات حسابی.

## زبان برنامه‌نویسی

Go (Golang)
