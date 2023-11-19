defmodule Oplan.Payments do
  @moduledoc """
  The Payments context.
  """

  import Ecto.Query, warn: false
  alias Oplan.Repo

  alias Oplan.Payments.Tickets

  @doc """
  Returns the list of event_tickets.

  ## Examples

      iex> list_event_tickets()
      [%Tickets{}, ...]

  """
  def list_event_tickets do
    Repo.all(Tickets)
  end

  @doc """
  Gets a single tickets.

  Raises `Ecto.NoResultsError` if the Tickets does not exist.

  ## Examples

      iex> get_tickets!(123)
      %Tickets{}

      iex> get_tickets!(456)
      ** (Ecto.NoResultsError)

  """
  def get_tickets!(id), do: Repo.get!(Tickets, id)

  @doc """
  Creates a tickets.

  ## Examples

      iex> create_tickets(%{field: value})
      {:ok, %Tickets{}}

      iex> create_tickets(%{field: bad_value})
      {:error, %Ecto.Changeset{}}

  """
  def create_tickets(attrs \\ %{}) do
    %Tickets{}
    |> Tickets.changeset(attrs)
    |> Repo.insert()
  end

  @doc """
  Updates a tickets.

  ## Examples

      iex> update_tickets(tickets, %{field: new_value})
      {:ok, %Tickets{}}

      iex> update_tickets(tickets, %{field: bad_value})
      {:error, %Ecto.Changeset{}}

  """
  def update_tickets(%Tickets{} = tickets, attrs) do
    tickets
    |> Tickets.changeset(attrs)
    |> Repo.update()
  end

  @doc """
  Deletes a tickets.

  ## Examples

      iex> delete_tickets(tickets)
      {:ok, %Tickets{}}

      iex> delete_tickets(tickets)
      {:error, %Ecto.Changeset{}}

  """
  def delete_tickets(%Tickets{} = tickets) do
    Repo.delete(tickets)
  end

  @doc """
  Returns an `%Ecto.Changeset{}` for tracking tickets changes.

  ## Examples

      iex> change_tickets(tickets)
      %Ecto.Changeset{data: %Tickets{}}

  """
  def change_tickets(%Tickets{} = tickets, attrs \\ %{}) do
    Tickets.changeset(tickets, attrs)
  end
end
