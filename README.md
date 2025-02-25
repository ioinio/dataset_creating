<B>Dataset Image Processor</B>

Этот скрипт на Go предназначен для обработки изображений, изменения их размера, увеличения контрастности и сохранения в новую папку. Скрипт автоматически обрабатывает все изображения в указанной папке, изменяет их размер до 728x728 пикселей, увеличивает контрастность и сохраняет их в формате JPEG с новым порядковым именем.

<B>Установка</B>

1. Убедитесь, что у вас установлен Go (версия 1.16 или выше).

2. Скопируйте репозиторий или скачайте исходный код.

3. Установите необходимые зависимости:
   
<code>go get github.com/nfnt/resize</code>


<B>Использование</B>

1. Поместите изображения, которые вы хотите обработать, в папку <code>img.</code>

2. Запустите скрипт:

3. Обработанные изображения будут сохранены в папку <code>dataset</code> с именами в формате <code>img_<номер>.jpg.</code>

<B>Структура проекта</B>

<code>img/</code> — папка с исходными изображениями.

<code>dataset/</code> — папка с обработанными изображениями.

<code>main.go</code> — основной файл скрипта.

<B>Настройки</B>

<code>inputFolder</code> — папка с исходными изображениями (по умолчанию <code>img</code>).

<code>outputFolder</code> — папка для сохранения обработанных изображений (по умолчанию <code>dataset</code>).

<code>imageSize</code> — размер изображения после обработки (по умолчанию 728x728 пикселей).

<code>contrastFactor</code> — коэффициент увеличения контрастности (по умолчанию 1.0).

<B>Пример</B>

![image](https://github.com/user-attachments/assets/f2bb68ef-cf49-4249-b3ad-17a4412a3972)


После выполнения скрипта все изображения из папки <code>img</code> будут обработаны и сохранены в папку <code>dataset</code> с новыми именами.

<B>Зависимости</B>

<code>github.com/nfnt/resize</code> — библиотека для изменения размера изображений.

<B>Лицензия</B>

Этот проект распространяется под лицензией MIT. Подробности см. в файле <a href="https://github.com/ioinio/dataset_creating/blame/main/LICENSE">LICENSE</a>.
