require 'json'
require 'stringio'

module BackendChecker
  class FailedTestcase
    attr_accessor :id, :input, :expected, :got, :stdOut

    def initialize(id, input, expected, got, std_out)
      @id = id
      @input = input
      @expected = expected
      @got = got
      @stdOut = std_out
    end

    def to_h
      {
        id: @id,
        input: @input,
        expected: @expected,
        got: @got,
        stdOut: @stdOut
      }
    end
  end

  class RunResult
    attr_accessor :isPassed, :failedTestcases, :memoryUsage, :executionTime, :stdOut, :stdErr

    def initialize
      @isPassed = true
      @failedTestcases = []
      @memoryUsage = 0.0
      @executionTime = 0.0
      @stdOut = ''
      @stdErr = ''
    end

    def to_h
      {
        isPassed: @isPassed,
        failedTestcases: @failedTestcases.map(&:to_h),
        memoryUsage: @memoryUsage,
        executionTime: @executionTime,
        stdOut: @stdOut,
        stdErr: @stdErr
      }
    end
  end

  def self.capture_stdout
    old_stdout = $stdout
    $stdout = StringIO.new
    
    result = yield
    
    output = $stdout.string
    $stdout = old_stdout
    
    { result: result, stdout: output }
  end

  def self.write_results(results, filename)
    File.write(filename, JSON.pretty_generate(results.to_h))
  rescue => e
    $stderr.puts "Failed to write results file."
    exit(1)
  end

  def self.run_tests(testcases, run_checker)
    start_time = Time.now
    
    results = run_checker.call(testcases)
    
    end_time = Time.now
    results.executionTime = ((end_time - start_time) * 1000).to_f
    
    write_results(results, 'results.json')
  end
end
