# ReRunner

A file monitor.Run some commands if file have changed.

## Getting Started

 - Go to the Release page and Download ReRunner.
 - Save ReRunner to the folder where code is located.
 - Run ReRunner once and Editing the RRSetting.json.
 - Run ReRuner again and enjoy it. 

## Edit the RRSetting.json
You can format it with json rules and it will maybe like - 
```json
{
    "ByLine": false,
    "Tasks": [
        {
            "RunCommands": [
                [
                    "cmd.exe",
                    "/c",
                    "cls"
                ],
                [
                    "cmd.exe",
                    "/c",
                    "dir"
                ]
            ],
            "WatchFiles": {
                "filename1": "filehash",
                "filename2": "filehash"
            }
        },
        {
            "RunCommands": [
                [
                    "cmd.exe",
                    "/c",
                    "cls"
                ],
                [
                    "cmd.exe",
                    "/c",
                    "dir"
                ]
            ],
            "WatchFiles": {
                "filename3": "filehash",
                "filename4": "filehash"
            }
        }
    ],
    "Version": "Mar 24,2020."
}
```
- `"ByLine": false`
   - If output is stream,you can set .But if you don't know what it is.Just skip.
- `"Tasks": [ {} ]`
   - a task list
- `            "RunCommands": [
                [
                    "cmd.exe",
                    "/c",
                    "cls"
                ],
                [
                    "cmd.exe",
                    "/c",
                    "dir"
                ]
            ]`
   - run cmd and give it /c , cls args.
   - if you cannot know what it is,search cmd at first.
- `            "WatchFiles": {
                "filename3": "filehash",
                "filename4": "filehash"
            }`
   - the filename or path with exe work path
- "Version": "Mar 24,2020."
   - It often get errors if version is not feat with program,don't need to change it.

## Running the tests

 * Windows
   * open command or powershell with the path
   * type `ReRunner.exe` and enter to get RRSetting.json
   * do it again after you edited RRSetting.json 
 - Unix
   - open terminal with the path
   - type `./ReRunner` and enter to get RRSetting.json
   - do it again after you edited RRSetting.json 

### Coding style

`go fmt` is enough.

## Built With

* [Golang](http://www.dropwizard.io/1.0.2/docs/) - version go1.13.4

## Contributing

Do it what you want under License.

## Authors

* **Anest Larry** - A Pythonista.

See also the list of contributors who participated in this project.

## License

This project is licensed under the "Anti 996" License - see the [LICENSE.md](LICENSE) file for details