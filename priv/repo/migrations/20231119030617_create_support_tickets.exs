defmodule Oplan.Repo.Migrations.CreateSupportTickets do
  use Ecto.Migration

  def change do
    create table(:support_tickets, primary_key: false) do
      add :id, :uuid, primary_key: true
      add :message, :text
      add :user_id, references(:users, on_delete: :nothing)
      add :event_id, references(:events, on_delete: :nothing)

      timestamps()
    end

    create index(:support_tickets, [:user_id])
    create index(:support_tickets, [:event_id])
  end
end
