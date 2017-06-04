defmodule Strain do
  @doc """
  Given a `list` of items and a function `fun`, return the list of items where
  `fun` returns true.

  Do not use `Enum.filter`.
  """
  @spec keep(list :: list(any), fun :: ((any) -> boolean)) :: list(any)
  def keep(list, fun) do
    do_keep(list, fun, [])
  end

  defp do_keep([], _fun, acc), do: Enum.reverse(acc)
  defp do_keep([h | t], fun, acc),
    do: do_keep(t, fun, (if fun.(h), do: [h | acc], else: acc))

  @doc """
  Given a `list` of items and a function `fun`, return the list of items where
  `fun` returns false.

  Do not use `Enum.reject`.
  """
  @spec discard(list :: list(any), fun :: ((any) -> boolean)) :: list(any)
  def discard(list, fun) do
    do_discard(list, fun, [])
  end

  defp do_discard([], _fun, acc), do: Enum.reverse(acc)
  defp do_discard([h | t], fun, acc),
    do: do_discard(t, fun, (if fun.(h), do: acc, else: [h | acc]))
end
