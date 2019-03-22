defmodule RNATranscription do
  @doc """
  Transcribes a character list representing DNA nucleotides to RNA

  ## Examples

  iex> RNATranscription.to_rna('ACTG')
  'UGAC'
  """
  @spec to_rna([char]) :: [char]
  def to_rna(dna) do
    dna
    |> Enum.map(
      fn dna_piece ->
        case dna_piece do
          ?G -> ?C
          ?C -> ?G
          ?T -> ?A
          ?A -> ?U
        end
      end
    )
  end

end
