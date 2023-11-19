defmodule Oplan.Repo.Migrations.CreateEventTickets do
  use Ecto.Migration

  def change do
    create table(:event_tickets) do
      add :price, :integer
      add :currency, :string
      add :valid_from, :utc_datetime
      add :expiry, :utc_datetime
      add :used, :boolean, default: false
      add :accepted_mode_of_payment, {:array, :string}
      add :user_id, references(:users, on_delete: :nothing)
      add :event_id, references(:events, on_delete: :nothing)

      timestamps()
    end

    create unique_index(:event_tickets, [:user_id, :event_id])
  end
end
