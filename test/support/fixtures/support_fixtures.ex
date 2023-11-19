defmodule Oplan.SupportFixtures do
  @moduledoc """
  This module defines test helpers for creating
  entities via the `Oplan.Support` context.
  """

  @doc """
  Generate a support_tickets.
  """
  def support_tickets_fixture(attrs \\ %{}) do
    {:ok, support_tickets} =
      attrs
      |> Enum.into(%{
        id: "7488a646-e31f-11e4-aace-600308960662",
        message: "some message"
      })
      |> Oplan.Support.create_support_tickets()

    support_tickets
  end
end
