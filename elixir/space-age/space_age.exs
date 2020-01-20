defmodule SpaceAge do
  @type planet :: :mercury | :venus | :earth | :mars | :jupiter
                | :saturn | :uranus | :neptune

  @doc """
  Return the number of years a person that has lived for 'seconds' seconds is
  aged on 'planet'.
  """
  @spec age_on(planet, pos_integer) :: float
  def age_on(:mercury, seconds), do: do_age_on(seconds, 87.97)
  def age_on(:venus, seconds), do: do_age_on(seconds, 224.7)
  def age_on(:mars, seconds), do: do_age_on(seconds, 687)
  def age_on(:jupiter, seconds), do: do_age_on(seconds, 4_332.59)
  def age_on(:saturn, seconds), do: do_age_on(seconds, 10_759)
  def age_on(:uranus, seconds), do: do_age_on(seconds, 30_688.5)
  def age_on(:neptune, seconds), do: do_age_on(seconds, 60_182)
  # Return earth years in the base case
  def age_on(_planet, seconds), do: do_age_on(seconds, 365.25)

  def do_age_on(seconds, earth_days_per_year) do
    seconds
    |> Kernel./(earth_days_per_year * 24 * 3600)
    |> Float.round(2)
  end
end
