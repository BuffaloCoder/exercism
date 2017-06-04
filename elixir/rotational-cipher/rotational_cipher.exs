defmodule RotationalCipher do
  @doc """
  Given a plaintext and amount to shift by, return a rotated string.

  Example:
  iex> RotationalCipher.rotate("Attack at dawn", 13)
  "Nggnpx ng qnja"
  """
  @spec rotate(text :: String.t(), shift :: integer) :: String.t()
  def rotate(text, shift) do
    text
    |> String.graphemes
    |> Enum.map(fn char ->
      cond do
        not(char =~ ~r/^\p{L}$/u) -> char
        char =~ ~r/^\p{Lu}$/u -> string_conversion(65, char, shift)
        char =~ ~r/^\p{Ll}$/u -> string_conversion(97, char, shift)
      end
    end)
    |> Enum.join
  end

  defp string_conversion(base, char, shift) do
    char
    |> to_charlist
    |> Enum.map(fn char_ascii ->
      char_ascii
      |> Kernel.-(base) # Account for ascii base (65 for capital, 97 for lower)
      |> Kernel.+(shift) # Shift it
      |> rem(26) # Modulo by letters in alphabet
      |> Kernel.+(base) # Re-apply base
    end)
    |> to_string
  end
end

