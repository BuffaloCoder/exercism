defmodule Roman do
  @doc """
  Convert the number to a roman number.
  """
  @spec numerals(pos_integer) :: String.t()
  def numerals(number) do
    do_numerals(number)
  end

  defp do_numerals(num) when num > 1000 do
    # We don't expect more than 3999 to be entered
    translation = do_translation_for_range(div(num, 1000), "M", "", "")
    translation <> do_numerals(rem(num, 1000))
  end
  defp do_numerals(num) when num > 100 do
    translation = do_translation_for_range(div(num, 100), "C", "D", "M")
    translation <> do_numerals(rem(num, 100))
  end
  defp do_numerals(num) when num > 10 do
    translation = do_translation_for_range(div(num, 10), "X", "L", "C")
    translation <> do_numerals(rem(num, 10))
  end
  defp do_numerals(num) do
    translation = do_translation_for_range(num, "I", "V", "X")
    translation
  end


  defp do_translation_for_range(digit, incrementer, halfway, next) do
    case digit do
      digit  when digit < 4 ->
        String.duplicate(incrementer, digit)
      4 ->
        incrementer<>halfway
      digit when digit < 9 ->
        halfway<>String.duplicate(incrementer, digit-5)
      9 ->
        incrementer<>next
    end
  end
end
