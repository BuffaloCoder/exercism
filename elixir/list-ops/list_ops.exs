defmodule ListOps do
  # Please don't use any external modules (especially List or Enum) in your
  # implementation. The point of this exercise is to create these basic
  # functions yourself. You may use basic Kernel functions (like `Kernel.+/2`
  # for adding numbers), but please do not use Kernel functions for Lists like
  # `++`, `--`, `hd`, `tl`, `in`, and `length`.

  @spec count(list) :: non_neg_integer
  def count(l) do
    do_count(l, 0)
  end
  def do_count([], acc), do: acc
  def do_count([_|t], acc), do: do_count(t, acc+1)

  @spec reverse(list) :: list
  def reverse(l) do
    do_reverse(l, [])
  end
  def do_reverse([], acc), do: acc
  def do_reverse([h|t], acc), do: do_reverse(t, [h | acc])

  @spec map(list, (any -> any)) :: list
  def map(l, f) do
    l
    |> do_map(f, [])
    |> reverse
  end
  def do_map([], _, acc), do: acc
  def do_map([h|t], f, acc), do: do_map(t, f, [f.(h)|acc])

  @spec filter(list, (any -> as_boolean(term))) :: list
  def filter(l, f) do
    l
    |> do_filter(f, [])
    |> reverse
  end
  def do_filter([], _, acc), do: acc
  def do_filter([h|t], f, acc), do: do_filter(t, f, (if f.(h), do: [h|acc], else: acc))

  @type acc :: any
  @spec reduce(list, acc, (any, acc -> acc)) :: acc
  def reduce([], acc, _), do: acc
  def reduce([h|t], acc, f) do
    reduce(t, f.(h, acc), f)
  end


  @spec append(list, list) :: list
  def append(a, b) do
    a
    |> reverse
    |> reduce(b, &[&1 | &2])
  end

  @spec concat([[any]]) :: [any]
  def concat(ll) do
    fun = &[&1 | &2]
    ll
    |> reduce([], &reduce(&1, &2, fun))
    |> reverse
  end
end
