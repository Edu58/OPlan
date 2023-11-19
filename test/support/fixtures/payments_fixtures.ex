defmodule Oplan.PaymentsFixtures do
  @moduledoc """
  This module defines test helpers for creating
  entities via the `Oplan.Payments` context.
  """

  @doc """
  Generate a tickets.
  """
  def tickets_fixture(attrs \\ %{}) do
    {:ok, tickets} =
      attrs
      |> Enum.into(%{
        accepted_mode_of_payment: ["option1", "option2"],
        currency: "some currency",
        price: 42
      })
      |> Oplan.Payments.create_tickets()

    tickets
  end
end
