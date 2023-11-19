defmodule Oplan.Support do
  @moduledoc """
  The Support context.
  """

  import Ecto.Query, warn: false
  alias Oplan.Repo

  alias Oplan.Support.SupportTickets

  @doc """
  Returns the list of support_tickets.

  ## Examples

      iex> list_support_tickets()
      [%SupportTickets{}, ...]

  """
  def list_support_tickets do
    Repo.all(SupportTickets)
  end

  @doc """
  Gets a single support_tickets.

  Raises `Ecto.NoResultsError` if the Support tickets does not exist.

  ## Examples

      iex> get_support_tickets!(123)
      %SupportTickets{}

      iex> get_support_tickets!(456)
      ** (Ecto.NoResultsError)

  """
  def get_support_tickets!(id), do: Repo.get!(SupportTickets, id)

  @doc """
  Creates a support_tickets.

  ## Examples

      iex> create_support_tickets(%{field: value})
      {:ok, %SupportTickets{}}

      iex> create_support_tickets(%{field: bad_value})
      {:error, %Ecto.Changeset{}}

  """
  def create_support_tickets(attrs \\ %{}) do
    %SupportTickets{}
    |> SupportTickets.changeset(attrs)
    |> Repo.insert()
  end

  @doc """
  Updates a support_tickets.

  ## Examples

      iex> update_support_tickets(support_tickets, %{field: new_value})
      {:ok, %SupportTickets{}}

      iex> update_support_tickets(support_tickets, %{field: bad_value})
      {:error, %Ecto.Changeset{}}

  """
  def update_support_tickets(%SupportTickets{} = support_tickets, attrs) do
    support_tickets
    |> SupportTickets.changeset(attrs)
    |> Repo.update()
  end

  @doc """
  Deletes a support_tickets.

  ## Examples

      iex> delete_support_tickets(support_tickets)
      {:ok, %SupportTickets{}}

      iex> delete_support_tickets(support_tickets)
      {:error, %Ecto.Changeset{}}

  """
  def delete_support_tickets(%SupportTickets{} = support_tickets) do
    Repo.delete(support_tickets)
  end

  @doc """
  Returns an `%Ecto.Changeset{}` for tracking support_tickets changes.

  ## Examples

      iex> change_support_tickets(support_tickets)
      %Ecto.Changeset{data: %SupportTickets{}}

  """
  def change_support_tickets(%SupportTickets{} = support_tickets, attrs \\ %{}) do
    SupportTickets.changeset(support_tickets, attrs)
  end
end
