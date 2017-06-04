defmodule ProteinTranslation do
  @mapping %{
      "UGU" => "Cysteine",
      "UGC" => "Cysteine",
      "UUA" => "Leucine",
      "UUG" => "Leucine",
      "AUG" => "Methionine",
      "UUU" => "Phenylalanine",
      "UUC" => "Phenylalanine",
      "UCU" => "Serine",
      "UCC" => "Serine",
      "UCA" => "Serine",
      "UCG" => "Serine",
      "UGG" => "Tryptophan",
      "UAU" => "Tyrosine",
      "UAC" => "Tyrosine",
      "UAA" => "STOP",
      "UAG" => "STOP",
      "UGA" => "STOP"
    }

  @doc """
  Given an RNA string, return a list of proteins specified by codons, in order.
  """
  @spec of_rna(String.t()) :: {atom,  list(String.t())}
  def of_rna(rna) do
    rna
    |> String.codepoints
    |> Enum.chunk(3)
    |> do_of_rna([])
  end

  defp do_of_rna([], acc), do: {:ok, Enum.reverse(acc)}
  defp do_of_rna([h|t], acc) do
      case h |> Enum.join |> of_codon do
        {:ok, "STOP"} ->
          {:ok, Enum.reverse(acc)}
        {:ok, codon} ->
          do_of_rna(t, [codon|acc])
        {:error, _} ->
          {:error, "invalid RNA"}
      end
  end

  @doc """
  Given a codon, return the corresponding protein

  UGU -> Cysteine
  UGC -> Cysteine
  UUA -> Leucine
  UUG -> Leucine
  AUG -> Methionine
  UUU -> Phenylalanine
  UUC -> Phenylalanine
  UCU -> Serine
  UCC -> Serine
  UCA -> Serine
  UCG -> Serine
  UGG -> Tryptophan
  UAU -> Tyrosine
  UAC -> Tyrosine
  UAA -> STOP
  UAG -> STOP
  UGA -> STOP
  """
  @spec of_codon(String.t()) :: {atom, String.t()}
  def of_codon(codon) do
    if Map.has_key?(@mapping, codon) do
      {:ok, @mapping[codon]}
    else
      {:error, "invalid codon"}
    end
  end
end

