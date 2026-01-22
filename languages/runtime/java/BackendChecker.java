import com.google.gson.Gson;
import java.io.ByteArrayOutputStream;
import java.io.FileWriter;
import java.io.IOException;
import java.io.PrintStream;
import java.util.ArrayList;
import java.util.List;
import java.util.function.Function;

public class BackendChecker {
  public static class FailedTestcase {
    public int id;
    public Object input;
    public Object expected;
    public Object got;
    public String stdOut;
  }

  public static class RunResult {
    public boolean isPassed = true;
    public List<FailedTestcase> failedTestcases = new ArrayList<>();
    public double memoryUsage;
    public double executionTime;
    public String stdOut;
    public String stdErr;
  }

  public static class StdoutCapture {
    public Object result;
    public String stdout;
    
    public StdoutCapture(Object result, String stdout) {
      this.result = result;
      this.stdout = stdout;
    }
  }

  public static <T> StdoutCapture captureStdout(java.util.function.Supplier<T> fn) {
    ByteArrayOutputStream baos = new ByteArrayOutputStream();
    PrintStream oldOut = System.out;
    System.setOut(new PrintStream(baos));
    
    T result = fn.get();
    
    System.setOut(oldOut);
    String stdout = baos.toString();
    
    return new StdoutCapture(result, stdout);
  }

  public static void writeResults(RunResult results, String filename) {
    try (FileWriter file = new FileWriter(filename)) {
      Gson gson = new Gson();
      file.write(gson.toJson(results));
    } catch (IOException e) {
      System.err.println("Failed to write results file.");
      System.exit(1);
    }
  }

  public static <T> void runTests(List<T> testcases, Function<List<T>, RunResult> runChecker) {
    long startTime = System.currentTimeMillis();

    RunResult results = runChecker.apply(testcases);

    long endTime = System.currentTimeMillis();
    results.executionTime = (double) (endTime - startTime);

    writeResults(results, "results.json");
  }

  public static void main(String[] args) {
    List<Driver.Testcase> testcases = Driver.loadTestcases("testcases.json");
    runTests(testcases, MainChecker::runChecker);
  }
}
