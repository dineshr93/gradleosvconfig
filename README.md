# gradleosvconfig
osv config for gradle build file


# Features

1. Produces customized report where we can see vulnerability, OSS name, affected source path details all in one report
2. Color coded

   low risk = no color

   medium risk = Yellow

   High risk = Red

3. Omits all other files which has no vulnerabilities.

### Prerequiites:

Go to Your Blackduck Project > Generate 'Create Version detail report' > checkbox Source and Vulnerabilities checked.

## How to install

```sh

osv gradleconfig.txt gradle_app_dir sourcing_env_shell_script
```

## Command to run

```sh


usage:osv gradleconfig.txt gradle_app_dir sourcing_env_shell_script

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
