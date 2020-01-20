defmodule Bob do
  def hey(input) do
    cond do
      Regex.match?(~r/^\s*$/, input) -> "Fine. Be that way!"
      String.last(input) == "?" -> "Sure."
      Regex.match?(~r/\p{Lu}/, input) && !Regex.match?(~r/\p{Ll}/, input) -> "Whoa, chill out!"
      true -> "Whatever."

    end
  end
end
