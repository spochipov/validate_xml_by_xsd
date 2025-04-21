# XML Validator

A simple Go application that validates XML files against XSD schemas and outputs all errors and inconsistencies.

> **Note:** For Russian documentation, see [README.ru.md](README.ru.md)

# XML Валидатор

Простое приложение на Go, которое проверяет XML файлы на соответствие XSD схемам и выводит все ошибки и несоответствия.

> **Примечание:** Для полной документации на русском языке, смотрите [README.ru.md](README.ru.md)

## Prerequisites | Предварительные требования

- Go 1.16 or higher | Go 1.16 или выше
- The application uses the `github.com/lestrrat-go/libxml2` package which requires the libxml2 library to be installed on your system. | Приложение использует пакет `github.com/lestrrat-go/libxml2`, который требует установки библиотеки libxml2 в вашей системе.

### Installing libxml2 | Установка libxml2

#### macOS
```
brew install libxml2
```

#### Ubuntu/Debian
```
sudo apt-get install libxml2-dev
```

#### CentOS/RHEL
```
sudo yum install libxml2-devel
```

## Building the Application | Сборка приложения

The application requires libxml2 and pkg-config to be installed. The build script will set the necessary environment variables for you. | Приложение требует установки libxml2 и pkg-config. Скрипт сборки установит необходимые переменные окружения за вас.

```
# Install dependencies (macOS) | Установка зависимостей (macOS)
brew install libxml2 pkg-config

# Build the application | Сборка приложения
./build.sh
```

## Usage | Использование

You can run the application directly: | Вы можете запустить приложение напрямую:

```
./validate_xml_by_xsd -xml <xml_file_path> -xsd <xsd_file_path>
```

Or use the provided wrapper script which sets the necessary environment variables: | Или использовать предоставленный скрипт-обертку, который устанавливает необходимые переменные окружения:

```
./validate.sh -xml <путь_к_xml_файлу> -xsd <путь_к_xsd_схеме>
```

### Examples | Примеры

Using the direct command: | Использование прямой команды:
```
./validate_xml_by_xsd -xml valid.xml -xsd schema.xsd
```

Using the wrapper script: | Использование скрипта-обертки:
```
./validate.sh -xml valid.xml -xsd schema.xsd
```

### Running Tests | Запуск тестов

A test script is provided to demonstrate validation with both valid and invalid XML files: | Предоставлен тестовый скрипт для демонстрации валидации как с корректными, так и с некорректными XML файлами:

```
./test.sh
```

This will run the validator against both sample files and display the results. | Это запустит валидатор для обоих примеров файлов и отобразит результаты.

## Sample Files | Примеры файлов

The repository includes sample files for testing: | Репозиторий включает примеры файлов для тестирования:

- `schema.xsd`: A sample XSD schema for a bookstore | Пример XSD схемы для книжного магазина
- `valid.xml`: A valid XML file that conforms to the schema | Корректный XML файл, соответствующий схеме
- `invalid.xml`: An invalid XML file with several validation errors | Некорректный XML файл с несколькими ошибками валидации

## Example Output | Пример вывода

For a valid XML file: | Для корректного XML файла:
```
XML validation successful! The XML file conforms to the XSD schema.
```

For an invalid XML file: | Для некорректного XML файла:
```
XML validation failed. Errors:
1. Element 'book': The attribute 'id' is required but missing.
2. Element 'year': 'Twenty Twenty-Two' is not a valid value of the atomic type 'xs:integer'.
3. Element 'category': This element is not expected. Expected is ( price ).
4. Element 'publisher': This element is not expected.
