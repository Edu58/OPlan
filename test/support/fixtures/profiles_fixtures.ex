defmodule Oplan.ProfilesFixtures do
  @moduledoc """
  This module defines test helpers for creating
  entities via the `Oplan.Profiles` context.
  """

  @doc """
  Generate a profile.
  """
  def profile_fixture(attrs \\ %{}) do
    {:ok, profile} =
      attrs
      |> Enum.into(%{
        avatar: "some avatar",
        city: "some city",
        date_of_birth: "some date_of_birth",
        nationality: "some nationality",
        phone_number: "some phone_number"
      })
      |> Oplan.Profiles.create_profile()

    profile
  end
end
