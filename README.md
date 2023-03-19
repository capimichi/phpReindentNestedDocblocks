# Php Reindent Nested Docblocks
Reindent nested docblocks (for example in swagger comments, ecc.)

## Motivation

You have worked in php projects like Symfony, Laravel, ecc. and you have to write a lot of docblocks. 

If they have a nested structure, for example when working with Swagger, you will end up with a mess.

So you have to find a way to reindent them.

If you usually do it manually, you will end up with a lot of time wasted.

So this project will help you to reindent them automatically:

![Preview Indent](https://raw.githubusercontent.com/capimichi/phpReindentNestedDocblocks/main/preview.gif)

## Usage

You can use in the terminal:

```bash
./php_reindent_nested_docblocks input_file.php [output_file.php]
```

The output file is optional, if you don't specify it, the input file will be overwritten.

## Installation

Just download the file `php_reindent_nested_docblocks` and make it executable:

```bash
chmod +x php_reindent_nested_docblocks
```

Then you can use it as described above.

## Next steps

- [ ] Include php attributes nested docblocks to the reindentation

## Contributing

If you want to contribute, you can fork the project and make a pull request.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details




