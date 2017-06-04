defmodule NucleotideCount do
  @nucleotides [?A, ?C, ?G, ?T]
  @default_histogram Map.new(@nucleotides, &({&1, 0}))

  @doc """
  Counts individual nucleotides in a NucleotideCount strand.

  ## Examples

  iex> NucleotideCount.count('AATAA', ?A)
  4

  iex> NucleotideCount.count('AATAA', ?T)
  1
  """
  @spec count([char], char) :: non_neg_integer
  def count(strand, nucleotide) do
    Enum.count(strand, fn(found) -> found == nucleotide end)
  end


  @doc """
  Returns a summary of counts by nucleotide.

  ## Examples

  iex> NucleotideCount.histogram('AATAA')
  %{?A => 4, ?T => 1, ?C => 0, ?G => 0}
  """
  @spec histogram([char]) :: map
  def histogram(strand) do
    Enum.reduce(strand, @default_histogram, fn(nucleotide, results) ->
      %{results | nucleotide => Map.get(results, nucleotide, 0) + 1}
    end)
  end
end
