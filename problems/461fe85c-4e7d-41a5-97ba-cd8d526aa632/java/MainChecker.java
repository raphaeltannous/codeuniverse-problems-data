import java.util.List;

public class MainChecker {
  public static BackendChecker.RunResult runChecker(List<Driver.Testcase> testcases) {
    BackendChecker.RunResult runResult = new BackendChecker.RunResult();
    Main.Solution solution = new Main.Solution();

    for (Driver.Testcase testcase : testcases) {
      String input = testcase.input;
      boolean expected = testcase.expected;
      
      BackendChecker.StdoutCapture capture = BackendChecker.captureStdout(() -> 
        solution.isValid(input)
      );
      
      boolean got = (boolean) capture.result;

      if (expected != got) {
        BackendChecker.FailedTestcase failedTestcase = new BackendChecker.FailedTestcase();
        failedTestcase.id = testcase.id;
        failedTestcase.input = testcase.input;
        failedTestcase.expected = testcase.expected;
        failedTestcase.got = got;
        failedTestcase.stdOut = capture.stdout;

        runResult.failedTestcases.add(failedTestcase);
        runResult.isPassed = false;
      }
    }

    return runResult;
  }
}
