defmodule Oplan.Event.EventFollower do
  use Ecto.Schema
  import Ecto.Changeset

  alias Oplan.Account.User
  alias Oplan.Event.Event

  schema "event_followers" do
    belongs_to :user, User
    belongs_to :event, Event

    timestamps()
  end

  @doc false
  def changeset(event_follower, attrs) do
    event_follower
    |> cast(attrs, [:user_id, :event_id])
    |> validate_required([:user_id, :event_id])
  end
end
