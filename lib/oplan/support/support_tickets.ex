defmodule Oplan.Support.SupportTickets do
  use Ecto.Schema
  import Ecto.Changeset

  alias Oplan.Account.User
  alias Oplan.Event.Event

  @primary_key false
  schema "support_tickets" do
    field :id, Ecto.UUID, primary_key: true
    field :message, :string
    belongs_to :user, User
    belongs_to :event, Event

    timestamps()
  end

  @doc false
  def changeset(support_tickets, attrs) do
    support_tickets
    |> cast(attrs, [:id, :message, :user_id, :event_id])
    |> validate_required([:id, :message, :user_id, :event_id])
  end
end
