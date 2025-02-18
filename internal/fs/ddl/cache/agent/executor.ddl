metadata :name        => "executor",
         :description => "Choria Process Executor Management",
         :author      => "R.I.Pienaar <rip@devco.net>",
         :license     => "Apache-2.0",
         :version     => "0.29.4",
         :url         => "https://choria.io",
         :provider    => "golang",
         :timeout     => 20


action "status", :description => "Requests the status of a job by ID" do
  display :always

  input :id,
        :prompt      => "Job ID",
        :description => "The unique ID for the job",
        :type        => :string,
        :validation  => '.',
        :maxlength   => 20,
        :optional    => false




  output :action,
         :description => "The RPC Action that started the process",
         :type        => "string",
         :display_as  => "Action"

  output :agent,
         :description => "The RPC Agent that started the process",
         :type        => "string",
         :display_as  => "Agent"

  output :caller,
         :description => "The Caller ID who started the process",
         :type        => "string",
         :display_as  => "Caller"

  output :exit_code,
         :description => "The exit code the process terminated with",
         :type        => "integer",
         :display_as  => "Exit Code"

  output :pid,
         :description => "The OS Process ID",
         :type        => "integer",
         :display_as  => "Pid"

  output :requestid,
         :description => "The Request ID that started the process",
         :type        => "string",
         :display_as  => "Request ID"

  output :running,
         :description => "Indicates if the process is still running",
         :type        => "boolean",
         :display_as  => "Running"

  output :start_time,
         :description => "Time that the process started",
         :type        => "string",
         :display_as  => "Started"

  output :started,
         :description => "Indicates if the process was started",
         :type        => "boolean",
         :display_as  => "Started"

  output :stderr_bytes,
         :description => "The number of bytes of STDERR output available",
         :type        => "integer",
         :display_as  => "STDERR Bytes"

  output :stdout_bytes,
         :description => "The number of bytes of STDOUT output available",
         :type        => "integer",
         :display_as  => "STDOUT Bytes"

  output :terminate_time,
         :description => "Time that the process terminated",
         :type        => "string",
         :display_as  => "Terminated"

end

