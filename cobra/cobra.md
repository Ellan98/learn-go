## Cobra

### cobra-cli初始化

```cobra
cobra-cli init
```

### 向项目添加命令

初始化cobra应用程序后，可以继续使用cobra生成器向应用程序添加其他命令。执行此操作的命令是`cobra-cli-add`。
假设您创建了一个应用程序，并希望使用以下命令：

- app serve
- app config
- app config create

在您的项目目录（main.go文件所在的位置）中，您将运行以下命令：

```
cobra-cli add serve
cobra-cli add config
cobra-cli add create -p 'configCmd'
```



`cobracli-add支持与cobracli-init（如上所述）相同的所有可选标志。

输入cobra-cli

```powershell
Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.

Usage:
  cobra-cli [command]

Available Commands:
  add         Add a command to a Cobra Application
  completion  Generate the autocompletion script for the specified shell
  help        Help about any command
  init        Initialize a Cobra Application

Flags:
  -a, --author string    author name for copyright attribution (default "YOUR NAME")
      --config string    config file (default is $HOME/.cobra.yaml)
  -h, --help             help for cobra-cli
  -l, --license string   name of license for the project
      --viper            use Viper for configuration

Use "cobra-cli [command] --help" for more information about a command.
```

### 修改信息

```
cobra-cli init --author Ellan -l MIT
```

### 命令基本结构

```go

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "learn_cobra",
	Short: "A brief description of your application",
	Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.learn_cobra.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

```

##### use

##### short

##### log

添加子命令

```
cobra-cli add version
```

执行子命令

```
go run main.go version    || go run main.go
```

查看命令相关信息

```
go run main.go -h || go run main.go version -h
```

为命令添加一个选项

```go
versionCmd.Flags().StringP("author", "a", "Ellan-copy", "作者name")
```



```

Flags:
  -a, --author string   作者name (default "Ellan-copy")
  -h, --help            help for version
```

获取命令信息

```go
Run: func(cmd *cobra.Command, args []string) {
		//cmd当前指向自身
		author, err := cmd.Flags().GetString("author")
		if err != nil {
			fmt.Println("请输入作者信息")
		}
		fmt.Println("作者信息", author)
	},
```

修改命令信息

```
go run main.go version --author ellan
作者信息 ellan
```

修改帮助信息 需要调用 setHelp...函数





```go
Run: func(cmd *cobra.Command, args []string) {
		//cmd当前指向自身
		author, err := cmd.Flags().GetString("author")
		if err != nil {
			fmt.Println("请输入作者信息")
		}
		fmt.Println("作者信息", author)
    fmt.Printf("当前接收的参数：%v\n", args)
	},
	
	
	//args []string
	
	
	接收命令行后面的参数
	
	E:\go\src\learn-go\cobra>go run main.go version --author ellan
作者信息 ellan
当前接收的参数：[]

E:\go\src\learn-go\cobra>go run main.go version --author ellan  test
作者信息 ellan
当前接收的参数：[test]

	
	
	
	
	
	

```

