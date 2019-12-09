namespace ConsoleRedirect
{
    using System;
    using System.Diagnostics;
    using System.IO;

    internal class TestCase
    {
        public TestCase(bool createNoWindow, bool redirectIO)
        {
            CreateNoWindow = createNoWindow;
            RedirectIO = redirectIO;
        }

        public string DisplayText => $"UseShell: [{!RedirectIO}], CreateNoWindow: [{CreateNoWindow}], RedirectIO: [{RedirectIO}]";
        public bool CreateNoWindow { get; }
        public bool RedirectIO { get; }
    }

    internal sealed class TestExecutionEngine
    {
        private readonly int waitTimeInMs;

        public TestExecutionEngine(int waitTimeInMs)
        {
            this.waitTimeInMs = waitTimeInMs < 2000 || waitTimeInMs > 10000 ? 2000 : waitTimeInMs;
        }

        public void ExecuteTest(TestCase testCase)
        {
            Console.WriteLine($"Press Enter for the test: {testCase.DisplayText}");
            Console.ReadLine();
            var startInfo = GetTestStartInfo(testCase);
            var process = Process.Start(startInfo);

            if (testCase.RedirectIO)
            {
                RedirectProcessIO(process);
            }
            process.WaitForExit(waitTimeInMs);
            process.Kill();
        }

        private ProcessStartInfo GetTestStartInfo(TestCase testCase)
        {
            var pStartInfo = new ProcessStartInfo
            {
                WorkingDirectory = Directory.GetCurrentDirectory(),
                UseShellExecute = !testCase.RedirectIO,
                CreateNoWindow = testCase.CreateNoWindow,
                FileName = Path.Combine(Directory.GetCurrentDirectory(), "test.bat"),
                RedirectStandardError = testCase.RedirectIO,
                RedirectStandardOutput = testCase.RedirectIO
            };

            return pStartInfo;
        }

        private void RedirectProcessIO(Process process)
        {
            process.ErrorDataReceived += (sender, args) =>
            {
                if (!string.IsNullOrEmpty(args?.Data))
                {
                    Console.Error.WriteLine(args.Data);
                }
            };
            process.OutputDataReceived += (sender, args) =>
            {
                if (!string.IsNullOrEmpty(args?.Data))
                {
                    Console.WriteLine(args.Data);
                }
            };
            process.BeginErrorReadLine();
            process.BeginOutputReadLine();
        }
    }

    internal class Program
    {
        internal static void Main(string[] args)
        {
            var cases = new[]
            {
                new TestCase(true, false),
                new TestCase(false, false),
                new TestCase(true, true),
                new TestCase(false, true)
            };

            var testEngine = new TestExecutionEngine(3000);

            foreach (var item in cases)
            {
                testEngine.ExecuteTest(item);
            }

            Console.WriteLine("Finish");
        }
    }
}
