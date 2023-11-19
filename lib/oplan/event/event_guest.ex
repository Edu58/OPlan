defmodule Oplan.Event.EventGuest do
  use Ecto.Schema
  import Ecto.Changeset

  alias Oplan.Account.User
  alias Oplan.Event.Event

  schema "event_guests" do
    belongs_to :user, User
    belongs_to :event, Event
    field :confirmed, :boolean, default: false
    field :arrived, :boolean, default: false
    field :departed, :boolean, default: false

    timestamps()
  end

  @doc false
  def changeset(event_guest, attrs) do
    event_guest
    |> cast(attrs, [:user_id, :event_id, :confirmed, :arrived, :departed])
    |> validate_required([:user_id, :event_id])
  end
end
