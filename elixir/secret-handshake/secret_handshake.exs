defmodule SecretHandshake do
  use Bitwise
  @commands %{1 => ["wink"], 2 => ["double blink"], 4 => ["close your eyes"], 8 => ["jump"]}
  @doc """
  Determine the actions of a secret handshake based on the binary
  representation of the given `code`.

  If the following bits are set, include the corresponding action in your list
  of commands, in order from lowest to highest.

  1 = wink
  10 = double blink
  100 = close your eyes
  1000 = jump

  10000 = Reverse the order of the operations in the secret handshake
  """
  @spec commands(code :: integer) :: list(String.t())
  def commands(code) do
    []
    |> check_command(code, 1)
    |> check_command(code, 2)
    |> check_command(code, 4)
    |> check_command(code, 8)
    |> (fn commands ->
      if (code &&& 16) != 0, do: Enum.reverse(commands), else: commands
    end).()
  end

  defp check_command(commands, code, check) do
    if (code &&& check) != 0, do: commands ++ Map.get(@commands, check), else: commands
  end
end

