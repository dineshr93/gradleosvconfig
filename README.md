# gradleosvconfig

osv vulnerability check binary config for gradle project

# Features

1. configure in CI/CD pipeline for gradle project.
2. If 0 vulnerability is detected it exits sucessfully
3. If any vulnerability is there it prints the details and fails the execution
4. to reset the config

```js
git status && echo "============" && git reset --hard && echo "============" && git clean -fd && echo "============" && git status
```

# Sample Output

![Sample](https://github.com/dineshr93/gradleosvconfig/blob/master/sample.png?raw=true)


### Prerequiites:

1.OSV

2.jq

3.goc gradleosvconfig

## How to install

```sh
with devenv
goc gradleconfig.txt gradle_app_dir sourcing_env_shell_script

without
goc gradleconfig.txt gradle_app_dir
```

## Command to run

```sh


usage:goc gradleconfig.txt gradle_app_dir sourcing_env_shell_script

options:
 gradleconfig.txt  find this file in this repo. config for releaseRuntimeClasspath. you can alter on your own
 gradle_app_dir      Gradle project directory
 sourcing_env_shell_script          your_env_path_gradle_repo_etc



```

## Issues

Please send your bugs to dineshr93@gmail.com

## License

[MIT](LICENSE)

```
MIT License

Copyright (c) 2022 Dinesh Ravi

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
```
