defmodule Words do
  @doc """
  Count the number of words in the sentence.

  Words are compared case-insensitively.
  """
  @spec count(String.t) :: map
  def count(sentence) do
    word_list = create_list(sentence)
    word_count(%{}, word_list)
  end

  @spec create_list(Integer.t) :: list
  defp create_list(sentence) do
    lower_sentence = String.downcase(sentence)
    Regex.scan(~r/([-\pL\d]+)/u, lower_sentence)
  end

  @spec word_count(map, List.t) :: map
  defp word_count(map, nil), do: map
  defp word_count(map, []), do: map
  defp word_count(map, [[capture, word]|other_words]) do
    map
    |> Map.get_and_update(
      word,
      fn
        nil -> {nil, 1}
        val -> {val, val + 1}
      end
    )
    |> elem(1)
    |>word_count(other_words)
  end
end
